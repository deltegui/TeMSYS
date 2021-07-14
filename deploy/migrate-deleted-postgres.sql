-- For PostgreSQL

select * from sensors;
alter table sensors add column deleted_bool boolean not null default false;
update sensors as s set deleted_bool = (select (case when deleted = 1 then true else false end) from sensors where name = s.name);
alter table sensors drop column deleted;
alter table sensors rename column deleted_bool to deleted;