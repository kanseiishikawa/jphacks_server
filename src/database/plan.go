package database

import (
	_ "github.com/lib/pq"
	"github.com/gocraft/dbr"
)

func Plan_Insert( id int, key string, db *dbr.Connection ) error {
	_, err := db.Exec(
		`INSERT INTO plan_data ( "id", "key", "answer_count" ) VALUES ( $1, $2, $3 )`,
		id,
		key,
		0, )

	return err
}

func Plan_Answer( sess *dbr.Session, key string ) ( int , error ) {
	var count int

	err := sess.QueryRow( `SELECT answer_count FROM plan_data WHERE key = $1`,
		key, ).Scan( &count )

	return count, err
}

func Plan_Answer_Update( db *dbr.Connection, key string, count int) error {

	_, err := db.Exec( `UPDATE plan_data SET answer_count = $1 WHERE key = $2`,
		count,
		key,)

	return err
}
