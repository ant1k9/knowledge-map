[pg_flame](https://github.com/mgartner/pg_flame)

*Description*: A flamegraph generator for Postgres EXPLAIN ANALYZE output.

*Labels*: #Go #PostgreSQL #Performance

*Examples*:

```bash
$ psql dbname -qAtc 'EXPLAIN (ANALYZE, BUFFERS, FORMAT JSON) SELECT id FROM users' \
    | pg_flame \
    > flamegraph.html \
    && open flamegraph.html
```
