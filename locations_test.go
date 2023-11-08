package turso

import (
	"testing"
)

func TestLocations(t *testing.T) {
	client, err := newClient()
	if err != nil || client == nil {
		t.Error(err)
	}
	locations, err := client.Locations.List()
	if err != nil {
		t.Error(err)
	}
	if locations == nil && len(locations.Locations) == 0 {
		t.Error("locations should not be nil")
	}
}
