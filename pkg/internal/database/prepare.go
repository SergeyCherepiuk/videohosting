package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func MustPrepareNamed(db *sqlx.DB, query string) *sqlx.NamedStmt {
	stmt, err := db.PrepareNamed(query)
	if err != nil {
		log.Fatal(err)
	}
	return stmt
}