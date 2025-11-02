drop table if exists "user" cascade;
create table "user"
(
    id    uuid primary key default gen_random_uuid(),
    name  text not null,
    phone text not null
);

drop table if exists user_email cascade;
create table user_email
(
    id      uuid primary key default gen_random_uuid(),
    user_id uuid not null references "user" (id),
    email   text not null
);

insert into "user"(name, phone)
values ('Peter', '79778725196'),
       ('Olga', '79778725196'),
       ('Masha', '79778725196'),
       ('Dasha', '79778725196');

insert into user_email(user_id, email) (select u.id, 'someemail@email.com' from "user" u);
insert into user_email(user_id, email) (select u.id, 'other@email.com' from "user" u);
insert into user_email(user_id, email) (select u.id, 'third@email.com' from "user" u);

select u.id, u.name, u.phone, json_agg(ue.*)
from "user" u
         join user_email ue on u.id = ue.user_id
where u.name = 'Peter'
group by 1, 2, 3