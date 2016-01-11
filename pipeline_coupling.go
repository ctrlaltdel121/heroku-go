package heroku

import "time"

type PipelineCoupling struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	App       struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"app"`
	Pipeline struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"pipeline"`
	Stage       string  `json:"stage"`
	PullRequest *string `json:"pull_request"`
}

func (c *Client) PipelineCouplingCreate(appID, pipelineID, stage string) (*PipelineCoupling, error) {
	var pc PipelineCoupling
	var pco PipelineCouplingCreateOpts
	pco.Stage = stage
	pco.Pipeline.ID = pipelineID
	return &pc, c.Post(&pc, "/apps/"+appID+"/pipeline-couplings", &pco)
}

type PipelineCouplingCreateOpts struct {
	Stage    string `json:"stage"`
	Pipeline struct {
		ID string `json:"id"`
	} `json:"pipeline"`
}
