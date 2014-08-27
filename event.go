package chain

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// WebhookEvent represents a Chain WebhooEvent.
//
// Chain documentation can be found here
// https://chain.com/docs#object-webhooks.
type WebhookEvent struct {
	ID            string
	WebhookID     string `json:"webhook_id"`
	Event         string
	BlockChain    string `json:"block_chain"`
	Address       string
	Confirmations int64
}

// WebhookEventType represents a webhook event type. Currently only
// address-transaction type is supported.
type WebhookEventType string

// AddressTransactionEventType triggers when a new transaction occurs on a
// specified address. The first POST will notify your application of the new,
// unconfirmed transaction. Additional POSTs will notify your application of
// subsequent confirmations for that transction.
const AddressTransactionEventType WebhookEventType = "address-transaction"

// CreateWebhookEvent creates a Webhook Event that makes POST requests to the
// associated Webhook when triggered.
//
// Chain documentation can be found here
// https://chain.com/docs#webhooks-event-list.
func (c *Chain) CreateWebhookEvent(webhookID string, event WebhookEventType,
	network Network, address string,
	confirmations int64) (WebhookEvent, error) {

	weResponse := WebhookEvent{}
	url := fmt.Sprintf("%s/webhooks/%s/events", c.network, webhookID)

	var blockChain string
	switch network {
	case TestNet3:
		blockChain = "testnet3"
	case MainNet:
		blockChain = "bitcoin"
	default:
		return weResponse, errors.New("unknown network")
	}

	jsonRequest := struct {
		Event         string
		BlockChain    string `json:"block_chain"`
		Address       string
		Confirmations int64
	}{string(event), blockChain, address, confirmations}
	requestBody, err := json.Marshal(&jsonRequest)

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

// ListWebhookEvents lists all Webhook Events associated with a Webhook.
//
// Chain documentation can be found here
// https://chain.com/docs#webhooks-event-list.
func (c *Chain) ListWebhookEvents(webhookID string) ([]WebhookEvent, error) {
	url := fmt.Sprintf("%s/webhooks/%s/events", c.network, webhookID)

	webhookEvents := []WebhookEvent{}
	return webhookEvents, c.httpGetJSON(url, &webhookEvents)
}

// DeleteWebhookEvent deletes a Webhook Event, which will stop all further
// POST requests for the event.
//
// Chain documentation can be found here
// https://chain.com/docs#webhooks-event-delete.
func (c *Chain) DeleteWebhookEvent(
	webhookID, eventType, address string) (WebhookEvent, error) {
	url := fmt.Sprintf("%s/webhooks/%s/events/%s/%s",
		c.network, webhookID, eventType, address)

	we := WebhookEvent{}
	return we, c.httpDeleteJSON(url, &we)
}
