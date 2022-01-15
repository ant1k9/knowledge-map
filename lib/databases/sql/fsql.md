[fsql](https://github.com/kashav/fsql)

*Description*: Find files with SQL.

*Labels*: #Go #SQL

*Examples*:

```bash
$ fsql "SELECT name, size, mode FROM . WHERE mode IS REG"
$ fsql "SELECT name, size, mode FROM . WHERE name LIKE .go"
```
