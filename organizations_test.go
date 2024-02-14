package turso

import (
	"os"
	"testing"
)

var org_name string = os.Getenv("TURSO_ORG_NAME")
var db_name string = os.Getenv("TURSO_DB_NAME")
var instance_name string = os.Getenv("TURSO_INSTANCE_NAME")

func TestOrganizationList(t *testing.T) {
	client, err := newClient()
	if err != nil || client == nil {
		t.Error(err)
	}
	orgs, err := client.Organizations.List()
	if err != nil {
		t.Error(err)
	}
	if orgs == nil && len(orgs.Orgs) == 0 {
		t.Error("organizations should not be nil")
	}
}

func TestOrganizationMembers(t *testing.T) {
	client, err := newClient()
	if err != nil || client == nil {
		t.Error(err)
	}
	members, err := client.Organizations.Members(org_name)
	if err != nil {
		t.Error(err)
	}
	if members == nil && len(members.Members) == 0 {
		t.Error("members should not be nil")
	}
	members, err = client.Organizations.Members("")
	if err.Error() != "organization slug is required" {
		t.Error(err)
	}
}

func TestOrganizationDatabases(t *testing.T) {
	client, err := newClient()
	if err != nil || client == nil {
		t.Error(err)
	}
	databases, err := client.Organizations.Databases(org_name)
	if err != nil {
		t.Error(err)
	}
	if databases == nil && len(databases.Databases) == 0 {
		t.Error("databases should not be nil")
	}
	databases, err = client.Organizations.Databases("")
	if err.Error() != "organization slug is required" {
		t.Error(err)
	}
}

func TestOrganizationDatabase(t *testing.T) {
	client, err := newClient()
	if err != nil || client == nil {
		t.Error(err)
	}
	database, err := client.Organizations.Database(org_name, db_name)
	if err != nil {
		t.Error(err)
	}
	if database == nil && database.Database.Name != db_name {
		t.Error("databases should not be nil")
	}
	database, err = client.Organizations.Database("", db_name)
	if err.Error() != "organization slug is required" {
		t.Error(err)
	}
	database, err = client.Organizations.Database(org_name, "")
	if err.Error() != "database name is required" {
		t.Error(err)
	}
}

func TestOrganizationInstances(t *testing.T) {
	client, err := newClient()
	if err != nil || client == nil {
		t.Error(err)
	}
	instances, err := client.Organizations.Instances(org_name, db_name)
	if err != nil {
		t.Error(err)
	}
	if instances == nil && len(instances.Instances) == 0 {
		t.Error("instaces should not be nil")
	}
	instances, err = client.Organizations.Instances("", db_name)
	if err.Error() != "organization slug is required" {
		t.Error(err)
	}
	instances, err = client.Organizations.Instances(org_name, "")
	if err.Error() != "database name is required" {
		t.Error(err)
	}
}

func TestOrganizationInstance(t *testing.T) {
	client, err := newClient()
	if err != nil || client == nil {
		t.Error(err)
	}
	instances, err := client.Organizations.Instance(org_name, db_name, instance_name)
	if err != nil {
		t.Error(err)
	}
	if instances == nil && instances.Instance.Name != instance_name {
		t.Error("instances should not be nil")
	}
	instances, err = client.Organizations.Instance("", db_name, instance_name)
	if err.Error() != "organization slug is required" {
		t.Error(err)
	}
	instances, err = client.Organizations.Instance(org_name, "", instance_name)
	if err.Error() != "database name is required" {
		t.Error(err)
	}
	instances, err = client.Organizations.Instance(org_name, db_name, "")
	if err.Error() != "instance name is required" {
		t.Error(err)
	}
}

func TestOrganizationDBUsage(t *testing.T) {
	client, err := newClient()
	if err != nil || client == nil {
		t.Error(err)
	}
	usage, err := client.Organizations.DBUsage(org_name, db_name)
	if err != nil {
		t.Error(err)
	}
	if usage == nil {
		t.Error("instaces should not be nil")
	}
	if usage.Database.UUID == "" {
		t.Error("uuid should not be empty")
	}
	usage, err = client.Organizations.DBUsage("", db_name)
	if err.Error() != "organization slug is required" {
		t.Error(err)
	}
	usage, err = client.Organizations.DBUsage(org_name, "")
	if err.Error() != "database name is required" {
		t.Error(err)
	}
}

func TestOrganizationInvitesList(t *testing.T) {
	client, err := newClient()
	if err != nil || client == nil {
		t.Error(err)
	}
	invites, err := client.Organizations.ListInvites(org_name)
	if err != nil {
		t.Error(err)
	}
	if invites == nil && len(invites.Invites) == 0 {
		t.Error("invites should not be nil")
	}
	invites, err = client.Organizations.ListInvites("")
	if err.Error() != "organization slug is required" {
		t.Error(err)
	}
}

func TestOrganizationGroups(t *testing.T) {
	client, err := newClient()
	if err != nil || client == nil {
		t.Error(err)
	}
	groups, err := client.Organizations.ListGroups(org_name)
	if err != nil {
		t.Error(err)
	}
	if groups == nil && len(groups.Groups) == 0 {
		t.Error("groups should not be nil")
	}
	groups, err = client.Organizations.ListGroups("")
	if err.Error() != "organization slug is required" {
		t.Error(err)
	}
}

