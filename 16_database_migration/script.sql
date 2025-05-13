create table category(
  id   integer      auto_increment primary key,
  name varchar(200) not null
);

insert into category(name)
  values("Gadget");

rename table category to categories;