package turso

import (
	"testing"
)

func TestAuthTokens(t *testing.T) {
	client, err := newClient()
	if err != nil || client == nil {
		t.Error(err)
	}
	tokens, err := client.Tokens.List()
	if err != nil {
		t.Error(err)
	}
	if tokens == nil && len(tokens.Tokens) == 0 {
		t.Error("organizations should not be nil")
	}
}

func TestAuthTokenValidate(t *testing.T) {
	client, err := newClient()
	if err != nil || client == nil {
		t.Error(err)
	}
	tokenValidate, err := client.Tokens.Validate()
	if err != nil {
		t.Error(err)
	}
	if tokenValidate == nil {
		t.Error("organizations should not be nil")
	}
}
