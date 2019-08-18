create table webdemo_admin
(
    admin_id int not null auto_increment,
    admin_name varchar(32) not null,
    admin_password varchar(32) not null,
    primary key(admin_id)
);

create table webdemo_jobtable
(
    id int not null auto_increment,
    department_name varchar(60) not null,
    department_id varchar(20) not null,
    position_numbers int not null,
    major varchar(200) not null,
    education varchar(200) not null,
    others varchar(200) ,
    primary key(admin_id)
);
