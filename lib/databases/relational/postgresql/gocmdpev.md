[gocmdpev](https://github.com/simon-engledew/gocmdpev)

*Description*: A command-line GO Postgres query visualizer, heavily inspired by pev (https://github.com/AlexTatiyants/pev)

*Labels*: #Go #PostgresSQL

*Examples*:

```bash
$ echo <query> \
    | sed '1s/^/EXPLAIN (ANALYZE, COSTS, VERBOSE, BUFFERS, FORMAT JSON) /' \
    | psql -qXAt <DATABASE> \
    | gocmdpev

```
