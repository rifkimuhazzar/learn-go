create table customers (
  id   varchar(100) not null,
  name varchar(100) not null,
  primary key (id)
);

delete from customers;

alter table customers 
  add column email      varchar(100),
  add column balance    int           default 0,
  add column rating     double        default 0.0,
  add column created_at  timestamp     default current_timestamp,
  add column birth_date date,
  add column married    boolean       default false;

insert into customers(id, name, email, balance, rating, birth_date, married)
  values("react", "React", "react@example.com", 1000000, 90.0, "2000-12-20", true),
        ("vue", "Vue", "vue@example.com", 2000000, 85.5, "2001-12-20", false),
        ("svelte", "Svelte", "svelte@example.com", 3000000, 95.7, "2002-12-20", false);

update customers
  set email = null,
      birth_date = null
  where id = "svelte"

create table users (
  username varchar(100) not null,
  password varchar(100) not null,
  primary key (username)
);

insert into users(username, password)
  value('admin', 'admin');

insert into users(username, password)
  value('admin2', 'admin2'),
       ('admin3', 'admin3'),
       ('admin4', 'admin4');

create table comments(
  id      int          not null auto_increment,
  email   varchar(100) not null,
  comment text,
  primary key(id)
);