package pipelines

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Focinfi/pipeline"
	"github.com/Focinfi/pipeline/line"
	"github.com/Focinfi/robo-api/dao"
)

var ErrPipelineAlreadyExists = errors.New("pipeline already exists")

type NewPipeline struct {
	UUID  string `json:"id"`
	Desc  string `json:"desc"`
	Confs string `json:"confs"`
}

type Pipeline struct {
	ID        int64          `json:"-"`
	UUID      string         `json:"id"`
	Desc      string         `json:"desc"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Confs     string         `json:"-"`
	handler   *line.Handlers `json:"-"`
}

func (p Pipeline) Handle(ctx context.Context, args pipeline.Args) (resp *pipeline.Resp, err error) {
	return p.handler.Handle(ctx, args)
}

func (p Pipeline) HandleVerbosely(ctx context.Context, args pipeline.Args) (*pipeline.Resp, []*pipeline.RespLog, error) {
	return p.handler.HandleVerbosely(ctx, args)
}

type Pipelines struct {
	Lines       []Pipeline
	lineMap     map[string]Pipeline
	handlersMap map[string]pipeline.Handler
}

func (ps Pipelines) GetOK(uuid string) (Pipeline, bool) {
	p, ok := ps.lineMap[uuid]
	return p, ok
}

func (ps *Pipelines) AddFromDB(dbPipeline *dao.Pipeline) error {
	handler, err := line.NewHandlers(dbPipeline.UUID, dbPipeline.Confs, ps.handlersMap)
	if err != nil {
		return fmt.Errorf("build line.Handlers failed, pipeline.id: %v, err: %v", dbPipeline.UUID, err)
	}
	p := Pipeline{
		ID:        dbPipeline.ID,
		UUID:      dbPipeline.UUID,
		Desc:      dbPipeline.Desc,
		CreatedAt: dbPipeline.CreatedAt,
		UpdatedAt: dbPipeline.UpdatedAt,
		Confs:     dbPipeline.Confs,
		handler:   handler,
	}

	_, ok := ps.lineMap[dbPipeline.UUID]
	if ok {
		return ErrPipelineAlreadyExists
	}

	ps.Lines = append(ps.Lines, p)
	ps.lineMap[dbPipeline.UUID] = p
	ps.handlersMap[dbPipeline.UUID] = p
	return nil
}
