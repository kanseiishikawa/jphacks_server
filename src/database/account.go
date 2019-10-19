package database

import (
	_ "github.com/lib/pq"
	"github.com/gocraft/dbr"
)

func Account_ID( sess *dbr.Session, account string, password string ) ( int, error ) {
	var ID int

	err :=  sess.QueryRow( `SELECT ID FROM account_info WHERE password = $1 AND name = $2`,
		password,
		account,).Scan( &ID )

	return ID, err
}


