package turso

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Organizations struct {
	client *client
}

type Organization struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Type string `json:"type"`
}

type OrganizationMembers struct {
	Role     string `json:"role"`
	Username string `json:"username"`
}

type OrganisationList struct {
	Orgs []Organization `json:"organizations"`
}

type Database struct {
	Name            string   `json:"name"`
	Hostname        string   `json:"hostname"`
	IssuedCertLimit int      `json:"issuedCertLimit"`
	IssuedCertCount int      `json:"issuedCertCount"`
	DbId            string   `json:"dbId"`
	Regions         []string `json:"regions"`
	PrimaryRegion   string   `json:"primaryRegion"`
}

type usage struct {
	RowsRead     int `json:"rows_read"`
	RowsWritten  int `json:"rows_written"`
	StorageBytes int `json:"storage_bytes"`
}

type instanceUsage struct {
	UUID  string `json:"uuid"`
	Usage usage  `json:"usage"`
}

type dbUsage struct {
	UUID      string        `json:"uuid"`
	Instances instanceUsage `json:"instances"`
	Usage     usage         `json:"usage"`
}

type DBMonthlyUsage struct {
	Database dbUsage `json:"database"`
}

type Instance struct {
	UUID     string `json:"uuid"`
	Hostname string `json:"hostname"`
	Region   string `json:"region"`
	Type     string `json:"type"`
	Name     string `json:"name"`
}

type organizationMembersList struct {
	Members []OrganizationMembers `json:"members"`
}

type organizationDatabaseList struct {
	Databases []Database `json:"databases"`
}

type organizationDatabase struct {
	Database struct {
		Database
	} `json:"database"`
}

type databaseInstances struct {
	Instances []Instance `json:"instances"`
}

type databaseInstance struct {
	Instance struct {
		Instance
	} `json:"instance"`
}

type jwtToken struct {
	JWT string `json:"jwt"`
}

func (org *Organizations) List() (*OrganisationList, error) {
	endpoint := fmt.Sprintf("%s/v1/organizations", tursoBaseURL)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var organizations = OrganisationList{}
	err = json.NewDecoder(resp.Body).Decode(&organizations.Orgs)
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

func (org *Organizations) MintToken(organizationSlug, dbName, expiration, authorization string) (*jwtToken, error) {
	if organizationSlug == "" {
		return nil, fmt.Errorf("organisation slug is required")
	}
	if dbName == "" {
		return nil, fmt.Errorf("database name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases/%s/auth/tokens", tursoBaseURL, organizationSlug, dbName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodPost, nil)
	if err != nil {
		return nil, err
	}
	var jwtToken = jwtToken{}
	err = json.NewDecoder(resp.Body).Decode(&jwtToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return &jwtToken, nil
}

func (org *Organizations) InvalidateTokens(organizationSlug, dbName string) error {
	if organizationSlug == "" {
		return fmt.Errorf("organization slug is required")
	}
	if dbName == "" {
		return fmt.Errorf("database name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases/%s/auth/tokens", tursoBaseURL, organizationSlug, dbName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodDelete, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
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

func (org *Organizations) CreateDatabase(orgName string, body map[string]string) (*Database, error) {
	if orgName == "" {
		return nil, fmt.Errorf("organization name is required")
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases", tursoBaseURL, orgName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodPost, b)
	if err != nil {
		return nil, err
	}
	var database = Database{}
	json.NewDecoder(resp.Body).Decode(&database)
	defer resp.Body.Close()
	return &database, nil
}

func (org *Organizations) DeleteDatabase(orgName, dbName string) error {
	if orgName == "" {
		return fmt.Errorf("organization name is required")
	}
	if dbName == "" {
		return fmt.Errorf("database name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases/%s", tursoBaseURL, orgName, dbName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodDelete, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (org *Organizations) UpdateAllInstances(orgName, dbName string) error {
	if orgName == "" {
		return fmt.Errorf("organization name is required")
	}
	if dbName == "" {
		return fmt.Errorf("database name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases/%s/update", tursoBaseURL, orgName, dbName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodPut, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (org *Organizations) DBUsage(orgName, dbName string) (*DBMonthlyUsage, error) {
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
	var usage = DBMonthlyUsage{}
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

func (org *Organizations) CreateInstance(orgName, dbName string, body map[string]string) (*databaseInstance, error) {
	if orgName == "" {
		return nil, fmt.Errorf("organization name is required")
	}
	if dbName == "" {
		return nil, fmt.Errorf("database name is required")
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases/%s/instances", tursoBaseURL, orgName, dbName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodPost, b)
	if err != nil {
		return nil, err
	}
	var instance = databaseInstance{}
	json.NewDecoder(resp.Body).Decode(&instance)
	defer resp.Body.Close()
	return &instance, nil
}

func (org *Organizations) DeleteInstance(orgName, dbName, instanceName string) error {
	if orgName == "" {
		return fmt.Errorf("organization name is required")
	}
	if dbName == "" {
		return fmt.Errorf("database name is required")
	}
	if instanceName == "" {
		return fmt.Errorf("instance name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases/%s/instances/%s", tursoBaseURL, orgName, dbName, instanceName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodDelete, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
