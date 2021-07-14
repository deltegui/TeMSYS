-- For MySQL
use temsys;
show tables;
describe reports;

insert into temsys.sensors (name, conntype, connvalue, update_interval, deleted)
select NAME, CONNTYPE, CONNVALUE, UPDATE_INTERVAL, DELETED
FROM tempanalizr.SENSORS;

select * from temsys.sensors;

insert into temsys.report_types (type_name)
select TYPE_NAME from tempanalizr.REPORT_TYPES;

select * from temsys.report_types;

insert into temsys.used_report_types (sensor, report_type, add_date)
select SENSOR, REPORT_TYPE, ADD_DATE from tempanalizr.USED_REPORT_TYPES;

select * from temsys.used_report_types;

insert into temsys.reports (id, sensor, type, value, report_date)
select ID, SENSOR, TYPE, VALUE, REPORT_DATE from tempanalizr.REPORTS;

select count(*) from temsys.reports;