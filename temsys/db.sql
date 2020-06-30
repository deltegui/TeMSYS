create database temsys;
use temsys;

create table users (
	username varchar(255) primary key,
    password varchar(255) not null,
    role varchar(255) not null
);