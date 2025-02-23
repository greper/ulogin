create table login_log (
    id int primary key auto_increment,
    user_id int not null,
    login_time datetime not null
);