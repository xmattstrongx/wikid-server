package models

type Session struct {
	Id        string `db:"id"`
	AccountId string `db:"account_id"`
}
