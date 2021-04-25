package productstore

import (
	"github.com/jmoiron/sqlx"
)

type sqlStore struct {
	db *sqlx.DB
}

func NewSQLStore(db *sqlx.DB) *sqlStore {
	return &sqlStore{db: db}
}

