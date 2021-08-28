[rqlite](https://github.com/rqlite/rqlite)

*Description*: The lightweight, distributed relational database built on SQLite

*Labels*: #Go #Database #Raft #SQLite

*Docs*: https://www.philipotoole.com/tag/rqlite/

*Examples*:

```bash
$ rqlited -node-id 1 ~/node.1
$ rqlited -node-id 2 -http-addr localhost:4003 -raft-addr localhost:4004 \
    -join http://localhost:4001 ~/node.2
$ rqlited -node-id 3 -http-addr localhost:4005 -raft-addr localhost:4006 \
    -join http://localhost:4001 ~/node.3
$ rqlite
    CREATE TABLE foo (id INTEGER NOT NULL PRIMARY KEY, name TEXT);
    .schema
    INSERT INTO foo(name) VALUES("fiona");
    SELECT * FROM foo;
```
