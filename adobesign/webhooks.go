package adobesign

import (
	"context"
	"fmt"
)

// WebhookService handles operations related to webhooks.
//
// ref: https://helpx.adobe.com/sign/using/adobe-sign-webhooks-api.html
// ref: https://secure.na1.echosign.com/public/docs/restapi/v6#!/webhooks/
type WebhookService service

const webhooksPath = "webhooks"

type WebhookUrlInfo struct {
	Url string `json:"url,omitempty"`
}

type WebhookAgreementEvents struct {
	IncludeDetailedInfo     bool `json:"includeDetailedInfo,omitempty"`
	IncludeDocumentsInfo    bool `json:"includeDocumentsInfo,omitempty"`
	IncludeParticipantsInfo bool `json:"includeParticipantsInfo,omitempty"`
	IncludeSignedDocuments  bool `json:"includeSignedDocuments,omitempty"`
}

type WebhookLibraryDocumentEvents struct {
	IncludeDetailedInfo  bool `json:"includeDetailedInfo,omitempty"`
	IncludeDocumentsInfo bool `json:"includeDocumentsInfo,omitempty"`
}

type WebhookMegaSignEvents struct {
	IncludeDetailedInfo bool `json:"includeDetailedInfo,omitempty"`
}

type WebhookWidgetEvents struct {
	IncludeDetailedInfo     bool `json:"includeDetailedInfo,omitempty"`
	IncludeDocumentsInfo    bool `json:"includeDocumentsInfo,omitempty"`
	IncludeParticipantsInfo bool `json:"includeParticipantsInfo,omitempty"`
}

type WebhookConditionalParams struct {
	WebhookAgreementEvents       WebhookAgreementEvents       `json:"webhookAgreementEvents,omitempty"`
	WebhookLibraryDocumentEvents WebhookLibraryDocumentEvents `json:"webhookLibraryDocumentEvents,omitempty"`
	WebhookMegaSignEvents        WebhookMegaSignEvents        `json:"webhookMegaSignEvents,omitempty"`
	WebhookWidgetEvents          WebhookWidgetEvents          `json:"webhookWidgetEvents,omitempty"`
}

type CreateWebhookRequest struct {
	Name                      string         `json:"name,omitempty"`
	Scope                     string         `json:"scope,omitempty"`
	State                     string         `json:"state,omitempty"`
	WebhookSubscriptionEvents []string       `json:"webhookSubscriptionEvents,omitempty"`
	WebhookUrlInfo            WebhookUrlInfo `json:"webhookUrlInfo,omitempty"`
	ApplicationDisplayName    string         `json:"applicationDisplayName,omitempty"`
	ApplicationName           string         `json:"applicationName,omitempty"`
	Created                   string         `json:"created,omitempty"`
	Id                        string         `json:"id,omitempty"`
	LastModified              string         `json:"lastModified,omitempty"`
	ProblemNotificationEmails []struct {
		Email string `json:"email,omitempty"`
	} `json:"problemNotificationEmails,omitempty"`
	ResourceId               string                   `json:"resourceId,omitempty"`
	ResourceType             string                   `json:"resourceType,omitempty"`
	Status                   string                   `json:"status,omitempty"`
	WebhookConditionalParams WebhookConditionalParams `json:"webhookConditionalParams,omitempty"`
}

type CreateWebhookResponse struct {
	Id string `json:"id"`
}

// CreateWebhook creates a new Adobe Sign Agreement
// ref: https://secure.na1.echosign.com/public/docs/restapi/v6#!/webhooks/createWebhook
// requires: `webhook_write` permissions https://secure.na1.echosign.com/public/static/oauthDoc.jsp#scope-webhook_write
func (s *WebhookService) CreateWebhook(ctx context.Context, request CreateWebhookRequest, onBehalfOf string) (*CreateWebhookResponse, error) {

	req, err := s.client.NewRequest("POST", webhooksPath, request)
	if err != nil {
		return nil, err
	}

	if onBehalfOf != "" { //impersonate user
		req.Header.Set("x-on-behalf-of-user", fmt.Sprintf("email:%s", onBehalfOf))
	}

	var response *CreateWebhookResponse
	if _, err := s.client.Do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}
