create database if not exists temsys;
use temsys;

drop table if exists users;

create table users(
    name varchar(255) primary key,
    password varchar(255) not null,
    role varchar(255) not null
);

create table report_types (
    type_name varchar(255) primary key
);

create table sensors (
	name varchar(255) primary key,
	conntype varchar(255) not null,
	connvalue varchar(255) not null,
    update_interval integer not null,
	deleted boolean not null default 0
);

create table used_report_types (
	sensor varchar(255) not null,
    report_type varchar(255) not null,
	add_date datetime(3)
);

create table reports (
	id integer primary key auto_increment,
	sensor varchar(255),
	type varchar(255) not null,
	value float not null,
	report_date datetime(3)
);

alter table reports add constraint foreign key (sensor) references sensors(name);
alter table reports add constraint foreign key (type) references report_types(type_name);
alter table used_report_types add constraint primary key (sensor, report_type);
alter table used_report_types add constraint foreign key (sensor) references sensors(name);
alter table used_report_types add constraint foreign key (report_type) references report_types(type_name);