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
		AuthenticationMethod string    `json:"authenticationMethod,omitempty"`
		NameInfo             NameInfo  `json:"nameInfo,omitempty"`
		NotaryAuthentication string    `json:"notaryAuthentication,omitempty"`
		Password             string    `json:"password,omitempty"`
		PhoneInfo            PhoneInfo `json:"phoneInfo,omitempty"`
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

type CCParticipantInfo struct {
	// Company of the CC participant, if available
	Company string `json:"company,omitempty"`
	// Email of the CC participant of the agreement
	Email string `json:"email"`
	// Hidden True if the agreement is hidden for the user that is calling the API. Only returned if self is true
	Hidden bool `json:"hidden,omitempty"`
	// Name of the CC participant, if available
	Name string `json:"name,omitempty"`
	// ParticipantId The unique identifier of the CC participant of the agreement
	ParticipantId string `json:"participantId"`
	// Self True if the CC participant is the same user that is calling the API
	Self bool `json:"self"`
}

type DigAuthInfo struct {
	// ProviderId Digital Identity Gateway Provider Id. When replacing a participant that has DIG_ID authentication
	//specified, you must supply a provider id for the new participant.,
	ProviderId string `json:"providerId"`
	// ProviderDesc Digital Identity Gateway Provider Description. This will be ignored as part of POST or PUT calls.,
	ProviderDesc string `json:"providerDesc,omitempty"`
	// ProviderName Digital Identity Gateway Provider Name. This will be ignored as part of POST or PUT calls.
	ProviderName string `json:"ProviderName,omitempty"`
}

type NameInfo struct {
	// FirstName Recipient's first name,
	FirstName string `json:"firstName,omitempty"`
	// LastName Recipient's last name
	LastName string `json:"lastName,omitempty"`
}

type PhoneInfo struct {
	// CountryCode The numeric country calling code (ISD code) required for the participant to view and sign the
	//document if authentication type is PHONE
	CountryCode string `json:"countryCode,omitempty"`
	// CountryIsoCode The country ISO Alpha-2 code required for the participant to view and sign the document if
	//authentication method is PHONE
	CountryIsoCode string `json:"countryIsoCode,omitempty"`
	// Phone The phone number required for the participant to view and sign the document if authentication method is
	//PHONE. When replacing a participant that has PHONE authentication specified, you must supply a phone number for
	//the new participant
	Phone string `json:"phone,omitempty"`
}

type ParticipantSecurityOption struct {
	// AuthenticationMethod ['NONE' or 'PASSWORD' or 'PHONE' or 'KBA' or 'WEB_IDENTITY' or 'ADOBE_SIGN' or 'GOV_ID' or
	//'DIG_ID']: The authentication method for the participants to have access to view and sign the document. When
	//replacing a participant that has PASSWORD or PHONE authentication specified, you must supply a password or phone
	//number for the new participant, and you cannot change the authentication method,
	AuthenticationMethod string `json:"authenticationMethod"`
	// DigAuthInfo Digital Identity Gateway Provider information.
	DigAuthInfo DigAuthInfo `json:"digAuthInfo,omitempty"`
	// NameInfo Recipient's full name,
	NameInfo NameInfo `json:"nameInfo,omitempty"`
	// NotaryAuthentication ['MULTI_FACTOR_AUTHENTICATION' or 'NONE']: The authentication method of the notary
	//participant to the notary,
	NotaryAuthentication string `json:"notaryAuthentication,omitempty"`
	// Password The password required for the participant to view and sign the document. Note that AdobeSign will never
	//show this password to anyone, so you will need to separately communicate it to any relevant parties. The password
	//will not be returned in GET call. When replacing a participant that has PASSWORD authentication specified, you
	//must supply a password for the new participant.
	Password string `json:"password,omitempty"`
	// PhoneInfo The phoneInfo required for the participant to view and sign the document
	PhoneInfo PhoneInfo `json:"phoneInfo,omitempty"`
}

type DetailedParticipantInfo struct {
	// Email of the participant. In case of modifying a participant set (PUT) this is a required field. In case of GET,
	//this is the required field and will always be returned unless it is a fax workflow (legacy agreements) that were
	//created using fax as input,
	Email string `json:"email"`
	// ParticipantSetInfo Security options that apply to the participant.
	SecurityOption ParticipantSecurityOption `json:"securityOption"`
	// Company of the participant, if available. This cannot be changed as part of the PUT call.
	Company string `json:"company,omitempty"`
	// CreatedDate The date when the participant was added. This will be returned as part of GET call but is ignored if
	//passed as part of PUT call.,
	CreatedDate string `json:"createdDate,omitempty"`
	// Hidden True if the agreement is hidden for the user that is calling the API. Only returned if self is true.
	//Ignored (not required) if modifying a participant (PUT).
	Hidden bool `json:"hidden,omitempty"`
	// Id The unique identifier of the participant. This will be returned as part of Get call but is not mandatory to
	//be passed as part of PUT call for agreements/{id}/members/participantSets/{id}.,
	Id string `json:"id,omitempty"`
	// Name The name of the participant, if available. This cannot be changed as part of the PUT call.
	Name string `json:"name,omitempty"`
	// PrivateMessage The private message of the participant, if available. This cannot be changed as part of the PUT
	//call.
	PrivateMessage string `json:"privateMessage,omitempty"`
	// True if this participant is the same user that is calling the API. Returned as part of Get. Ignored
	//(not required) if modifying a participant set (PUT).,
	Self bool `json:"self,omitempty"`
	// Status ['REPLACED' or 'ACTIVE']: The status of the participant. This cannot be changed as part of the PUT call.
	//New participants will be ignored if added with a REPLACED status.,
	Status string `json:"status,omitempty"`
	// UserId The user ID of the participant. This will be returned as part of GET call but is ignored if passed as
	//part of PUT call.
	UserId string `json:"userId,omitempty"`
}

type DetailedParticipantSetInfo struct {
	// MemberInfos  Array of ParticipantInfo objects, containing participant-specific data (e.g. email). All
	//participants in the array belong to the same set
	MemberInfos []DetailedParticipantInfo `json:"memberInfos"`
	// Order Index indicating sequential signing group (specified for hybrid routing). This cannot be changed as part
	//of the PUT call.
	Order int `json:"order"`
	// Role ['SIGNER' or 'SENDER' or 'APPROVER' or 'ACCEPTOR' or 'CERTIFIED_RECIPIENT' or 'FORM_FILLER' or
	//'DELEGATE_TO_SIGNER' or 'DELEGATE_TO_APPROVER' or 'DELEGATE_TO_ACCEPTOR' or 'DELEGATE_TO_CERTIFIED_RECIPIENT' or
	//'DELEGATE_TO_FORM_FILLER' or 'SHARE' or 'NOTARY_SIGNER']: Role assumed by all participants in the set (signer,
	//approver etc.). This cannot be changed as part of the PUT call.,
	Role string `json:"role"`
	// Id The unique identifier of the participant set. This cannot be changed as part of the PUT call.
	Id string `json:"id,omitempty"`
	// Name of ParticipantSet (it can be empty, but needs not to be unique in a single agreement). Maximum no of
	//characters in participant set name is restricted to 255. This cannot be changed as part of the PUT call.
	Name string `json:"name,omitempty"`
	// PrivateMessage Participant set's private message - all participants in the set will receive the same message.
	//This cannot be changed as part of the PUT call.
	PrivateMessage string `json:"privateMessage,omitempty"`
	// Status ['CANCELLED' or 'COMPLETED' or 'EXPIRED' or 'NOT_YET_VISIBLE' or 'WAITING_FOR_NOTARIZATION' or
	//'WAITING_FOR_OTHERS' or 'WAITING_FOR_MY_APPROVAL' or 'WAITING_FOR_AUTHORING' or 'WAITING_FOR_MY_ACKNOWLEDGEMENT'
	//or 'WAITING_FOR_MY_ACCEPTANCE' or 'WAITING_FOR_MY_FORM_FILLING' or 'WAITING_FOR_MY_DELEGATION' or
	//'WAITING_FOR_MY_SIGNATURE' or 'WAITING_FOR_MY_VERIFICATION' or 'WAITING_FOR_PREFILL']: The agreement status with
	//respect to the participant set. This cannot be changed as part of the PUT call.
	Status string `json:"status,omitempty"`
}

type SenderInfo struct {
	// Company of the sender, if available.
	Company string `json:"company"`
	// Email of the sender of the agreement.
	Email string `json:"email`
	// Hidden True if the agreement is hidden for the user that is calling the API. Only returned if self is true.
	Hidden bool `json:"hidden"`
	// Name of the sender, if available.
	Name string `json:"name"`
	// ParticipantId The unique identifier of the sender of the agreement.
	ParticipantId string `json:"participantId"`
	// Self True if the sender is the same user that is calling the API.
	Self bool `json:"self"`
	// Status ['CANCELLED' or 'COMPLETED' or 'EXPIRED' or 'NOT_YET_VISIBLE' or 'WAITING_FOR_AUTHORING' or
	//'WAITING_FOR_MY_DELEGATION' or 'WAITING_FOR_MY_ACCEPTANCE' or 'WAITING_FOR_MY_ACKNOWLEDGEMENT' or
	//'WAITING_FOR_MY_APPROVAL' or 'WAITING_FOR_MY_FORM_FILLING' or 'WAITING_FOR_MY_SIGNATURE' or
	//'WAITING_FOR_NOTARIZATION' or 'WAITING_FOR_OTHERS']: The agreement status with respect to the participant set.
	//This cannot be changed as part of the PUT call.,
	Status string `json:"status"`
	// CreatedDate The date when the sender was added. This will be returned as part of GET call but is ignored if
	//passed as part of PUT call.,
	CreatedDate string `json:"createdDate"`
	// UserId The user ID of the sender. This will be returned as part of GET call but is ignored if passed as part of
	//PUT call.
	UserId string `json:"userId"`
}

type ShareParticipantInfo struct {
	// Company of the sharee participant, if available.
	Company string `json:"company"`
	// Email of the sharee participant of the agreement.
	Email string `json:"email"`
	// Hidden True if the agreement is hidden for the user that is calling the API. Only returned if self is true.
	Hidden bool `json:"hidden"`
	// Name of the sharee participant, if available.
	Name string `json:"name"`
	// The unique identifier of the sharee participant of the agreement.
	ParticipantId string `json:"participantId"`
	// Self True if the Share participant is the same user that is calling the API.,
	Self bool `json:"self"`
	// SharerParticipantId The unique identifier of the participant who shared the agreement.
	SharerParticipantId string `json:"sharerParticipantId"`
}

type MembersInfo struct {
	// CCsInfo Information of CC participants of the agreement.
	CCsInfo []CCParticipantInfo `json:"ccsInfo"`
	// NextParticipantSets Information of next participant sets.,
	NextParticipantSets []DetailedParticipantSetInfo `json:"nextParticipantSets"`
	// ParticipantSets Information about the participant Sets.,
	ParticipantSets []DetailedParticipantSetInfo `json:"participantSets"`
	// SenderInfo Information of the sender of the agreement.,
	SenderInfo SenderInfo `json:"senderInfo"`
	// Information of the participants with whom the agreement has been shared.
	SharesInfo []ShareParticipantInfo `json:"sharesInfo"`
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

// CreateReminder Creates a reminder on the specified participants of an existing AdobeSign Agreement
// ref: https://secure.na1.adobesign.com/public/docs/restapi/v6#!/agreements/createReminderOnParticipant
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

// GetAgreementMembers Retrieves information of members of an existing AdobeSign Agreement
// ref: https://secure.na1.adobesign.com/public/docs/restapi/v6#!/agreements/getAllMembers
func (s *AgreementService) GetAgreementMembers(ctx context.Context, agreementId string) (*MembersInfo, error) {
	u := fmt.Sprintf("%s/%s/members", agreementsPath, agreementId)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var response *MembersInfo
	if _, err := s.client.Do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}
