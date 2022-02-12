[quinedb](https://github.com/gfredericks/quinedb)

*Description*: QuineDB is a quine that is also a key-value store.

*Labels*: #Shell #Database #KeyValue

*Examples*:

```bash
$ cat > ~/local/bin/qdb << EOF
qdb () {
    ./quinedb "$@" > qdb.out 2> qdb.err || return "$?"
        cat qdb.out > quinedb
        cat qdb.err
        rm qdb.err qdb.out
}
$ qdb set foo bar
$ qdb get foo
```
