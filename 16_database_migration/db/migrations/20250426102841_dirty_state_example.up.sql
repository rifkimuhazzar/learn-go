create table correct(
  id int not null auto_increment,
  name varchar(255) not null,
  primary key(id)
) engine = innodb;

create table wrong(
  id int not null auto_increment,
  name varchar(255) not null,
  primary key(id)
) engine = innodb;
