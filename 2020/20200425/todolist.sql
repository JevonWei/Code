create table task (
    id int,
    name varchar(64),
    `desc` varchar(128),
    status int,
    create_time datetime,
    complate_time datetime
) engine=innodb default charset utf8mb4;

insert into task values(1, "dan", "xxx", 0, Now(), Now())

insert into task(id, name, `desc`, create_time, complate_time, status) values(2, "ran", "xxx", Now(), Now(), 0)

create table task (
    id int primary key auto_increment,
    name varchar(64) not null default '',
    progress float not null default 0,
    user varchar(64) not null default '',
    `desc` varchar(128) not null default '',
    status int not null default 0,
    create_time datetime not null,
    complate_time datetime
) engine=innodb default charset utf8mb4;

create table user (
    id int primary key auto_increment,
    name varchar(64) unique not null default '',
    password varchar(1024) not null default '',
    `desc` varchar(128) not null default '',
    tel varchar(128) not null default '',
    addr varchar(512) not null default '',
    super boolean not null default false
) engine=innodb default charset utf8mb4;

insert into user(name, password, `desc`, tel) values("ran", "123", "xxx", "1231213");

insert into user(name, password, `desc`, tel) values
    ("ran03", "123", "xxx", "1231213"),
    ("ran04", "123", "xxx", "1231213"),
    ("ran05", "123", "xxx", "1231213");

create table accesslog (
    id int primary key auto_increment,
    ip varchar(128) not null default "",
    logtime datetime not null,
    method varchar(8) not null default "",
    url varchar(4096) not null default "",
    status_code int not null default 0,
    bytes int not null default 0
) engine=innodb default charset utf8mb4;

insert into accesslog(ip, logtime, method, url, status_code, bytes) values
("11.1.1.1", "2019-01-01", 'GET', "/", 200, 250),
("11.1.1.2", "2019-03-02", 'GET', "/login/", 200, 250),
("11.1.2.1", "2019-01-01", 'POST', "/", 200, 250),
("11.1.3.1", "2019-11-04", 'GET', "/logout", 200, 4094),
("11.1.3.3", "2019-05-21", 'POST', "/", 301, 250),
("11.1.1.4", "2015-04-23", 'GET', "/error/", 404, 5320),
("11.1.1.1", "2019-01-01", 'GET', "/", 200, 250),
("11.1.1.3", "2019-03-02", 'GET', "/login/", 200, 250),
("11.1.2.1", "2019-01-01", 'POST', "/", 200, 250),
("11.1.3.2", "2019-11-04", 'GET', "/logout", 200, 4094),
("11.1.1.3", "2019-05-21", 'POST', "/", 301, 250),
("11.1.1.3", "2015-04-23", 'GET', "/error/", 404, 5320);