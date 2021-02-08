package mysql

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type sqlRepository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) sqlRepository {
	return sqlRepository{db}
}

func (repo sqlRepository) beginOrFatal() *sqlx.Tx {
	tx, err := repo.db.Beginx()
	if err != nil {
		log.Fatal(err)
	}
	return tx
}
