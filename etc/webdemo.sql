create table webdemo_admin
(
    admin_id int not null auto_increment,
    admin_name varchar(32) not null,
    admin_password varchar(32) not null,
    primary key(admin_id)
);

create table jobtable
(
    id int not null auto_increment,
    job_name varchar(60) not null,
    job_code varchar(20) not null,
    job_nums int not null,
    major varchar(200) not null,
    education varchar(200) not null,
    others varchar(200) ,
    primary key(id)
);

create table userinfo
(
    id_card int not null ,
    mobile_phone varchar(15) not null,
    user_name varchar(20) not null,
    school varchar(100) not null,
    sub_school varchar(100) not null,
    major varchar(100) not null,
    job_code1 varchar(20) not null,
    job_name1 varchar(40) not null,
    job_code2 varchar(20) default 'no',
    job_name2 varchar(40) default 'no',
    score1 int  default 0,
    score2 int  default 0,
    rank1 int  default 0,
    rank2 int  default 0,
    primary key(id_card)
);
