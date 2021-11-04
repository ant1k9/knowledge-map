[tbls](https://github.com/k1LoW/tbls)

*Description*: tbls is a CI-Friendly tool for document a database, written in Go.

*Labels*: #Go #SQL #Docs

*Examples*:

```bash
$ echo 'dsn: postgres://dbuser:dbpass@localhost:5432/dbname
docPath: doc/schema' > .tbls.yml
$ tbls doc
```
