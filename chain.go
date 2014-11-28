package chain

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://api.chain.com/v2"

// Network is used to let the Chain context know which network it should
// connect to.
type Network string

const (
	// TestNet3 is used to make the Chain context connect to the Bitcon
	// TestNet3 network.
	TestNet3 Network = "testnet3"

	// MainNet is used to make the Chain context to connect to the Bitcon
	// MainNet network.
	MainNet Network = "bitcoin"
)

// Chain contains the context for connecting with the Chain.com API endpoints.
type Chain struct {
	client  *http.Client
	network Network

	apiKeyID     string
	apiKeySecret string
}

// MultiError is returned by *Multi functions when there are errors with
// particular elements. Errors will be in a one-to-one correspondence with the
// input elements; sucessful elements will have a nil entry.
type MultiError []error

func (m MultiError) Error() string {
	s, n := "", 0
	for _, e := range m {
		if e != nil {
			if n == 0 {
				s = e.Error()
			}
			n++
		}
	}
	switch n {
	case 0:
		return "(0 errors)"
	case 1:
		return s
	case 2:
		return s + " (and 1 other error)"
	}
	return fmt.Sprintf("%s (and %d other errors)", s, n-1)
}

// New creates a new chain object.
func New(c *http.Client, n Network, apiKeyID, apiKeySecret string) *Chain {
	return &Chain{c, n, apiKeyID, apiKeySecret}
}

func checkHTTPResponse(r *http.Response) error {
	switch r.StatusCode {
	case http.StatusOK, http.StatusCreated:
		return nil
	}

	errData, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		return err
	}

	jsonError := struct {
		Message string
		Error   string
	}{}
	if err := json.Unmarshal(errData, &jsonError); err == nil {
		message := string(errData)
		if jsonError.Message != "" {
			message = jsonError.Message
		} else if jsonError.Error != "" {
			message = jsonError.Error
		}
		return fmt.Errorf("%s %s", r.Request.URL, message)
	}

	return fmt.Errorf("%s %s %s", r.Request.URL, r.Status, errData)
}

func (c *Chain) httpGetJSON(url string, v interface{}) error {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	return c.doRequest(req, v)
}

func (c *Chain) httpDeleteJSON(url string, v interface{}) error {

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	return c.doRequest(req, v)
}

func (c *Chain) doRequest(req *http.Request, v interface{}) error {
	req.SetBasicAuth(c.apiKeyID, c.apiKeySecret)
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if err := checkHTTPResponse(resp); err != nil {
		return err
	}

	defer resp.Body.Close()
	return decodeJSON(resp.Body, v)
}

func (c *Chain) httpPostJSON(url string,
	body io.Reader) (io.ReadCloser, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	return c.doRequestWithBody(req)
}

func (c *Chain) httpPutJSON(url string, body io.Reader) (io.ReadCloser, error) {
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, err
	}
	return c.doRequestWithBody(req)
}

func (c *Chain) doRequestWithBody(req *http.Request) (io.ReadCloser, error) {
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.apiKeyID, c.apiKeySecret)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err := checkHTTPResponse(resp); err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func decodeJSON(r io.Reader, v interface{}) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("%s with data %.30q", err.Error(), data)
	}
	return nil
}
