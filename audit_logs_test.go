package turso

import (
	"testing"
)

func TestAuditLogs(t *testing.T) {
	client, err := newClient()
	if err != nil || client == nil {
		t.Error(err)
	}
	auditLogs, err := client.AuditLogs.List(org_name)
	if err != nil {
		t.Error(err)
	}
	if auditLogs == nil && len(auditLogs.AuditLogs) == 0 {
		t.Error("auditLogs should not be nil")
	}
}
