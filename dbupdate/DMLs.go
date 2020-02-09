package dbupdate

const addClientDML = `insert into clients(name, surname, login, password, locked) 
values ($1, $2, $3, $4, $5)`

const addATM = `insert into ATMs(name, locked)
values ($1, $2)`