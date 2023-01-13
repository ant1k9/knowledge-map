[visql](https://github.com/paulfitz/visql)

*Description*: edit slices of SQL databases in vi

*Labels*: #Python #SQL

*Examples*:

```bash
$ visql postgresql://test:test@localhost:5432/test --table my-table --sql 'id = 10'
$ visql postgresql://test:test@localhost:5432/test --table my-table --grep 'options'
```
