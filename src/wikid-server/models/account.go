package models

import "time"

type Account struct {
	Id          int64     `json:"id"          db:"id"`
	Email       string    `json:"email"       db:"email"`
	Password    []byte    `json:"password"    db:"password"`
	Salt        []byte    `json:"salt"        db:"salt"`
	CreatedTime time.Time `json:"createdTime" db:"created_time"`
	UpdatedTime time.Time `json:"updatedTime" db:"updated_time"`
	DeletedTime time.Time `json:"deletedTime" db:"deleted_time"`
}
