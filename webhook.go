package chain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Webhook struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

func (c *Chain) CreateWebhook(id, url string) (Webhook, error) {
	endpointURL := fmt.Sprintf("%s/webhooks", c.network)
	requestBody, err := json.Marshal(&Webhook{id, endpointURL})

	wh := Webhook{}
	if err != nil {
		return wh, err
	}
	response, err := c.httpPostJSON(url, bytes.NewReader(requestBody))

	responseBody, err := ioutil.ReadAll(response)
	if err != nil {
		return wh, err
	}

	if err := json.Unmarshal(responseBody, &wh); err != nil {
		return wh, nil
	}
	return wh, nil
}

func (c *Chain) ListWebhooks() ([]Webhook, error) {
	url, webhooks := fmt.Sprintf("%s/webhooks", c.network), []Webhook{}
	return webhooks, c.httpGetJSON(url, &webhooks)
}

func (c *Chain) UpdateWebhook(id, url string) (Webhook, error) {
	endpointURL := fmt.Sprintf("%s/webhooks/%s", c.network, id)
	requestBody, err := json.Marshal(&Webhook{id, endpointURL})

	wh := Webhook{}
	if err != nil {
		return wh, err
	}
	response, err := c.httpPutJSON(url, bytes.NewReader(requestBody))

	responseBody, err := ioutil.ReadAll(response)
	if err != nil {
		return wh, err
	}

	if err := json.Unmarshal(responseBody, &wh); err != nil {
		return wh, nil
	}
	return wh, nil
}

// DeleteWebhook deletes a Webhook and all associated Webhook Events.
//
// Chain.com documentation can be found here
// https://chain.com/docs#webhooks-delete.
func (c *Chain) DeleteWebhook(id string) ([]Webhook, error) {
	url := fmt.Sprintf("%s/webhooks/%s", c.network, id)

	webhooks := []Webhook{}
	return webhooks, c.httpDeleteJSON(url, &webhooks)
}
