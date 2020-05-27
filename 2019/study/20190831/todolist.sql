create database todolist default utf8mb4;
create database testgorm default utf8mb4;
use database todolist;
create table task (
    id int,
    name varchar(64),
    `desc` varchar(1024),
    status int,
    create_time datetime,
    complete_time datetime
) engine=innodb default charset utf8mb4;



create table user (
    id int,
    name varchar(64),
    password varchar(1024),
    `desc`text,
    sex boolean,
    height float,
    birthday date,
    create_time datetime
) engine=innodb default charset utf8mb4;

create table user2 (
    id int,
    name varchar(64) not null default "",
    password varchar(1024),
    `desc`varchar(1024) not null default "",
    sex boolean not null default false,
    height float not null default 0.0,
    birthday date,
    create_time datetime
) engine=innodb default charset utf8mb4;

create table user3 (
    id int primary key,
    name varchar(64) not null default "",
    password varchar(1024),
    `desc`varchar(1024) not null default "",
    sex boolean not null default false,
    height float not null default 0.0,
    birthday date,
    create_time datetime
) engine=innodb default charset utf8mb4;


create table user4 (
    id int primary key,
    name varchar(64) unique not null default "",
    password varchar(1024),
    `desc`varchar(1024) not null default "",
    sex boolean not null default false,
    height float not null default 0.0,
    birthday date,
    create_time datetime
) engine=innodb default charset utf8mb4;

create table user5 (
    id int primary key auto_increment,
    name varchar(64) unique not null default "",
    password varchar(1024),
    `desc`varchar(1024) not null default "",
    sex boolean not null default false,
    height float not null default 0.0,
    birthday date,
    create_time datetime
) engine=innodb default charset utf8mb4;

create table test01 (
    col1 varchar(10),
    col2 varchar(10) not null default ""
) engine=innodb default charset utf8mb4;

insert into `tablename` value ('col1',....,'coln')
insert into test01 value ("1", "2");

insert into `tablename`('col1',....,'coln') values('col1',.....,'coln');
insert into test01(col2) values ("2");
insert into test01(col1) values ("1");
insert into test01(col2, col1) values ("2", "1");


select * from test01;

create table test02 (
    id int primary key,
    name VARCHAR(32) UNIQUE not null DEFAULT ''
) engine=innodb default charset utf8mb4;

INSERT INTO test02 values(1, "Jevon");
insert into test02(id) values (2);
select * from test02;


create table test03 (
    id int primary key auto_increment,
    name VARCHAR(32) UNIQUE not null DEFAULT ''
) engine=innodb default charset utf8mb4;

INSERT INTO test03 value (1, "Jevon");
INSERT INTO test03(name) values("Dan");
INSERT INTO test03 value (10, "Jevon10");
INSERT INTO test03(name) values("Dan11");
select * from test03;


create table users (
    id int primary key auto_increment,
    name varchar(64) unique not null default "",
    password varchar(1024) not null default "",
    sex boolean not null default false,
    birthday date,
    tel varchar(32) not null default "",
    addr varchar(128) not null default "",
    `desc` text,
    create_time datetime not null
) engine=innodb default charset utf8mb4;

INSERT INTO users value (1, "dan", "123", true, "2010-01-01", "10000000", "shanghai","abcd", "2014-01-01");
INSERT INTO users value (2, "ran", "123", false, "2010-02-01", "10000000", "shanghai","abcd12", "2010-01-01");
INSERT INTO users value (3, "wei", "abc", true, "2010-01-01", "10000000", "shanghai","abcd", "2011-01-01");
INSERT INTO users value (4, "jevon", "abc", false, "2010-04-01", "10000000", "shanghai","abcd1", "2010-01-01");
INSERT INTO users value (5, "jevonwei", "abc", false, "2010-04-01", "10000000", "shanghai","abcd1", now());
INSERT INTO users(name, password, sex, tel, addr, create_time) values ("danran", md5("abc"), false, "10000000", "shanghai",now());
INSERT INTO users(name, password, sex, birthday, tel, addr, create_time) values ("danran", md5("abc"), false, "2018-01-01", "10000000", "shanghai",now());

INSERT INTO users(name, password, sex, tel, addr, create_time) values 
("danran00", md5("abc"), false, "10000000", "shanghai",now()),
("danran01", md5("abc"), false, "10000000", "shanghai",now()),
("danran02", md5("abc"), false, "10000000", "shanghai",now());

create table tasks (
    id int primary key auto_increment,
    name varchar(64) not null default "",
    progress float not NULL DEFAULT 0,
    user VARCHAR(64) not NULL DEFAULT "",
    `desc` varchar(1024) not null default "",
    status int not NULL DEFAULT 0,
    create_time datetime not null,
    complete_time datetime
) engine=innodb default charset utf8mb4;


INSERT INTO tasks(name, progress, user, status, create_time, complete_time) values 
("danran01",0, "wei", 0, now(), "2019-09-20"),
("danran02",0, "dan", 0, now(), "2019-09-20"),
("danran03",0, "wei", 0, now(), "2019-09-20");

drop table users;

selece now();

select md5("a");
select name,sex, desc from users;
select name as n,sex as s,addr from users;

select count(*) from users;
select count(*) as a from users;
select count(*) from users as u;


select * from users where id < 3;
select * from users where name = "jevon";
select * from users where create_time > "2019-01-01";

select * from users where name in ("jevon", "danran");
select * from users where id BETWEEN 2 and 4;

SELECT * from users WHERE name like '%wei';
SELECT * from users WHERE name like 'dan%';
SELECT * from users WHERE name like '%an%';
SELECT * from users WHERE password like '%ab%';

SELECT * from users WHERE name like '%wei' and id > 3;
SELECT * from users WHERE name like '%wei' or id < 3;
SELECT * from users WHERE not id < 3;

SELECT * from test01 where col1 is not NULL;
SELECT * from test01 where not (col2 is NULL);

select * from users limit 2;
select * from users limit 5;
SELECT * from users WHERE name like '%an%' limit 2;
SELECT * from users WHERE name like '%an%' limit 2 offset 1;
SELECT * from users WHERE name like '%an%' limit 2 offset 2;
select * from users limit 2 offset 0;
select * from users limit 2 offset 2;
select * from users limit 2 offset 4;
SELECT * from users ORDER by name;
select * from users order by name;
select * from users order by name asc;
select * from users order by name desc;
select * from users order by birthday;
select * from users order by birthday asc, name desc;
select * from users where name like "%an%" ORDER by birthday desc limit 1 offset 1;
select count(*) from users where name like "%an%" ORDER by birthday desc;

select count(name) from users where name="jevon" and password = "abc";
select count(name) from users where name="jevon" and password = md5("abc");
select * from users where name="jevon" and password = md5("abc");


update tablename set colname = value1, colname = value2 where XXXX;
update users set addr="上海" where name = "ran";
update users set `desc`="test", sex = 0 where id > 4;
update users set birthday="2000-01-01", addr="henan" where id = 4;


delete from users  where id = 3;
delete from users  where id >6;
TRUNCATE TABLE users;


create table accesslog (
    id int primary key auto_increment,
    ip VARCHAR(64) not NULL DEFAULT "0.0.0.0",
    logtime datetime not null,
    method VARCHAR(8) not null DEFAULT "GET",
    url VARCHAR(1024) not null DEFAULT "",
    status_code int not NULL DEFAULT 200,
    bytes int Not null DEFAULT 0
) engine=innodb default charset utf8mb4;

INSERT INTO accesslog(ip, logtime, method, url, status_code, bytes) values 
("10.10.10.1", now(), "POST", "10.127.0.1/login", 302,122),
("10.10.10.2", now(), "POST", "/logout", 200,111),
("10.10.10.5", now(), "GET", "10.127.0.1/manager", 200,111),
("10.10.11.2", now(), "POST", "/", 200,111),
("10.10.10.11", now(), "GET", "/manager", 200,111),
("10.10.10.3", now(), "POST", "10.127.0.1/login", 302 ,100);

update accesslog set ip="100.1.1.1", status_code = 200 where id > 5 or id %3 = 0;
update accesslog set ip="200.2.2.2" where id > 3 and id %5 = 0;

delete from accesslog  where logtime <  "2019-08-31 15:22:35";

select ip,count(*) from accesslog group by ip;
select url,count(*) from accesslog group by ip;
select method,count(*) from accesslog group by method;

select ip,max(bytes) from accesslog group by ip;
select ip,min(bytes) from accesslog group by ip;
select ip,sum(bytes),max(bytes),min(bytes), avg(bytes),count(bytes) from accesslog group by ip;
select ip,url, sum(bytes),max(bytes),min(bytes), avg(bytes),count(bytes) from accesslog group by ip, url;

select ip,url, sum(bytes),max(bytes),min(bytes), avg(bytes),count(bytes) from accesslog group by ip, url having count(*) = 1;

select date_format(now(), "%Y-%M-%D");



