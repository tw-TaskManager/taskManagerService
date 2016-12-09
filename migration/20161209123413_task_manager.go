package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up_20161209123413, Down_20161209123413)
}

func Up_20161209123413(tx *sql.Tx) error {
	return nil
}

func Down_20161209123413(tx *sql.Tx) error {
	return nil
}
