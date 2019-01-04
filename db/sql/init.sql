-- CREATE DATABASE sample_db;
use sample_db;

CREATE TABLE users (
  id int(10) unsigned not null auto_increment,
  uid varchar(64) not null unique,
  name varchar(255),
  email varchar(255) not null unique,
  password varchar(255) not null,
  primary key (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE todos (
  id int(10) unsigned not null auto_increment,
  uid varchar(64) not null unique,
  title varchar(255),
  contents text,
  user_id int(10) unsigned,
  foreign key(user_id) references users(id),
  primary key (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO users (uid, name, email, password) VALUES ('ksdjlfkj098', 'taka', 'lskdjfks@email.com', 'sodlkjwelkj99');