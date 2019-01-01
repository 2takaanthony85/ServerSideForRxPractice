CREATE DATABASE sample_db;
use sample_db;

CREATE TABLE users (
  id int(10) unsigned not null auto_increment,
  uuid varchar(64) not null unique,
  name varchar(255),
  email varchar(255) not null unique,
  password varchar(255) not null,
  primary key (id)
);

CREATE TABLE todos (
  id int(10) unsigned not null auto_increment,
  uuid varchar(64) not null unique,
  title varchar(255),
  contents text,
  user_id int(10) unsigned,
  foreign key(user_id) references users(id),
  primary key (id)
);