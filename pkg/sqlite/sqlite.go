package sqlite

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	path string
}

func NewSqlite(path string) *Sqlite {
	return &Sqlite{path: path}
}

func (s *Sqlite) GetDB() (*sqlx.DB, error) {
	return sqlx.Open("sqlite3", s.path)
}
