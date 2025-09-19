package repository

import (
	"database/sql"
	"errors"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jmoiron/sqlx"
	"myasd/internal/errs"
)

type RepositoryStruct struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *RepositoryStruct {
	return &RepositoryStruct{
		db: db,
	}
}

func (r *RepositoryStruct) translateError(err error) error {
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case pgerrcode.UniqueViolation:
			return errs.ErrDuplicateArticle
		case pgerrcode.NotNullViolation:
			return errs.ErrInvalidArticle
		case pgerrcode.ForeignKeyViolation:
			return errs.ErrUserNotExists
		}
	}

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return errs.ErrNotFound
	default:
		return err
	}
}
