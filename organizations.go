package turso

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Organizations struct {
	client *client
}

type Organization struct {
	Name          string `json:"name"`
	Slug          string `json:"slug"`
	Type          string `json:"type"`
	BlockedReads  bool   `json:"blocked_reads"`
	BlockedWrites bool   `json:"blocked_writes"`
	Overages      bool   `json:"overages"`
}

type OrganizationMembers struct {
	Role     string `json:"role"`
	Username string `json:"username"`
}

type OrganisationList struct {
	Orgs []Organization `json:"organizations"`
}

type OrganizationGroup struct {
	Name      string   `json:"name"`
	Primary   string   `json:"primary"`
	UUID      string   `json:"uuid"`
	Archived  bool     `json:"archived"`
	Locations []string `json:"locations"`
}

type OrganizationInvite struct {
	Accepted       bool         `json:"Accepted"`
	CreatedAt      string       `json:"CreatedAt"`
	DeletedAt      string       `json:"DeletedAt"`
	UpdatedAt      string       `json:"UpdatedAt"`
	Email          string       `json:"Email"`
	ID             int          `json:"Id"`
	Organization   Organization `json:"Organization"`
	OrganizationID int          `json:"OrganizationID"`
	Role           string       `json:"Role"`
	Token          string       `json:"Token"`
}

type OrganizationInvites struct {
	Invites []OrganizationInvite `json:"invites"`
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

type topQueries struct {
	Query       string `json:"query"`
	RowsRead    int    `json:"rows_read"`
	RowsWritten int    `json:"rows_written"`
}

type DatabaseStats struct {
	TopQueries []topQueries `json:"top_queries"`
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

type organizationGroup struct {
	Group OrganizationGroup `json:"group"`
}

type organizationGroupList struct {
	Groups []OrganizationGroup `json:"groups"`
}

type organizationDatabaseList struct {
	Databases []Database `json:"databases"`
}

type organizationDatabase struct {
	Database struct {
		Database
	} `json:"database"`
}

type organizationDatabaseConfig struct {
	AllowAttach bool   `json:"allow_attach"`
	SizeLimit   string `json:"size_limit"`
}

type DatabaseConfiguration struct {
	organizationDatabaseConfig
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

func (org *Organizations) Update(organizationSlug string, body map[string]string) error {
	if organizationSlug == "" {
		return fmt.Errorf("organization slug is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s", tursoBaseURL, organizationSlug)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodPut, b)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (org *Organizations) Members(organizationSlug string) (*organizationMembersList, error) {
	if organizationSlug == "" {
		return nil, fmt.Errorf("organization slug is required")
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

func (org *Organizations) AddMembers(organizationSlug string, body map[string]string) error {
	if organizationSlug == "" {
		return fmt.Errorf("organization slug is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/members", tursoBaseURL, organizationSlug)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodPost, b)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (org *Organizations) RemoveMembers(organizationSlug string, username string) error {
	if organizationSlug == "" {
		return fmt.Errorf("organization slug is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/members/%s", tursoBaseURL, organizationSlug, username)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodDelete, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (org *Organizations) MintToken(organizationSlug, dbName, expiration, authorization string) (*jwtToken, error) {
	if organizationSlug == "" {
		return nil, fmt.Errorf("organization slug is required")
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

func (org *Organizations) CreateDatabase(orgSlug string, body map[string]string) (*Database, error) {
	if orgSlug == "" {
		return nil, fmt.Errorf("organization slug is required")
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases", tursoBaseURL, orgSlug)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodPost, b)
	if err != nil {
		return nil, err
	}
	var database = Database{}
	json.NewDecoder(resp.Body).Decode(&database)
	defer resp.Body.Close()
	return &database, nil
}

func (org *Organizations) UpdateDatabaseConfiguration(orgSlug, dbName string, body map[string]string) (*DatabaseConfiguration, error) {
	if orgSlug == "" {
		return nil, fmt.Errorf("organization slug is required")
	}
	if dbName == "" {
		return nil, fmt.Errorf("database name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases/%s/configuration", tursoBaseURL, orgSlug, dbName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodPatch, body)
	if err != nil {
		return nil, err
	}
	var database = DatabaseConfiguration{}
	json.NewDecoder(resp.Body).Decode(&database)
	defer resp.Body.Close()
	return &database, nil
}

func (org *Organizations) DeleteDatabase(orgSlug, dbName string) error {
	if orgSlug == "" {
		return fmt.Errorf("organization slug is required")
	}
	if dbName == "" {
		return fmt.Errorf("database name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases/%s", tursoBaseURL, orgSlug, dbName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodDelete, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (org *Organizations) UpdateDatabasesInGroup(orgSlug, groupName string) error {
	if orgSlug == "" {
		return fmt.Errorf("organization slug is required")
	}
	if groupName == "" {
		return fmt.Errorf("group name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/groups/%s/update", tursoBaseURL, orgSlug, groupName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodPut, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (org *Organizations) UpdateAllInstances(orgSlug, dbName string) error {
	if orgSlug == "" {
		return fmt.Errorf("organization slug is required")
	}
	if dbName == "" {
		return fmt.Errorf("database name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases/%s/update", tursoBaseURL, orgSlug, dbName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodPut, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (org *Organizations) DBUsage(orgSlug, dbName string) (*DBMonthlyUsage, error) {
	if orgSlug == "" {
		return nil, fmt.Errorf("organization slug is required")
	}
	if dbName == "" {
		return nil, fmt.Errorf("database name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases/%s/usage", tursoBaseURL, orgSlug, dbName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var usage = DBMonthlyUsage{}
	json.NewDecoder(resp.Body).Decode(&usage)
	defer resp.Body.Close()
	return &usage, nil
}

func (org *Organizations) Instances(orgSlug, dbName string) (*databaseInstances, error) {
	if orgSlug == "" {
		return nil, fmt.Errorf("organization slug is required")
	}
	if dbName == "" {
		return nil, fmt.Errorf("database name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases/%s/instances", tursoBaseURL, orgSlug, dbName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var instances = databaseInstances{}
	json.NewDecoder(resp.Body).Decode(&instances)
	defer resp.Body.Close()
	return &instances, nil
}

func (org *Organizations) Instance(orgSlug, dbName, instanceName string) (*databaseInstance, error) {
	if orgSlug == "" {
		return nil, fmt.Errorf("organization slug is required")
	}
	if dbName == "" {
		return nil, fmt.Errorf("database name is required")
	}
	if instanceName == "" {
		return nil, fmt.Errorf("instance name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases/%s/instances/%s", tursoBaseURL, orgSlug, dbName, instanceName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var instance = databaseInstance{}
	json.NewDecoder(resp.Body).Decode(&instance)
	defer resp.Body.Close()
	return &instance, nil
}

func (org *Organizations) CreateInstance(orgSlug, dbName string, body map[string]string) (*databaseInstance, error) {
	if orgSlug == "" {
		return nil, fmt.Errorf("organization slug is required")
	}
	if dbName == "" {
		return nil, fmt.Errorf("database name is required")
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases/%s/instances", tursoBaseURL, orgSlug, dbName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodPost, b)
	if err != nil {
		return nil, err
	}
	var instance = databaseInstance{}
	json.NewDecoder(resp.Body).Decode(&instance)
	defer resp.Body.Close()
	return &instance, nil
}

func (org *Organizations) DeleteInstance(orgSlug, dbName, instanceName string) error {
	if orgSlug == "" {
		return fmt.Errorf("organization slug is required")
	}
	if dbName == "" {
		return fmt.Errorf("database name is required")
	}
	if instanceName == "" {
		return fmt.Errorf("instance name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases/%s/instances/%s", tursoBaseURL, orgSlug, dbName, instanceName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodDelete, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (org *Organizations) ListGroups(orgSlug string) (*organizationGroupList, error) {
	if orgSlug == "" {
		return nil, fmt.Errorf("organization slug is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/groups", tursoBaseURL, orgSlug)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var groups = organizationGroupList{}
	json.NewDecoder(resp.Body).Decode(&groups)
	defer resp.Body.Close()
	return &groups, nil
}

func (org *Organizations) Group(orgSlug, groupName string) (*organizationGroup, error) {
	if orgSlug == "" {
		return nil, fmt.Errorf("organization slug is required")
	}
	if groupName == "" {
		return nil, fmt.Errorf("group name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/groups/%s", tursoBaseURL, orgSlug, groupName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var group = organizationGroup{}
	json.NewDecoder(resp.Body).Decode(&group)
	defer resp.Body.Close()
	return &group, nil
}

func (org *Organizations) CreateGroup(orgSlug string, body map[string]string) (*organizationGroup, error) {
	if orgSlug == "" {
		return nil, fmt.Errorf("organization slug is required")
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/groups", tursoBaseURL, orgSlug)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodPost, b)
	if err != nil {
		return nil, err
	}
	var group = organizationGroup{}
	json.NewDecoder(resp.Body).Decode(&group)
	defer resp.Body.Close()
	return &group, nil
}

func (org *Organizations) DeleteGroup(orgSlug, groupName string) error {
	if orgSlug == "" {
		return fmt.Errorf("organization slug is required")
	}
	if groupName == "" {
		return fmt.Errorf("group name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/groups/%s", tursoBaseURL, orgSlug, groupName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodDelete, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (org *Organizations) AddLocationToGroup(orgSlug, groupName, location string) (*organizationGroup, error) {
	if orgSlug == "" {
		return nil, fmt.Errorf("organization slug is required")
	}
	if groupName == "" {
		return nil, fmt.Errorf("group name is required")
	}
	if location == "" {
		return nil, fmt.Errorf("location is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/groups/%s/locations/%s", tursoBaseURL, orgSlug, groupName, location)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodPost, nil)
	if err != nil {
		return nil, err
	}
	var group = organizationGroup{}
	json.NewDecoder(resp.Body).Decode(&group)
	defer resp.Body.Close()
	return &group, nil
}

func (org *Organizations) RemoveLocationFromGroup(orgSlug, groupName, location string) (*organizationGroup, error) {
	if orgSlug == "" {
		return nil, fmt.Errorf("organization slug is required")
	}
	if groupName == "" {
		return nil, fmt.Errorf("group name is required")
	}
	if location == "" {
		return nil, fmt.Errorf("location is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/groups/%s/locations/%s", tursoBaseURL, orgSlug, groupName, location)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodDelete, nil)
	if err != nil {
		return nil, err
	}
	var group = organizationGroup{}
	json.NewDecoder(resp.Body).Decode(&group)
	defer resp.Body.Close()
	return &group, nil
}

func (org *Organizations) UploadDumpFile(orgSlug string, file io.Reader) error {
	if orgSlug == "" {
		return fmt.Errorf("organization slug is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases/dumps", tursoBaseURL, orgSlug)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodPost, file)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (org *Organizations) InvalidateAllDBTokens(orgSlug, dbName string) error {
	if orgSlug == "" {
		return fmt.Errorf("organization slug is required")
	}
	if dbName == "" {
		return fmt.Errorf("database name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases/%s/auth/rotate", tursoBaseURL, orgSlug, dbName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodPost, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (org *Organizations) InvalidateAllGroupTokens(orgSlug, groupName, token string) error {
	if orgSlug == "" {
		return fmt.Errorf("organization slug is required")
	}
	if groupName == "" {
		return fmt.Errorf("group name is required")
	}
	if token == "" {
		return fmt.Errorf("token is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/groups/%s/auth/rotate", tursoBaseURL, orgSlug, groupName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodPost, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (org *Organizations) ListInvites(orgSlug string) (*OrganizationInvites, error) {
	if orgSlug == "" {
		return nil, fmt.Errorf("organization slug is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/invites", tursoBaseURL, orgSlug)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var invites = OrganizationInvites{}
	json.NewDecoder(resp.Body).Decode(&invites)
	defer resp.Body.Close()
	return &invites, nil
}

func (org *Organizations) CreateInvite(orgSlug string, body map[string]string) (*OrganizationInvite, error) {
	if orgSlug == "" {
		return nil, fmt.Errorf("organization slug is required")
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/invites", tursoBaseURL, orgSlug)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodPost, b)
	if err != nil {
		return nil, err
	}
	var invite = OrganizationInvite{}
	json.NewDecoder(resp.Body).Decode(&invite)
	defer resp.Body.Close()
	return &invite, nil
}

func (org *Organizations) TransferOrganisation(orgSlug, groupName, ToOrgSlug string) (*organizationGroup, error) {
	if orgSlug == "" {
		return nil, fmt.Errorf("organization slug is required")
	}
	if groupName == "" {
		return nil, fmt.Errorf("group name is required")
	}
	if ToOrgSlug == "" {
		return nil, fmt.Errorf("the organization slug to be transfer to is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/groups/%s/transfer", tursoBaseURL, orgSlug, groupName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodPost, nil)
	if err != nil {
		return nil, err
	}
	var transfer = organizationGroup{}
	json.NewDecoder(resp.Body).Decode(&transfer)
	defer resp.Body.Close()
	return &transfer, nil
}

func (org *Organizations) DatabaseStats(orgSlug, dbName string) (*DatabaseStats, error) {
	if orgSlug == "" {
		return nil, fmt.Errorf("organization slug is required")
	}
	if dbName == "" {
		return nil, fmt.Errorf("database name is required")
	}
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/databases/%s/stats", tursoBaseURL, orgSlug, dbName)
	resp, err := org.client.tursoAPIrequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var stats = DatabaseStats{}
	json.NewDecoder(resp.Body).Decode(&stats)
	defer resp.Body.Close()
	return &stats, nil
}
