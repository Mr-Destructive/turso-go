package turso

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Organizations struct {
	client *client
}

type organization struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Type string `json:"type"`
}

type organizationMembers struct {
	Role     string `json:"role"`
	Username string `json:"username"`
}

type database struct {
	Name            string   `json:"name"`
	Hostname        string   `json:"hostname"`
	IssuedCertLimit int      `json:"issuedCertLimit"`
	IssuedCertCount int      `json:"issuedCertCount"`
	DbId            string   `json:"dbId"`
	Regions         []string `json:"regions"`
	PrimaryRegion   string   `json:"primaryRegion"`
}

type orgDBMonthUsage struct {
	Database struct {
		UUID      string `json:"uuid"`
		Instances struct {
			UUID  string `json:"uuid"`
			Usage struct {
				RowsRead     int `json:"rows_read"`
				RowsWritten  int `json:"rows_written"`
				StorageBytes int `json:"storage_bytes"`
			} `json:"usage"`
		} `json:"instances"`
		Usage struct {
			RowsRead     int `json:"rows_read"`
			RowsWritten  int `json:"rows_written"`
			StorageBytes int `json:"storage_bytes"`
		} `json:"usage"`
	} `json:"database"`
}

type instance struct {
	UUID     string `json:"uuid"`
	Hostname string `json:"hostname"`
	Region   string `json:"region"`
	Type     string `json:"type"`
	Name     string `json:"name"`
}

type organizationList struct {
	Organizations []organization `json:"organizations"`
}

type organizationMembersList struct {
	Members []organizationMembers `json:"members"`
}

type organizationDatabaseList struct {
	Databases []database `json:"databases"`
}

type organizationDatabase struct {
	Database struct {
		database
	} `json:"database"`
}

type databaseInstances struct {
	Instances []instance `json:"instances"`
}

type databaseInstance struct {
	Instance struct {
		instance
	} `json:"instance"`
}

func (org *Organizations) List() (*organizationList, error) {
	endpoint := fmt.Sprintf("%s/v1/organizations", tursoBaseURL)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var organizations = organizationList{}
	err = json.NewDecoder(resp.Body).Decode(&organizations)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return &organizations, nil
}

func (org *Organizations) Members(organizationSlug string) (*organizationMembersList, error) {
	if organizationSlug == "" {
		return nil, fmt.Errorf("organization name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/members", tursoBaseURL, organizationSlug)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var members = organizationMembersList{}
	err = json.NewDecoder(resp.Body).Decode(&members)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return &members, nil
}

func (org *Organizations) Databases(organizationSlug string) (*organizationDatabaseList, error) {
	if organizationSlug == "" {
		return nil, fmt.Errorf("organization slug is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases", tursoBaseURL, organizationSlug)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var database = organizationDatabaseList{}
	err = json.NewDecoder(resp.Body).Decode(&database)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return &database, nil
}

func (org *Organizations) Database(orgSlug, dbName string) (*organizationDatabase, error) {
	if orgSlug == "" {
		return nil, fmt.Errorf("organization slug is required")
	}
	if dbName == "" {
		return nil, fmt.Errorf("database name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases/%s", tursoBaseURL, orgSlug, dbName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var database = organizationDatabase{}
	json.NewDecoder(resp.Body).Decode(&database)
	defer resp.Body.Close()
	return &database, nil
}

func (org *Organizations) DBUsage(orgName, dbName string) (*orgDBMonthUsage, error) {
	if orgName == "" {
		return nil, fmt.Errorf("organization name is required")
	}
	if dbName == "" {
		return nil, fmt.Errorf("database name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases/%s/usage", tursoBaseURL, orgName, dbName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var usage = orgDBMonthUsage{}
	json.NewDecoder(resp.Body).Decode(&usage)
	defer resp.Body.Close()
	return &usage, nil
}

func (org *Organizations) Instances(orgName, dbName string) (*databaseInstances, error) {
	if orgName == "" {
		return nil, fmt.Errorf("organization name is required")
	}
	if dbName == "" {
		return nil, fmt.Errorf("database name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases/%s/instances", tursoBaseURL, orgName, dbName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var instances = databaseInstances{}
	json.NewDecoder(resp.Body).Decode(&instances)
	defer resp.Body.Close()
	return &instances, nil
}

func (org *Organizations) Instance(orgName, dbName, instanceName string) (*databaseInstance, error) {
	if orgName == "" {
		return nil, fmt.Errorf("organization name is required")
	}
	if dbName == "" {
		return nil, fmt.Errorf("database name is required")
	}
	if instanceName == "" {
		return nil, fmt.Errorf("instance name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases/%s/instances/%s", tursoBaseURL, orgName, dbName, instanceName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var instance = databaseInstance{}
	json.NewDecoder(resp.Body).Decode(&instance)
	defer resp.Body.Close()
	return &instance, nil
}
