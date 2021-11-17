package adobesign

import (
	"context"
)

// WebhookService handles operations related to webhooks.
//
// ref: https://helpx.adobe.com/sign/using/adobe-sign-webhooks-api.html
// ref: https://secure.na1.echosign.com/public/docs/restapi/v6#!/webhooks/
type WebhookService service

const webhooksPath = "webhooks"

type WebhookPayload struct {
	WebhookId                          string         `json:"webhookId"`
	WebhookName                        string         `json:"webhookName"`
	WebhookNotificationId              string         `json:"webhookNotificationId"`
	WebhookUrlInfo                     WebhookUrlInfo `json:"webhookUrlInfo"`
	WebhookScope                       string         `json:"webhookScope"`
	WebhookNotificationApplicableUsers []struct {
		Id                string `json:"id"`
		Email             string `json:"email"`
		Role              string `json:"role"`
		PayloadApplicable string `json:"payloadApplicable"`
	} `json:"webhookNotificationApplicableUsers"`
	Event                   string `json:"event"`
	SubEvent                string `json:"subEvent"`
	EventDate               string `json:"eventDate"`
	EventResourceType       string `json:"eventResourceType"`
	EventResourceParentType string `json:"eventResourceParentType"`
	EventResourceParentId   string `json:"eventResourceParentId"`
	ParticipantRole         string `json:"participantRole"`
	ActionType              string `json:"actionType"`
	ParticipantUserId       string `json:"participantUserId"`
	ParticipantUserEmail    string `json:"participantUserEmail"`
	ActingUserId            string `json:"actingUserId"`
	ActingUserEmail         string `json:"actingUserEmail"`
	ActingUserIpAddress     string `json:"actingUserIpAddress"`
	InitiatingUserId        string `json:"initiatingUserId"`
	InitiatingUserEmail     string `json:"initiatingUserEmail"`
	Agreement               struct {
		Id            string `json:"id"`
		Name          string `json:"name"`
		SignatureType string `json:"signatureType"`
		Status        string `json:"status"`
		Ccs           []struct {
			Email        string   `json:"email"`
			Label        string   `json:"label"`
			VisiblePages []string `json:"visiblePages"`
		} `json:"ccs"`
		DeviceInfo struct {
			ApplicationDescription string `json:"applicationDescription"`
			DeviceDescription      string `json:"deviceDescription"`
			Location               struct {
				Latitude  string `json:"latitude"`
				Longitude string `json:"longitude"`
			} `json:"location"`
			DeviceTime string `json:"deviceTime"`
		} `json:"deviceInfo"`
		DocumentVisibilityEnabled string `json:"documentVisibilityEnabled"`
		CreatedDate               string `json:"createdDate"`
		ExpirationTime            string `json:"expirationTime"`
		ExternalId                struct {
			Id string `json:"id"`
		} `json:"externalId"`
		PostSignOption struct {
			RedirectDelay string `json:"redirectDelay"`
			RedirectUrl   string `json:"redirectUrl"`
		} `json:"postSignOption"`
		FirstReminderDelay string `json:"firstReminderDelay"`
		Locale             string `json:"locale"`
		Message            string `json:"message"`
		ReminderFrequency  string `json:"reminderFrequency"`
		SenderEmail        string `json:"senderEmail"`
		VaultingInfo       struct {
			Enabled string `json:"enabled"`
		} `json:"vaultingInfo"`
		WorkflowId          string `json:"workflowId"`
		ParticipantSetsInfo struct {
			ParticipantSets []struct {
				MemberInfos []struct {
					Id             string `json:"id"`
					Email          string `json:"email"`
					Company        string `json:"company"`
					Name           string `json:"name"`
					PrivateMessage string `json:"privateMessage"`
					Status         string `json:"status"`
				} `json:"memberInfos"`
				Order          string `json:"order"`
				Role           string `json:"role"`
				Status         string `json:"status"`
				Id             string `json:"id"`
				Name           string `json:"name"`
				PrivateMessage string `json:"privateMessage"`
			} `json:"participantSets"`
		} `json:"participantSetsInfo"`
		DocumentsInfo struct {
			Documents []struct {
				Id       string `json:"id"`
				Label    string `json:"label"`
				NumPages string `json:"numPages"`
				MimeType string `json:"mimeType"`
				Name     string `json:"name"`
			} `json:"documents"`
			SupportingDocuments []struct {
				DisplayLabel string `json:"displayLabel"`
				FieldName    string `json:"fieldName"`
				Id           string `json:"id"`
				MimeType     string `json:"mimeType"`
				NumPages     string `json:"numPages"`
			} `json:"supportingDocuments"`
		} `json:"documentsInfo"`
	} `json:"agreement"`
}

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
func (s *WebhookService) CreateWebhook(ctx context.Context, request CreateWebhookRequest) (*CreateWebhookResponse, error) {

	req, err := s.client.NewRequest("POST", webhooksPath, request)
	if err != nil {
		return nil, err
	}

	var response *CreateWebhookResponse
	if _, err := s.client.Do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}
