package adobesign

import (
	"context"
	"fmt"
)

const agreementsPath = "agreements"

// AgreementService handles operations related to agreements
//
// ref: https://secure.na1.echosign.com/public/docs/restapi/v6#!/agreements
type AgreementService service

type Document struct {
	CreatedDate string `json:"createdDate,omitempty"`
	Id          string `json:"id,omitempty"`
	Label       string `json:"label,omitempty"`
	NumPages    int    `json:"numPages,omitempty"`
	MimeType    string `json:"mimeType,omitempty"`
	Name        string `json:"name,omitempty"`
}

type FileInfo struct {
	Document            Document `json:"document,omitempty"`
	Label               string   `json:"label,omitempty"`
	LibraryDocumentId   string   `json:"libraryDocumentId,omitempty"`
	Notarize            bool     `json:"notarize,omitempty"`
	TransientDocumentId string   `json:"transientDocumentId,omitempty"`
	UrlFileInfo         struct {
		MimeType string `json:"mimeType,omitempty"`
		Name     string `json:"name,omitempty"`
		Url      string `json:"url,omitempty"`
	} `json:"urlFileInfo,omitempty"`
}

type MemberInfo struct {
	Email          string `json:"email,omitempty"`
	Id             string `json:"id,omitempty"`
	SecurityOption struct {
		AuthenticationMethod string `json:"authenticationMethod,omitempty"`
		NameInfo             struct {
			FirstName string `json:"firstName,omitempty"`
			LastName  string `json:"lastName,omitempty"`
		} `json:"nameInfo,omitempty"`
		NotaryAuthentication string `json:"notaryAuthentication,omitempty"`
		Password             string `json:"password,omitempty"`
		PhoneInfo            struct {
			CountryCode    string `json:"countryCode,omitempty"`
			CountryIsoCode string `json:"countryIsoCode,omitempty"`
			Phone          string `json:"phone,omitempty"`
		} `json:"phoneInfo,omitempty"`
	} `json:"securityOption,omitempty"`
}

type ParticipantSetInfo struct {
	MemberInfos    []MemberInfo `json:"memberInfos,omitempty"`
	Order          int          `json:"order,omitempty"`
	Role           string       `json:"role,omitempty"`
	Id             string       `json:"id,omitempty"`
	Label          string       `json:"label,omitempty"`
	Name           string       `json:"name,omitempty"`
	PrivateMessage string       `json:"privateMessage,omitempty"`
	VisiblePages   []string     `json:"visiblePages,omitempty"`
}

// Agreement defines the request body for creating an agreement
// ref: https://secure.na1.echosign.com/public/docs/restapi/v6#!/agreements/createAgreement
type Agreement struct {
	FileInfos           []FileInfo           `json:"fileInfos,omitempty"`
	Name                string               `json:"name,omitempty"`
	ParticipantSetsInfo []ParticipantSetInfo `json:"participantSetsInfo,omitempty"`
	SignatureType       string               `json:"signatureType,omitempty"`
	State               string               `json:"state,omitempty"`
	Ccs                 []struct {
		Email        string   `json:"email,omitempty"`
		Label        string   `json:"label,omitempty"`
		VisiblePages []string `json:"visiblePages,omitempty"`
	} `json:"ccs,omitempty"`
	CreatedDate string `json:"createdDate,omitempty"`
	DeviceInfo  struct {
		ApplicationDescription string `json:"applicationDescription,omitempty"`
		DeviceDescription      string `json:"deviceDescription,omitempty"`
		DeviceTime             string `json:"deviceTime,omitempty"`
	} `json:"deviceInfo,omitempty"`
	DocumentVisibilityEnabled bool `json:"documentVisibilityEnabled,omitempty"`
	EmailOption               struct {
		SendOptions struct {
			CompletionEmails string `json:"completionEmails,omitempty"`
			InFlightEmails   string `json:"inFlightEmails,omitempty"`
			InitEmails       string `json:"initEmails,omitempty"`
		} `json:"sendOptions,omitempty"`
	} `json:"emailOption,omitempty"`
	ExpirationTime string `json:"expirationTime,omitempty"`
	ExternalId     struct {
		Id string `json:"id,omitempty"`
	} `json:"externalId,omitempty"`
	FirstReminderDelay      int `json:"firstReminderDelay,omitempty"`
	FormFieldLayerTemplates []struct {
		Document struct {
			CreatedDate string `json:"createdDate,omitempty"`
			Id          string `json:"id,omitempty"`
			Label       string `json:"label,omitempty"`
			NumPages    int    `json:"numPages,omitempty"`
			MimeType    string `json:"mimeType,omitempty"`
			Name        string `json:"name,omitempty"`
		} `json:"document,omitempty"`
		Label               string `json:"label,omitempty"`
		LibraryDocumentId   string `json:"libraryDocumentId,omitempty"`
		Notarize            bool   `json:"notarize,omitempty"`
		TransientDocumentId string `json:"transientDocumentId,omitempty"`
		UrlFileInfo         struct {
			MimeType string `json:"mimeType,omitempty"`
			Name     string `json:"name,omitempty"`
			Url      string `json:"url,omitempty"`
		} `json:"urlFileInfo,omitempty"`
	} `json:"formFieldLayerTemplates,omitempty"`
	GroupId                    string `json:"groupId,omitempty"`
	HasFormFieldData           bool   `json:"hasFormFieldData,omitempty"`
	HasSignerIdentityReport    bool   `json:"hasSignerIdentityReport,omitempty"`
	Id                         string `json:"id,omitempty"`
	IsDocumentRetentionApplied bool   `json:"isDocumentRetentionApplied,omitempty"`
	LastEventDate              string `json:"lastEventDate,omitempty"`
	Locale                     string `json:"locale,omitempty"`
	MergeFieldInfo             []struct {
		DefaultValue string `json:"defaultValue,omitempty"`
		FieldName    string `json:"fieldName,omitempty"`
	} `json:"mergeFieldInfo,omitempty"`
	Message    string `json:"message,omitempty"`
	NotaryInfo struct {
		Appointment string `json:"appointment,omitempty"`
		NotaryEmail string `json:"notaryEmail,omitempty"`
		NotaryType  string `json:"notaryType,omitempty"`
		Note        string `json:"note,omitempty"`
		Payment     string `json:"payment,omitempty"`
	} `json:"notaryInfo,omitempty"`
	ParentId       string `json:"parentId,omitempty"`
	PostSignOption struct {
		RedirectDelay int    `json:"redirectDelay,omitempty"`
		RedirectUrl   string `json:"redirectUrl,omitempty"`
	} `json:"postSignOption,omitempty"`
	ReminderFrequency string `json:"reminderFrequency,omitempty"`
	SecurityOption    struct {
		OpenPassword string `json:"openPassword,omitempty"`
	} `json:"securityOption,omitempty"`
	SenderEmail  string `json:"senderEmail,omitempty"`
	Status       string `json:"status,omitempty"`
	Type         string `json:"type,omitempty"`
	VaultingInfo struct {
		Enabled bool `json:"enabled,omitempty"`
	} `json:"vaultingInfo,omitempty"`
	WorkflowId string `json:"workflowId,omitempty"`
}

type CreateAgreementResponse struct {
	Id string `json:"id,omitempty"`
}

// CreateAgreement creates a new Adobe Sign Agreement
// ref: https://secure.na1.echosign.com/public/docs/restapi/v6#!/agreements/createAgreement
func (s *AgreementService) CreateAgreement(ctx context.Context, request Agreement) (*CreateAgreementResponse, error) {

	req, err := s.client.NewRequest("POST", agreementsPath, request)
	if err != nil {
		return nil, err
	}

	var response *CreateAgreementResponse
	if _, err := s.client.Do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil

}

// GetAgreement retrieves an existing Adobe Sign Agreement
// ref: https://secure.na1.echosign.com/public/docs/restapi/v6#!/agreements/getAgreementInfo
func (s *AgreementService) GetAgreement(ctx context.Context, agreementId string) (*Agreement, error) {

	u := fmt.Sprintf("%s/%s", agreementsPath, agreementId)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var response *Agreement
	if _, err := s.client.Do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetAuditTrail retrieves the PDF file stream containing audit trail information
// ref: https://secure.na1.echosign.com/public/docs/restapi/v6#!/agreements/getAuditTrail
func (s *AgreementService) GetAuditTrail(ctx context.Context, agreementId string) (string, error) {
	u := fmt.Sprintf("%s/%s/auditTrail", agreementsPath, agreementId)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return "", err
	}

	var response string
	if _, err := s.client.Do(ctx, req, &response); err != nil {
		return "", err
	}

	return response, nil
}

// GetCombinedDocument retrieves a single combined PDF document for the documents associated with an agreement
// ref: https://secure.na1.echosign.com/public/docs/restapi/v6#!/agreements/getCombinedDocument
func (s *AgreementService) GetCombinedDocument(ctx context.Context, agreementId string) (string, error) {
	u := fmt.Sprintf("%s/%s/combinedDocument", agreementsPath, agreementId)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return "", err
	}

	var response string
	if _, err := s.client.Do(ctx, req, &response); err != nil {
		return "", err
	}

	return response, nil
}
