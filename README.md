# SNSProject

#数据库创建记录
create table if not exists comment
(
	content varchar(300) not null,
	cid int not null,
	uid int not null,
	pid int not null
);

create table if not exists contact
(
	uid int not null
		primary key,
	body text null
);

create table if not exists pclass
(
	class char(10) not null,
	cid int not null,
	constraint pclass_class_uindex
		unique (class)
);

create table if not exists post
(
	title varchar(60) not null,
	pid int auto_increment
		primary key,
	uperid int not null,
	class int not null,
	comment int null,
	content varchar(500) not null,
	files text null
);

create table if not exists pulink
(
	pid int not null
		primary key,
	uid int not null
);

create table if not exists user
(
	uid int auto_increment
		primary key,
	username varchar(30) null,
	account varchar(20) not null,
	password char(255) not null,
	avatar varchar(255) null,
	slogan varchar(40) null,
	gender int default 0 not null,
	constraint user_account_uindex
		unique (account),
	constraint user_username_uindex
		unique (username)
);
#