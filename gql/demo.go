package gql

import (
	"github.com/Focinfi/pipeline/builders"
	"github.com/Focinfi/robo-api/models/gql"
)

var Builders = func() []gql.Builder {
	names := builders.GetBuilderNames()
	builders := make([]gql.Builder, len(names))
	for i, name := range names {
		builders[i] = gql.Builder{Name: name, PipelineIDs: []string{"demo_expr"}}
	}
	return builders
}()
