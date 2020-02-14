package dbupdate

const addClientDML = `insert into clients(name, surname, login, password, phone, locked) 
values ($1, $2, $3, $4, $5, $6)`

const addATM = `insert into ATMs(address, locked)
values ($1, $2)`

const addAccount  = `insert into accounts(user_id, name, accountNumber, locked) 
values (?, ?, ?, ?)`

const addService = `insert into services(name, price)
values (?, ?)`

const getAllClients = `select *from clients`

const getAllAccounts = `select *from accounts a left join clients c on a.user_id = c.id`

const getAllATMs = `select *from ATMs`

const loginSQL = `select login, password from clients where login = ?`

const searchClientByLogin = `select id, surname from clients where login = ?`

const searchAccountByIDSql = `select id, name, accountNumber, locked from accounts where user_id = ?`