package tender

import (
	_ "github.com/lib/pq" // PostgreSQL driver
)

type Tender struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	ServiceType     string `json:"service_type"`
	Status          string `json:"status"`
	OrganizationID  int    `json:"organization_id"`
	CreatorUsername string `json:"creator_username"`
}
