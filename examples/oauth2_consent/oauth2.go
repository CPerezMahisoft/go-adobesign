package main

import (
	"github.com/aesadde/go-adobesign/adobesign"
)

func main() {
	params := adobesign.Oauth2Params{
		ClientId:     "YOUR_CLIENT_ID",
		ClientSecret: "YOUR_CLIENT_SECRET",
		Scopes:       []string{"user_login:self", "agreement_send:account"},
		BaseUrl:      "YOUR_BASE_URL (example: secure.na1.adobesign.com)",
		RedirectUri:  "YOUR_REDIRECT_URI",
	}
	_ = adobesign.NewOauth2Client(params)
}