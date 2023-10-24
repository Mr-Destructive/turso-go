package turso

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type client struct {
	baseURL  string
	apiToken string
	api      *http.Client
	Tokens
	Organizations
	Locations
}

const tursoBaseURL = "https://api.turso.tech"

func NewClient(baseURL, apiToken string) (*client, error) {
	if baseURL == "" {
		baseURL = tursoBaseURL
	}
	if apiToken == "" {
		return nil, fmt.Errorf("apiToken is required")
	}
	client := &client{
		baseURL:  baseURL,
		apiToken: apiToken,
		api:      &http.Client{},
	}
	client.Tokens = Tokens{
		client: client,
	}
	client.Organizations = Organizations{
		client: client,
	}
	client.Locations = Locations{
		client: client,
	}
	return client, nil
}

func (client *client) tursoAPIrequest(endpoint string, method string, body interface{}) (*http.Response, error) {
	endpointURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprintf("Bearer %s", client.apiToken))
	req := &http.Request{
		Method: method,
		URL:    endpointURL,
		Header: headers,
	}
	if body != nil {
		req.Body = io.NopCloser(bytes.NewBuffer([]byte(body.(string))))
	}
	resp, err := client.api.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
