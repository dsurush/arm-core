package cmodels

type Client struct{
	ID int
	Name string
	Surname string
	Login string
	Password string
	NumberPhone string
	Locked bool
}

type ClientList struct {
	Clients []Client
}