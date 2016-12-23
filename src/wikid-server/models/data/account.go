package data

import "time"

type Account struct {
	ID          string    `db:"id"`
	Email       string    `db:"email"`
	Password    []byte    `db:"password"`
	Salt        []byte    `db:"salt"`
	CreatedTime time.Time `db:"created_time"`
	UpdatedTime time.Time `db:"updated_time"`
	DeletedTime time.Time `db:"deleted_time"`
}
