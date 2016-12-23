package app

import (
	"sync"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var dbOnce sync.Once

// GetDatabase panics if it fails get the database.
func GetDatabase() *sqlx.DB {
	dbOnce.Do(func() {
		_db, err := sqlx.Open("mysql", GetConfig().DBConnStr)
		if err != nil {
			panic(err)
		}

		db = _db
	})

	return db
}
