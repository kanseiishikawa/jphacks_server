package database

import (
	_ "github.com/lib/pq"
	"github.com/gocraft/dbr"
)

func Plan_Generate( id int, key string, db *dbr.Connection ) error {
	_, err := db.Exec(
		`INSERT INTO plan_data ( "id", "key", "answer_count" ) VALUES ( $1, $2, $3 )`,
		id,
		key,
		0, )

	return err
}
