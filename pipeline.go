package heroku

import (
	"time"
)

type Pipeline struct {
	ID             string    `json:"id"`
	AccountID      string    `json:"-"`
	OrganizationID string    `json:"-"`
	Name           string    `json:"name"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (c *Client) PipelineInfo(name string) (*Pipeline, error) {
	var p Pipeline
	return &p, c.Get(&p, "/pipelines/"+name)
}
