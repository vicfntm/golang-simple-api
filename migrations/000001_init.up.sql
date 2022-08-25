create table if not exists customers (
id serial primary key,
login varchar unique not null,
password varchar not null,
role varchar default null,
created_on timestamp default null 
);

create table if not exists groups (
id serial primary key,
name varchar not null,
definition varchar default null,
created_on timestamp default null
);

create table if not exists customer_groups (
id serial primary key,
customer_id int not null,
group_id int not null,
created_on timestamp default null,
foreign key (customer_id) references customers (id),
foreign key (group_id) references groups (id)
);

create table if not exists shits (
id serial primary key,
group_id int not null,
customer_id int not null,
created_on timestamp default null,
foreign key(customer_id) references customers (id),
foreign key (group_id) references groups (id)
);
