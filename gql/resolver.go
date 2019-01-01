package gql

import (
	"context"
	"errors"
	"fmt"

	"github.com/Focinfi/pipeline"
	"github.com/Focinfi/robo-api/gql/gen"
	"github.com/Focinfi/robo-api/models/gql"
	"github.com/Focinfi/robo-api/services/pipelines"
)

var ErrPipelineNotFound = errors.New("pipeline not found")

type Resolver struct{}

func (r *Resolver) Builder() gen.BuilderResolver {
	return &builderResolver{r}
}

func (r *Resolver) Mutation() gen.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() gen.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) AddPipeline(ctx context.Context, input gen.NewPipeline) (gql.Pipeline, error) {
	if err := pipelines.Add(context.Background(), pipelines.NewPipeline{
		UUID:  input.ID,
		Desc:  input.Desc,
		Confs: input.Confs,
	}); err != nil {
		return gql.Pipeline{}, err
	}

	p, ok := pipelines.All().GetOK(input.ID)
	if !ok {
		return gql.Pipeline{}, ErrPipelineNotFound
	}
	return buildPipelineForService(p), nil
}

type builderResolver struct{ *Resolver }

func (r *builderResolver) Pipelines(ctx context.Context, obj *gql.Builder) ([]gql.Pipeline, error) {
	ps := make([]gql.Pipeline, 0, len(obj.PipelineIDs))
	for _, id := range obj.PipelineIDs {

		pipeline, ok := pipelines.All().GetOK(id)
		if !ok {
			return nil, ErrPipelineNotFound
		}
		ps = append(ps, buildPipelineForService(pipeline))
	}
	return ps, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Builders(ctx context.Context) ([]gql.Builder, error) {
	return Builders, nil
}

func (r *queryResolver) Pipelines(ctx context.Context) ([]gql.Pipeline, error) {
	all := pipelines.All()
	ps := make([]gql.Pipeline, 0, len(all.Lines))
	for _, p := range all.Lines {
		ps = append(ps, buildPipelineForService(p))
	}
	return ps, nil
}

func (r *queryResolver) RunPipeline(ctx context.Context, pipelineId string, input *gen.HandleInput) (string, error) {
	p, ok := pipelines.All().GetOK(pipelineId)
	if !ok {
		return "", ErrPipelineNotFound
	}

	args := pipeline.Args{}
	if input != nil {
		args = pipeline.Args{
			InValue: input.InValue,
			Params:  buildMapFromKV(input.Params),
		}
	}

	resp, err := p.Handle(ctx, args)

	if err != nil {
		return "", err
	}
	return fmt.Sprint(resp.OutValue), nil
}

func (r *queryResolver) RubPipelineVerbosely(ctx context.Context, pipelineId string, input *gen.HandleInput) (gen.VerboseResp, error) {
	p, ok := pipelines.All().GetOK(pipelineId)
	if !ok {
		return gen.VerboseResp{}, ErrPipelineNotFound
	}

	args := pipeline.Args{}
	if input != nil {
		args = pipeline.Args{
			InValue: input.InValue,
			Params:  buildMapFromKV(input.Params),
		}
	}

	resp, logs, err := p.HandleVerbosely(ctx, args)
	if err != nil {
		return gen.VerboseResp{}, err
	}

	genLogs := make([]gen.HandleLog, 0, len(logs))
	for _, log := range logs {
		errMsg := ""
		if log.Err != nil {
			errMsg = log.Err.Error()
		}
		genLogs = append(genLogs, gen.HandleLog{
			Log:    log.Out,
			ErrMsg: errMsg,
		})
	}

	return gen.VerboseResp{
		Result: fmt.Sprint(resp.OutValue),
		Logs:   genLogs,
	}, nil
}
