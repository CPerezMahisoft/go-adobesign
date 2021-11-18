package adobesign

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	oauthApiVersion = "v2"
	apiVersion = "v6"
	userAgent       = "go-adobesign"
	apiBaseUrl      = "https://api.%s.adobesign.com/api/rest/v6/"

	headerRateLimit = "Retry-After"

	headerApiVersion = "Accept-Version"
)

var errNonNilContext = errors.New("context must be non-nil")

func Endpoint(baseUrl string) oauth2.Endpoint {
	return oauth2.Endpoint{
		AuthURL:  fmt.Sprintf("%s/public/oauth/%s", baseUrl, oauthApiVersion),
		TokenURL: fmt.Sprintf("%s/oauth/%s/token", baseUrl, oauthApiVersion),
	}
}

// A Client manages communication with the Adobe Sign API.
type Client struct {
	clientMu sync.Mutex   // clientMu protects the client during calls that modify the CheckRedirect func.
	client   *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. Defaults to the public Adobe Sign API, but can be
	// set to a domain endpoint to use with Adobe Sign Enterprise. BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL

	// Send requests impersonating user
	ImpersonatedUser string

	// User agent used when communicating with the Adobe Sign API.
	UserAgent string

	rateMu    sync.Mutex
	rateLimit Rate // Rate limits for the client as determined by the most recent API calls.

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	//// Services used for talking to different parts of the Adobe Sign API.
	TransientDocumentService *TransientDocumentService
	AgreementService         *AgreementService
	WebhookService           *WebhookService
}

type service struct {
	client *Client
}

// Client returns the http.Client used by this Adobe Sign client.
func (c *Client) Client() *http.Client {
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	clientCopy := *c.client
	return &clientCopy
}

type Oauth2Params struct {
	ClientId     string   `json:"clientId"`
	ClientSecret string   `json:"clientSecret"`
	Scopes       []string `json:"scopes"`
	BaseUrl      string   `json:"baseUrl"`
	RedirectUri  string   `json:"redirectUri"`
}

func NewOauth2Client(params Oauth2Params) *Client {
	ctx := context.Background()
	conf := &oauth2.Config{
		RedirectURL:  params.RedirectUri,
		ClientID:     params.ClientId,
		ClientSecret: params.ClientSecret,
		Scopes:       params.Scopes,
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("")
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	return newClient(conf.Client(ctx, tok), params.BaseUrl, "")
}

// NewClient creates an adobe sign client using an Integration Key, this method is deprecated.
// New integrations should use the NewOauth2Client method.
// ref: https://helpx.adobe.com/sign/kb/how-to-create-an-integration-key.html
func NewClient(integrationKey string, shard string, impersonating string) *Client {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: integrationKey},
	)
	tc := oauth2.NewClient(ctx, ts)

	return newClient(tc, fmt.Sprintf(apiBaseUrl, shard), impersonating)
}

func newClient(httpClient *http.Client, baseUrl string, impersonating string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	baseURL, _ := url.Parse(baseUrl)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	c.common.client = c

	c.ImpersonatedUser = impersonating

	c.TransientDocumentService = (*TransientDocumentService)(&c.common)
	c.AgreementService = (*AgreementService)(&c.common)
	c.WebhookService = (*WebhookService)(&c.common)

	return c
}

func (c *Client) impersonate(req *http.Request) *http.Request {
	if c.ImpersonatedUser != "" {
		req.Header.Set("x-on-behalf-of-user", fmt.Sprintf("email:%s", c.ImpersonatedUser))
	}
	return req
}

func (c *Client) NewMultiPartRequest(urlStr string, body io.ReadWriter) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set(headerApiVersion, oauthApiVersion)
	req = c.impersonate(req)

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil

}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set(headerApiVersion, apiVersion)
	req = c.impersonate(req)

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Response is a Adobe Sign API response. This wraps the standard http.Response
// returned from Adobe Sign and provides convenient access to things like
// pagination links.
type Response struct {
	*http.Response

	//// These fields provide the page values for paginating through a set of
	//// results. Any or all of these may be set to the zero value for
	//// responses that are not part of a paginated set, or for which there
	//// are no additional pages.
	////
	//// These fields support what is called "offset pagination" and should
	//// be used with the ListOptions struct.
	NextCursor int

	// Explicitly specify the Rate type so Rate's String() receiver doesn't
	// propagate to Response.
	Rate Rate
	//
	//// token's expiration date
	//TokenExpiration time.Time
}

// newResponse creates a new Response for the provided http.Response.
// r must not be nil.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}

	//TODO: response.populatePageValues()

	response.Rate = parseRate(r)
	//response.TokenExpiration = parseTokenExpiration(r)
	return response
}

// Rate represents the rate limit for the current client.
// ref: https://www.adobe.io/apis/documentcloud/sign/docs.html#!adobedocs/adobe-sign/master/api_usage/throttling.md
type Rate struct {
	// RetryAfter is the number of seconds that the client should wait before retrying new requests.
	RetryAfterSeconds int `json:"retryAfter"`
}

// parseRate parses the rate related headers.
func parseRate(r *http.Response) Rate {
	var rate Rate
	if wait := r.Header.Get(headerRateLimit); wait != "" {
		rate.RetryAfterSeconds, _ = strconv.Atoi(wait)
	}

	return rate
}

// RateLimitError occurs when Adobe Sign returns 403 Forbidden response with a rate limit
// remaining value of 0.
type RateLimitError struct {
	Rate     Rate           // Rate specifies last known rate limit for the client
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // error message
}

func (r *RateLimitError) Error() string {
	return fmt.Sprintf("%v %v: %d %v %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message, r.Rate.RetryAfterSeconds)
}

// Is returns whether the provided error equals this error.
func (r *RateLimitError) Is(target error) bool {
	v, ok := target.(*RateLimitError)
	if !ok {
		return false
	}

	return r.Rate == v.Rate &&
		r.Message == v.Message &&
		compareHTTPResponse(r.Response, v.Response)
}

// compareHTTPResponse returns whether two http.Response objects are equal or not.
// Currently, only StatusCode is checked. This function is used when implementing the
// Is(error) bool interface for the custom error types in this package.
func compareHTTPResponse(r1, r2 *http.Response) bool {
	if r1 == nil && r2 == nil {
		return true
	}

	if r1 != nil && r2 != nil {
		return r1.StatusCode == r2.StatusCode
	}
	return false
}

// checkRateLimitBeforeDo does not make any network calls, but uses existing knowledge from
// current client state in order to quickly check if *RateLimitError can be immediately returned
// from Client.Do, and if so, returns it so that Client.Do can skip making a network API call unnecessarily.
// Otherwise, it returns nil, and Client.Do should proceed normally.
func (c *Client) checkRateLimitBeforeDo(req *http.Request) *RateLimitError {
	c.rateMu.Lock()
	rate := c.rateLimit
	c.rateMu.Unlock()
	if rate.RetryAfterSeconds > 0 {
		// Create a fake response.
		resp := &http.Response{
			Status:     http.StatusText(http.StatusForbidden),
			StatusCode: http.StatusForbidden,
			Request:    req,
			Header:     make(http.Header),
			Body:       ioutil.NopCloser(strings.NewReader("")),
		}
		return &RateLimitError{
			Rate:     rate,
			Response: resp,
			Message: fmt.Sprintf("API rate limit reached. Wait %v seconds before making more requests.",
				rate.RetryAfterSeconds),
		}
	}

	return nil
}

// BareDo sends an API request and lets you handle the api response. If an error
// or API Error occurs, the error will contain more information. Otherwise you
// are supposed to read and close the response's Body. If rate limit is exceeded
// and reset time is in the future, BareDo returns *RateLimitError immediately
// without making a network API call.
//
// The provided ctx must be non-nil, if it is nil an error is returned. If it is
// canceled or times out, ctx.Err() will be returned.
func (c *Client) BareDo(ctx context.Context, req *http.Request) (*Response, error) {
	if ctx == nil {
		return nil, errNonNilContext
	}

	//if bypass := ctx.Value(bypassRateLimitCheck); bypass == nil {
	// If we've hit rate limit, don't make further requests before Reset time.
	if err := c.checkRateLimitBeforeDo(req); err != nil {
		return &Response{
			Response: err.Response,
			//Rate:     err.Rate,
		}, err
	}
	//}

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}

	response := newResponse(resp)

	//c.rateMu.Lock()
	//c.rateLimit = response.Rate
	//c.rateMu.Unlock()

	err = CheckResponse(resp)
	return response, err
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer interface,
// the raw response body will be written to v, without attempting to first
// decode it. If v is nil, and no error happens, the response is returned as is.
// If rate limit is exceeded and reset time is in the future, Do returns
// *RateLimitError immediately without making a network API call.
//
// The provided ctx must be non-nil, if it is nil an error is returned. If it
// is canceled or times out, ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.BareDo(ctx, req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}
	return resp, err
}

// An ErrorResponse reports one or more errors caused by an API request.
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"`
	Code     string         `json:"code"`
	Err      string         `json:"err"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message, r.Err)
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		if err := json.Unmarshal(data, errorResponse); err != nil {
			return err
		}
	}

	switch {
	case r.StatusCode == http.StatusTooManyRequests && r.Header.Get(headerRateLimit) != "0":
		return &RateLimitError{
			Rate:     parseRate(r),
			Response: errorResponse.Response,
			Message:  errorResponse.Message,
		}
	default:
		return errorResponse
	}
}

// ListOptions specifies the optional parameters to various List methods that
// support offset pagination.
type ListOptions struct {
	// Maximum number of Items to be returned (max limit: 100)
	Cursor int `url:"cursor,omitempty"`

	// 	Offset used for pagination if collection has more than limit items
	PageSize int `url:"pageSize,omitempty"`
}

// addOptions adds the parameters in opts as URL query parameters to s. opts
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

//PageInfo holds the pagination information for a Adobe Sign API request
type PageInfo struct {
	NextCursor int `json:"nextCursor,omitempty"`
}
