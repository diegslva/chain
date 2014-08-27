package chain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type WebhookEvent struct {
	ID            string
	WebhookID     string `json:"webhook_id"`
	Event         string
	BlockChain    string `json:"block_chain"`
	Address       string
	Confirmations int64
}

func (c *Chain) CreateWebhookEvent(id, event, blockChain, address string,
	confirmations int64) (WebhookEvent, error) {

	url := fmt.Sprintf("%s/webhooks/%s/events", c.network, id)

	jsonRequest := struct {
		Event         string
		BlockChain    string `json:"block_chain"`
		Address       string
		Confirmations int64
	}{event, blockChain, address, confirmations}
	requestBody, err := json.Marshal(&jsonRequest)

	weResponse := WebhookEvent{}
	if err != nil {
		return weResponse, err
	}
	response, err := c.httpPostJSON(url, bytes.NewReader(requestBody))

	responseBody, err := ioutil.ReadAll(response)
	if err != nil {
		return weResponse, err
	}

	if err := json.Unmarshal(responseBody, &weResponse); err != nil {
		return weResponse, nil
	}
	return weResponse, nil
}

func (c *Chain) ListWebhookEvents(webhookID string) ([]WebhookEvent, error) {
	url := fmt.Sprintf("%s/webhooks/%s/events", c.network, webhookID)

	webhookEvents := []WebhookEvent{}
	return webhookEvents, c.httpGetJSON(url, &webhookEvents)
}

func (c *Chain) DeleteWebhookEvent(
	webhookID, eventType, address string) (WebhookEvent, error) {
	url := fmt.Sprintf("%s/webhooks/%s/events/%s/%s",
		c.network, webhookID, eventType, address)

	we := WebhookEvent{}
	return we, c.httpDeleteJSON(url, &we)
}
