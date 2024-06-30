-- 创建user表
create table go_zero_demo.user
(
    id   int auto_increment
        primary key,
    name varchar(20) not null,
    age  int         not null
);