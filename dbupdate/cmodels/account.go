package cmodels

type Account struct {
	ID int64
	UserId int64
	Name string
	AccountNumber int64
	Balance int64
	Locked bool
}

type AccountWithUserName struct {
	Account Account
	Client Client
}
type AccountList struct {
	AccountWithUserName []AccountWithUserName
}

type AccountForUser struct {
	ID int64
	Name string
	AccountNumber int64
	Balance int64
	Locked bool
}

//type AccountListByID struct {
//	Accounts []AccountForUser
//}