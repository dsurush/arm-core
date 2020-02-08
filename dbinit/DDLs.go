package dbinit

const clientsDDL  = `create table clients (
    id integer primary key autoincrement,
    name text not null,
    surname text not null,
    login text not null unique,
    password text not null check ( length(password) > 4 ),
    locked boolean not null
);`

const accountsDDL = `create table accounts (
    id integer primary key autoincrement,
    user_id integer not null references clients,
    name text not null,
    locked boolean not null
);`

const ATMsDDL  = `create table ATMs (
    id integer primary key autoincrement,
    name text not null,
    locked boolean not null
);`

const servicesDDL =  `create table services (
    id integer primary key autoincrement,
    name text not null,
    price boolean not null
);`