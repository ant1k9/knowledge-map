[ql](https://gitlab.com/cznic/ql/-/tree/master)

*Description*: Package ql is a pure Go embedded (S)QL database.

*Labels*: #Go #Database

*Docs*: https://pkg.go.dev/modernc.org/ql/ql?utm_source=godoc

*Examples*:

```bash
$ ql 'create table t (i int, s string)'
$ ql << EOF
insert into t values
(1, "a"),
(2, "b"),
(3, "c"),
EOF
$ ql 'select * from t'
$ ql -fld 'select * from t where i != 2 order by s'
```
