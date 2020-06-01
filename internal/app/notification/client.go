package notification

import (
	"bytes"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"

	"github.com/gmbanaag/wallet2020/internal/app/logger"
)

//NotifyUser struct of message to be sent to the notificaion service
type NotifyUser struct {
	TransactionID     string `json:"transaction_id"`
	SourceUserID      uint32 `json:"source_user_id"`
	DestinationUserID string `json:"destination_user_id"`
	Amount            string `json:"amount"`
	Message           string `json:"message"`
	TransactionTs     string `json:"transaction_ts"`
}

//Client options
type Client struct {
	Host   string
	APIKey string
}

//SendNotification sends notification to the service
func (c Client) SendNotification(notifyUser NotifyUser) {
	request, _ := json.Marshal(notifyUser)
	uri := fmt.Sprintf("%s/notify", c.Host)
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(request))

	if err != nil {
		logger.LogError(fmt.Sprintf("http.Do %s", err.Error()))
	}

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", c.APIKey))
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.LogError(fmt.Sprintf("http.Do %s", err.Error()))
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.LogError(fmt.Sprintf("ReadAll: %s", err.Error()))

	}
	defer resp.Body.Close()
}
