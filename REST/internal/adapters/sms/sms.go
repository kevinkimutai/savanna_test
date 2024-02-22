package smspkg

import (
	"net/http"
	"time"

	"github.com/edwinwalela/africastalking-go/pkg/sms"
)

type Adapter struct {
	client      sms.Client
	bulkrequest sms.BulkRequest
}

func NewAdapter(username string, apiKey string) *Adapter {
	client := sms.Client{
		ApiKey:    apiKey,
		Username:  username,
		IsSandbox: true,
		Client:    &http.Client{},
	}

	bulkRequest := sms.BulkRequest{
		Username:      username,
		To:            []string{},
		Message:       "",
		From:          "AFRICASTKNG",
		BulkSMSMode:   true,
		Enqueue:       true,
		RetryDuration: 24 * time.Hour,
	}

	return &Adapter{client: client, bulkrequest: bulkRequest}

}

func (a Adapter) SendSMS() (string, error) {
	response, err := a.client.SendBulk(&a.bulkrequest)
	if err != nil {
		return response.Message, err
	}

	return response.Message, nil
}
