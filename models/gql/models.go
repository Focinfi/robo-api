package gql

type Builder struct {
	Name        string   `json:"name"`
	PipelineIDs []string `json:"pipeline_ids"`
}

type Pipeline struct {
	ID        string `json:"id"`
	Desc      string `json:"desc"`
	Confs     string `json:"confs"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
