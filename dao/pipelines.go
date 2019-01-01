package dao

import "context"

func GetAllPipelines(ctx context.Context) ([]Pipeline, error) {
	var pipelines []Pipeline
	_, err := newSess().Select("*").From(tablePipeline).Load(&pipelines)
	return pipelines, err
}

func InsertPipeline(ctx context.Context, pipeline *Pipeline) error {
	_, err := newSess().InsertInto(tablePipeline).Columns("uuid", "desc", "confs", "created_at").Record(pipeline).Exec()
	return err
}
