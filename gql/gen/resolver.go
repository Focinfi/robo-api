package gen

import (
	"context"

	"github.com/Focinfi/robo-api/models/gql"
)

type Resolver struct{}

func (r *Resolver) Builder() BuilderResolver {
	return &builderResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type builderResolver struct{ *Resolver }

func (r *builderResolver) Pipelines(ctx context.Context, obj *gql.Builder) ([]gql.Pipeline, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) AddPipeline(ctx context.Context, input NewPipeline) (gql.Pipeline, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Builders(ctx context.Context) ([]gql.Builder, error) {
	panic("not implemented")
}
func (r *queryResolver) Pipelines(ctx context.Context) ([]gql.Pipeline, error) {
	panic("not implemented")
}
func (r *queryResolver) RunPipeline(ctx context.Context, pipelineId string, input *HandleInput) (string, error) {
	panic("not implemented")
}
func (r *queryResolver) RubPipelineVerbosely(ctx context.Context, pipelineId string, input *HandleInput) (VerboseResp, error) {
	panic("not implemented")
}
