package turso

import (
	"net/http"
	"os"
	"testing"
)

func newClient() (*Client, error) {
	baseURL := ""
	apiToken := os.Getenv("TURSO_AUTH_TOKEN")
	client, err := NewClient(baseURL, apiToken)
	if err != nil || client == nil {
		return nil, err
	}
	return client, err
}

func TestNewClientNil(t *testing.T) {
	baseURL := ""
	apiToken := ""
	client, err := NewClient(baseURL, apiToken)
	if err == nil || client != nil {
		t.Errorf("NewClient without a apiToken should fail")
	}
}

func TestNewClient(t *testing.T) {
	client, err := newClient()
	if err != nil || client == nil {
		t.Error(err)
	}
	client.client.tursoAPIrequest("/", http.MethodGet, nil)
	if client.client.api == nil {
		t.Errorf("Failed to create api connection")
	}
}
