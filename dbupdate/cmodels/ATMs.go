package cmodels

type ATM struct {
	ID int64
	Name string
	Locked bool
}

type AtmList struct {
	ATMs []ATM
}

type CreateNewATM struct {
	Address string
	Locked bool
}