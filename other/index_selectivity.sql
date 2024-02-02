drop table if exists selectivity_test;
create table selectivity_test
(
    id            serial primary key,
    internal_id   uuid unique,
    birthday_year int,
    gender        text,
    level         int
);

insert into selectivity_test(internal_id, birthday_year, gender, level)
select uuid_generate_v4(),
       (random() * 124 + 1900):: int,
       case when random() > 0.5 then 'male' else 'female' end,
       (random() * 10 + 0):: int
from generate_series(1, 1000000);


explain analyse
select *
from selectivity_test
where birthday_year = 2000
limit 10;
-- Seq Scan
-- Execution Time: 0.144 ms


drop index if exists selectivity_test_birthday_year;
create index selectivity_test_birthday_year on selectivity_test (birthday_year);

explain analyse
select *
from selectivity_test
where birthday_year = 2000
limit 10;
-- Index Scan
-- Execution Time: 0.096 ms


explain analyse
select *
from selectivity_test
where gender = 'female'
limit 10;
-- Seq Scan
-- Execution Time: 0.027 ms

drop index if exists selectivity_test_gender;
create index selectivity_test_gender on selectivity_test (gender);

explain analyse
select *
from selectivity_test
where gender = 'female'
limit 10;
-- Seq Scan
-- Execution Time: 0.023 ms


explain analyse
select *
from selectivity_test
where level = 1
limit 10;
-- Seq Scan
-- Execution Time: 0.030 ms

drop index if exists selectivity_test_level;
create index selectivity_test_level on selectivity_test (level);

explain analyse
select *
from selectivity_test
where level = 1
limit 10;
-- Index Scan
-- Execution Time: 0.061 ms


drop index if exists selectivity_test_level;
drop index if exists selectivity_test_gender;
drop index if exists selectivity_test_birthday_year;

explain analyse
select *
from selectivity_test
where level = 1 and birthday_year = 2000
limit 10;
-- Seq Scan
-- Execution Time: 0.795 ms

drop index if exists selectivity_test_level_birthday_year;
create index selectivity_test_level_birthday_year on selectivity_test (level, birthday_year);

explain analyse
select *
from selectivity_test
where level = 1 and birthday_year = 2000
limit 10;
-- Index Scan
-- Execution Time: 0.095 ms
