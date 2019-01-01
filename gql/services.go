package gql

import (
	"github.com/Focinfi/robo-api/gql/gen"
	"github.com/Focinfi/robo-api/models/gql"
	"github.com/Focinfi/robo-api/services/pipelines"
)

const (
	timeLayout = "2006-01-02 15:04:05"
)

func buildPipelineForService(p pipelines.Pipeline) gql.Pipeline {
	return gql.Pipeline{
		ID:        p.UUID,
		Desc:      p.Desc,
		Confs:     string(p.Confs),
		CreatedAt: p.CreatedAt.Format(timeLayout),
		UpdatedAt: p.UpdatedAt.Format(timeLayout),
	}
}

func buildMapFromKV(kvs []gen.KV) map[string]interface{} {
	m := make(map[string]interface{}, len(kvs))
	for _, kv := range kvs {
		m[kv.Key] = kv.Value
	}
	return m
}
