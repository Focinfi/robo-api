package pipelines

import (
	"context"

	"time"

	"fmt"

	"github.com/Focinfi/pipeline"
	"github.com/Focinfi/robo-api/dao"
)

var pipelines *Pipelines

func All() *Pipelines {
	return pipelines
}

func Add(ctx context.Context, newPipeline NewPipeline) error {
	dbPipeline := &dao.Pipeline{
		UUID:      newPipeline.UUID,
		Desc:      newPipeline.Desc,
		Confs:     newPipeline.Confs,
		CreatedAt: time.Now(),
	}
	// try add
	if err := pipelines.AddFromDB(dbPipeline); err != nil {
		return fmt.Errorf("add pipeline failed, err: %v", err)
	}

	// try set into db
	if err := dao.InsertPipeline(ctx, dbPipeline); err != nil {
		return fmt.Errorf("store pipeline failed, err: %v", err)
	}

	return nil
}

func init() {
	if ps, err := MakePipelines(); err != nil {
		panic("init pipelines failed, err: " + err.Error())
	} else {
		pipelines = ps
	}
}

func MakePipelines() (*Pipelines, error) {
	// get from db
	dbPipelines, err := dao.GetAllPipelines(context.Background())
	if err != nil {
		panic(err)
	}

	return MakePipelinesFromDB(dbPipelines)
}

func MakePipelinesFromDB(dbPipelines []dao.Pipeline) (*Pipelines, error) {
	// init ps
	ps := &Pipelines{
		Lines:       make([]Pipeline, 0, len(dbPipelines)),
		lineMap:     make(map[string]Pipeline, len(dbPipelines)),
		handlersMap: make(map[string]pipeline.Handler, len(dbPipelines)),
	}
	for _, dbP := range dbPipelines {
		if err := ps.AddFromDB(&dbP); err != nil {
			return nil, err
		}
	}
	return ps, nil
}
