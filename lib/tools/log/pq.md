[pq](https://github.com/iximiuz/pq)

*Description*: Interactive query analytics

*Labels*: #Rust #Logs

*Examples*:

```bash
$ cat /tmp/some.log | pq 'json | map { .level as level, .ts:ts } |  select count_over_time(__line__[1s])'
```
