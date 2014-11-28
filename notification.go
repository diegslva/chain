package chain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type NotificationResponse struct {
	ID         string
	State      string
	URL        string
	Type       string
	Address    string
	BlockChain string `json:"block_chain"`
}

func (c *Chain) CreateNewTxNotification(url string) (
	*NotificationResponse, error) {
	return c.createNewNotification(url, "new-transaction")
}

func (c *Chain) CreateNewBlockNotification(url string) (
	*NotificationResponse, error) {
	return c.createNewNotification(url, "new-block")
}

func (c *Chain) createNewNotification(url, ty string) (
	*NotificationResponse, error) {
	endpointURL := fmt.Sprintf("%s/notifications", baseURL)
	req := &struct {
		Type       string `json:"type"`
		BlockChain string `json:"block_chain"`
		URL        string `json:"url"`
	}{ty, string(c.network), url}

	requestBody, err := json.Marshal(req)
	response, err := c.httpPostJSON(endpointURL, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	responseBody, err := ioutil.ReadAll(response)
	if err != nil {
		return nil, err
	}

	resp := &NotificationResponse{}
	if err := json.Unmarshal(responseBody, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Chain) ListNotifications() ([]*NotificationResponse, error) {
	url := fmt.Sprintf("%s/notifications", baseURL)
	resp := []*NotificationResponse{}
	return resp, c.httpGetJSON(url, &resp)
}

func (c *Chain) DeleteNotification(id string) (*NotificationResponse, error) {
	url := fmt.Sprintf("%s/notifications/%s", baseURL, id)

	resp := &NotificationResponse{}
	return resp, c.httpDeleteJSON(url, resp)
}
