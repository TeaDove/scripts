select table_name,
       pg_size_pretty(pg_total_relation_size(table_schema || '.' || quote_ident(table_name))) as "total size",
       pg_size_pretty(pg_indexes_size(table_schema || '.' || quote_ident(table_name)))        as "index size",
       pg_size_pretty(pg_relation_size(table_schema || '.' || quote_ident(table_name)))       as "relation size",
       pg_relation_size(table_schema || '.' || quote_ident(table_name))
from information_schema.tables
where table_schema = 'public'
order by 5 desc;
