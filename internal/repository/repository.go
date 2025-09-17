package repository

import (
	"github.com/jmoiron/sqlx"
)

type RepositoryStruct struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *RepositoryStruct {
	return &RepositoryStruct{
		db: db,
	}
}
