package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aesadde/go-adobesign/adobesign"
)

func main() {
	client := adobesign.NewClient("YOUR_INTEGRATION_KEY", "YOUR_SHARD")

	file, err := os.Open("PATH_TO_FILE")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	document, err := client.TransientDocumentService.UploadTransientDocument(context.Background(), data, "TEST.docx", "")
	if err != nil {
		log.Fatal(err)
	}

	agreement, err := client.AgreementService.CreateAgreement(context.Background(), adobesign.CreateAgreementRequest{
		FileInfos: []adobesign.FileInfo{{TransientDocumentId: document.TransientDocumentId}},
		Name:      "test document",
		ParticipantSetsInfo: []adobesign.ParticipantSetInfo{{
			MemberInfos: []adobesign.MemberInfo{{Email: "test@test.com"}},
			Order:       1,                                // the order in which the signer appears on the document
			Role:        adobesign.ParticipantRole.Signer, // the role of the member
		},
		},
		SignatureType: adobesign.SignatureType.Esign,
		State:         adobesign.AgreementState.InProcess,
	}, "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(agreement)

	webhook, err := client.WebhookService.CreateWebhook(context.Background(), adobesign.CreateWebhookRequest{
		Name:                      "testing webhook",
		Scope:                     adobesign.Scope.Resource,
		State:                     "ACTIVE",
		WebhookSubscriptionEvents: []string{adobesign.WebhookSubscriptionEvent.AgreementAll},
		ResourceType:              adobesign.Resource.Agreement,
		ResourceId:                agreement.Id,
		WebhookUrlInfo:            adobesign.WebhookUrlInfo{Url: ""},
		WebhookConditionalParams: adobesign.WebhookConditionalParams{WebhookAgreementEvents: adobesign.WebhookAgreementEvents{
			IncludeDetailedInfo:     true,
			IncludeDocumentsInfo:    true,
			IncludeParticipantsInfo: true,
			IncludeSignedDocuments:  true,
		}},
	}, "")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(webhook)
}
