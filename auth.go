package turso

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Tokens struct {
	client *client
}

type token struct {
	Name  string `json:"name"`
	Id    string `json:"id"`
	Token string `json:"token,omitempty"`
}

type tokenList struct {
	Tokens []token `json:"tokens"`
}

type tokenValidate struct {
	Expiration time.Duration `json:"exp"`
}

func (t *Tokens) List() (*tokenList, error) {
	endpoint := fmt.Sprintf("%s/v1/auth/api-tokens", tursoBaseURL)
	resp, err := t.client.tursoAPIrequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var tokens tokenList
    if err := json.NewDecoder(resp.Body).Decode(&tokens); err != nil {
        return nil, err
    }
	defer resp.Body.Close()
	return &tokens, nil
}

func (t *Tokens) Mint(name string) (*token, error) {
	if name == "" {
		return nil, fmt.Errorf("token name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/auth/api-tokens/%s", tursoBaseURL, name)
	resp, err := t.client.tursoAPIrequest(endpoint, http.MethodPost, nil)
	if err != nil {
		return nil, err
	}
	var token token
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return &token, nil
}

func (t *Tokens) Revoke(tokenName string) error {
	if tokenName == "" {
		return fmt.Errorf("token name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/auth/api-tokens/%s", tursoBaseURL, tokenName)
	resp, err := t.client.tursoAPIrequest(endpoint, http.MethodDelete, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (t *Tokens) Validate() (*tokenValidate, error) {
	endpoint := fmt.Sprintf("%s/v1/auth/validate", tursoBaseURL)
	resp, err := t.client.tursoAPIrequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var tokenValidate tokenValidate
	if err := json.NewDecoder(resp.Body).Decode(&tokenValidate); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return &tokenValidate, nil
}
