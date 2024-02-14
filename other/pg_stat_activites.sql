select datname, usename, application_name, query, count(1)
from pg_stat_activity
group by 1, 2, 3, 4
order by 2, 3 desc;
