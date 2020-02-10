package dbinit

const clientsDDL  = `create table if not exists clients (
    id integer primary key autoincrement,
    name text not null,
    surname text not null,
    login text not null unique,
    password text not null check ( length(password) > 4 ),
	phone text not null,
    locked boolean not null
);`

const accountsDDL = `create table if not exists accounts (
    id integer primary key autoincrement,
    user_id integer references clients,
    name text not null,
	accountNumber integer, 
    locked boolean not null
);`

const ATMsDDL  = `create table if not exists ATMs (
    id integer primary key autoincrement,
    name text not null,
    locked boolean not null
);`

const servicesDDL =  `create table if not exists services (
    id integer primary key autoincrement,
    name text not null,
    price int not null
);`
