package dao

import (
	"time"
)

type Pipeline struct {
	ID        int64     `db:"id"`
	UUID      string    `db:"uuid"`
	Desc      string    `db:"desc"`
	Confs     string    `db:"confs"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
