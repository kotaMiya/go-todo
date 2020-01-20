package postgres

import (
	"github.com/go-pg/pg/v9"
	_ "github.com/lib/pg"
)

func New(opts *pg.Options) *pg.DB {
	db := pg.Connect(opts)

	return db
}
