package turso

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type auditLog struct {
	Author    string                 `json:"author"`
	Code      string                 `json:"code"`
	CreatedAt string                 `json:"created_at"`
	Data      map[string]interface{} `json:"data"`
	Message   string                 `json:"message"`
	Origin    string                 `json:"origin"`
}

type pagination struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalPages int `json:"total_pages"`
	TotalRows  int `json:"total_rows"`
}

type AuditLogs struct {
	client     *client
	AuditLogs  []auditLog `json:"audit_logs"`
	Pagination pagination `json:"pagination"`
}

func (t *AuditLogs) List(orgSlug string) (*AuditLogs, error) {
	endpoint := fmt.Sprintf("%s/v1/organizations/%s/audit-logs", tursoBaseURL, orgSlug)
	resp, err := t.client.tursoAPIrequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var auditLogs AuditLogs
	if err := json.NewDecoder(resp.Body).Decode(&auditLogs); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return &auditLogs, nil
}
