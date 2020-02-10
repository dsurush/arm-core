package dbupdate

const addClientDML = `insert into clients(name, surname, login, password, phone, locked) 
values ($1, $2, $3, $4, $5, $6)`

const addATM = `insert into ATMs(name, locked)
values ($1, $2)`

const addAccount  = `insert into accounts(user_id, name, accountNumber, locked) 
values (?, ?, ?, ?)`

const addService = `insert into services(name, price)
values (?, ?)`