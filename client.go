package turso

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
    client client
	Tokens
	Organizations
	Locations
}

type client struct {
	baseURL  string
	apiToken string
	api      *http.Client
}

const tursoBaseURL = "https://api.turso.tech"

func NewClient(baseURL, apiToken string) (*Client, error) {
	if baseURL == "" {
		baseURL = tursoBaseURL
	}
	if apiToken == "" {
		return nil, fmt.Errorf("apiToken is required")
	}
	connection := &client{
		baseURL:  baseURL,
		apiToken: apiToken,
		api:      &http.Client{},
	}
    client := &Client{
        client: *connection,
    }
	client.Tokens = Tokens{
		client: connection,
	}
	client.Organizations = Organizations{
		client: connection,
	}
	client.Locations = Locations{
		client: connection,
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
