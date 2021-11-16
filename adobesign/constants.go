package adobesign

// Resource defines the type of resources that can be used in the API.
var Resource = struct {
	Agreement       string
	Widget          string
	MegaSign        string
	LibraryDocument string
}{
	Agreement:       "AGREEMENT",
	Widget:          "WIDGET",
	MegaSign:        "MEGASIGN",
	LibraryDocument: "LIBRARY_DOCUMENT",
}

// Scope defines the supported scopes for the API.
var Scope = struct {
	Account  string
	Group    string
	User     string
	Resource string
}{
	Account:  "ACCOUNT",
	Group:    "GROUP",
	User:     "USER",
	Resource: "RESOURCE",
}

// WebhookSubscriptionEvent defines the supported webhook events.
var WebhookSubscriptionEvent = struct {
	AgreementAll                            string
	AgreementCreated                        string
	AgreementActionRequested                string
	AgreementActionCompleted                string
	AgreementWorkflowCompleted              string
	AgreementExpired                        string
	AgreementDocumentsDeleted               string
	AgreementRecalled                       string
	AgreementRejected                       string
	AgreementShared                         string
	AgreementActionDelegated                string
	AgreementActionReplacedSigner           string
	AgreementModified                       string
	AgreementUserAckAgreementModified       string
	AgreementEmailViewed                    string
	AgreementEmailBounced                   string
	AgreementAutoCancelledConversionProblem string
	AgreementOfflineSync                    string
	AgreementUploadedBySender               string
	AgreementVaulted                        string
	AgreementWebIdentityAuthenticated       string
	AgreementKbaAuthenticated               string
}{
	AgreementAll:                            "AGREEMENT_ALL",
	AgreementCreated:                        "AGREEMENT_CREATED",
	AgreementActionRequested:                "AGREEMENT_ACTION_REQUESTED",
	AgreementActionCompleted:                "AGREEMENT_ACTION_COMPLETED",
	AgreementWorkflowCompleted:              "AGREEMENT_WORKFLOW_COMPLETED",
	AgreementExpired:                        "AGREEMENT_EXPIRED",
	AgreementDocumentsDeleted:               "AGREEMENT_DOCUMENTS_DELETED",
	AgreementRecalled:                       "AGREEMENT_RECALLED",
	AgreementRejected:                       "AGREEMENT_REJECTED",
	AgreementShared:                         "AGREEMENT_SHARED",
	AgreementActionDelegated:                "AGREEMENT_ACTION_DELEGATED",
	AgreementActionReplacedSigner:           "AGREEMENT_ACTION_REPLACED_SIGNER",
	AgreementModified:                       "AGREEMENT_MODIFIED",
	AgreementUserAckAgreementModified:       "AGREEMENT_USER_ACK_AGREEMENT_MODIFIED",
	AgreementEmailViewed:                    "AGREEMENT_EMAIL_VIEWED",
	AgreementEmailBounced:                   "AGREEMENT_EMAIL_BOUNCED",
	AgreementAutoCancelledConversionProblem: "AGREEMENT_AUTO_CANCELLED_CONVERSION_PROBLEM",
	AgreementOfflineSync:                    "AGREEMENT_OFFLINE_SYNC",
	AgreementUploadedBySender:               "AGREEMENT_UPLOADED_BY_SENDER",
	AgreementVaulted:                        "AGREEMENT_VAULTED",
	AgreementWebIdentityAuthenticated:       "AGREEMENT_WEB_IDENTITY_AUTHENTICATED",
	AgreementKbaAuthenticated:               "AGREEMENT_KBA_AUTHENTICATED",
}

// AgreementState defines the valid states of an Agreement.
var AgreementState = struct {
	Authoring string
	Draft     string
	InProcess string
}{
	Authoring: "AUTHORING",
	Draft:     "DRAFT",
	InProcess: "IN_PROCESS",
}

// ParticipantRole defines the valid roles of a participant in an Agreement.
var ParticipantRole = struct {
	Signer                       string
	Approver                     string
	Acceptor                     string
	CertifiedRecipient           string
	FormFiller                   string
	DelegateToSigner             string
	DelegateToApprover           string
	DelegateToAcceptor           string
	DelegateToCertifiedRecipient string
	DelegateToFormFiller         string
	Share                        string
	NotarySigner                 string
}{
	Signer:                       "SIGNER",
	Approver:                     "APPROVER",
	Acceptor:                     "ACCEPTOR",
	CertifiedRecipient:           "CERTIFIED_RECIPIENT",
	FormFiller:                   "FORM_FILLER",
	DelegateToSigner:             "DELEGATE_TO_SIGNER",
	DelegateToApprover:           "DELEGATE_TO_APPROVER",
	DelegateToAcceptor:           "DELEGATE_TO_ACCEPTOR",
	DelegateToCertifiedRecipient: "DELEGATE_TO_CERTIFIED_RECIPIENT",
	DelegateToFormFiller:         "DELEGATE_TO_FORM_FILLER",
	Share:                        "SHARE",
	NotarySigner:                 "NOTARY_SIGNER",
}

// SignatureType defines the valid signature methods of an Agreement.
var SignatureType = struct {
	Esign   string
	Written string
}{
	Esign:   "ESIGN",
	Written: "WRITTEN",
}
