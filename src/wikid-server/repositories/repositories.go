package repositories

import "github.com/jmoiron/sqlx"

var _db *sqlx.DB

// Init initializes the repositories with an underlying database.
func Init(db *sqlx.DB) {
	_db = db
}
