package turso

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Locations struct {
	client *client
}

type locations struct {
	Locations map[string]string `json:"locations"`
}

func (loc *Locations) List() (*locations, error) {
	endpoint := fmt.Sprintf("%s/v1/locations", tursoBaseURL)
	resp, err := loc.client.tursoAPIrequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var locations = locations{}
    if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
        return nil, err
    }
	defer resp.Body.Close()
	return &locations, nil
}
