package models

type Account struct {
	ID int64
	UserId int64
	Name string
	AccountNumber int64
	Locked bool
}
