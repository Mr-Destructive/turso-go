package turso

import (
	"testing"
)

func TestOrganizationList(t *testing.T) {
	client, err := newClient()
	if err != nil || client == nil {
		t.Error(err)
	}
	_, err = client.Organizations.List()
	if err != nil {
		t.Error(err)
	}
}

func TestOrganizationMembers(t *testing.T) {
	client, err := newClient()
	if err != nil || client == nil {
		t.Error(err)
	}
	_, err = client.Organizations.Members("abc")
	if err != nil {
		t.Error(err)
	}
}

