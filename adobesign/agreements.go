package adobesign

import (
	"bytes"
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

type MergeFieldInfo struct {
	DefaultValue string `json:"defaultValue,omitempty"`
	FieldName    string `json:"fieldName,omitempty"`
}

type Cc struct {
	Email        string   `json:"email,omitempty"`
	Label        string   `json:"label,omitempty"`
	VisiblePages []string `json:"visiblePages,omitempty"`
}

// Agreement defines the request body for creating an agreement
// ref: https://secure.na1.echosign.com/public/docs/restapi/v6#!/agreements/createAgreement
type Agreement struct {
	FileInfos           []FileInfo           `json:"fileInfos,omitempty"`
	Name                string               `json:"name,omitempty"`
	ParticipantSetsInfo []ParticipantSetInfo `json:"participantSetsInfo,omitempty"`
	SignatureType       string               `json:"signatureType,omitempty"`
	State               string               `json:"state,omitempty"`
	Ccs                 []Cc                 `json:"ccs,omitempty"`
	CreatedDate         string               `json:"createdDate,omitempty"`
	DeviceInfo          struct {
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
	GroupId                    string           `json:"groupId,omitempty"`
	HasFormFieldData           bool             `json:"hasFormFieldData,omitempty"`
	HasSignerIdentityReport    bool             `json:"hasSignerIdentityReport,omitempty"`
	Id                         string           `json:"id,omitempty"`
	IsDocumentRetentionApplied bool             `json:"isDocumentRetentionApplied,omitempty"`
	LastEventDate              string           `json:"lastEventDate,omitempty"`
	Locale                     string           `json:"locale,omitempty"`
	MergeFieldInfo             []MergeFieldInfo `json:"mergeFieldInfo,omitempty"`
	Message                    string           `json:"message,omitempty"`
	NotaryInfo                 struct {
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

type ReminderInfo struct {
	// RecipientParticipantIds A list of one or more participant IDs that the reminder should be sent to. These must
	//be recipients of the agreement and not sharees or cc's.
	RecipientParticipantIds []string `json:"recipientParticipantIds"`
	// Status ['ACTIVE' or 'CANCELED' or 'COMPLETE']: Current status of the reminder. The only valid update in a PUT
	//is from ACTIVE to CANCELED. Must be provided as ACTIVE in a POST.
	Status string `json:"status"`
	// FirstReminderDelay Integer which specifies the delay in hours before sending the first reminder.
	//This is an optional field. The minimum value allowed is 1 hour and the maximum value can’t be more than the
	//difference of agreement creation and expiry time of the agreement in hours.
	//If this is not specified but the reminder frequency is specified, then the first reminder will be sent based
	//on frequency. Cannot be updated in a PUT.
	FirstReminderDelay int `json:"firstReminderDelay,omitempty"`
	// Frequency ['DAILY_UNTIL_SIGNED' or 'WEEKDAILY_UNTIL_SIGNED' or 'EVERY_OTHER_DAY_UNTIL_SIGNED' or
	//'EVERY_THIRD_DAY_UNTIL_SIGNED' or 'EVERY_FIFTH_DAY_UNTIL_SIGNED' or 'WEEKLY_UNTIL_SIGNED' or 'ONCE']: The
	//frequency at which reminder will be sent until the agreement is completed.
	//If frequency is not provided, the reminder will be sent once (if the agreement is available at the specified
	//time) with the delay based on the firstReminderDelay field and will never repeat again. If the agreement is
	//not available at that time, reminder will not be sent. Cannot be updated in a PUT,
	Frequency string `json:"frequency,omitempty"`
	// LastSentDate The date when the reminder was last sent. Only provided in GET.
	LastSentDate string `json:"lastSentDate,omitempty"`
	// NextSentDate The date when the reminder is scheduled to be sent next. When provided in POST request, frequency
	//needs to be ONCE (or not specified), startReminderCounterFrom needs to be REMINDER_CREATION (or not specified)
	//and firstReminderDelay needs to be 0 (or not specified). Cannot be updated in a PUT. Format would be
	//yyyy-MM-dd'T'HH:mm:ssZ
	NextSentDate string `json:"nextSentDate,omitempty"`
	// Note An optional message sent to the recipients, describing why their participation is required.
	Note string `json:"note,omitempty"`
	// ReminderId An identifier of the reminder resource created on the server. If provided in POST or PUT, it will
	//be ignored
	ReminderId string `json:"reminderId,omitempty"`
	// StartReminderCounterFrom ['AGREEMENT_AVAILABILITY' or 'REMINDER_CREATION']: Reminder can be sent based on when
	//the agreement becomes available or when the reminder is created
	StartReminderCounterFrom string `json:"startReminderCounterFrom,omitempty"`
}

type ReminderCreationResult struct {
	Id string `json:"id"`
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
func (s *AgreementService) GetCombinedDocument(ctx context.Context, agreementId string) ([]byte, error) {
	u := fmt.Sprintf("%s/%s/combinedDocument", agreementsPath, agreementId)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var response bytes.Buffer
	if _, err := s.client.Do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response.Bytes(), nil
}

type AgreementCancellationInfo struct {
	Comment      string `json:"comment"`
	NotifyOthers bool   `json:"notifyOthers"`
}

type UpdateAgreementRequest struct {
	State                     string                    `json:"state"`
	AgreementCancellationInfo AgreementCancellationInfo `json:"agreementCancellationInfo"`
}

// UpdateAgreementState updates the state of an existing Adobe Sign Agreement
// ref: https://secure.na1.echosign.com/public/docs/restapi/v6#!/agreements/updateAgreementState
func (s *AgreementService) UpdateAgreementState(ctx context.Context, agreementId string, request UpdateAgreementRequest) error {

	u := fmt.Sprintf("%s/%s/state", agreementsPath, agreementId)

	req, err := s.client.NewRequest("PUT", u, request)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	return err
}

func (s *AgreementService) CreateReminder(ctx context.Context, agreementId string, request ReminderInfo) (*ReminderCreationResult, error) {
	u := fmt.Sprintf("%s/%s/reminders", agreementsPath, agreementId)

	req, err := s.client.NewRequest("POST", u, request)
	if err != nil {
		return nil, err
	}

	var response ReminderCreationResult
	if _, err = s.client.Do(ctx, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
