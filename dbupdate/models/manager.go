package models

type Client struct{
	ID int
	Name string
	Surname string
	NumberPhone string
	Login string
	Password string
	Locked bool
}

type ClientList struct {
}