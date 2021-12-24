[dsq](https://github.com/multiprocessio/datastation/tree/main/runner/cmd/dsq)

*Description*: Run SQL queries against JSON, CSV, Excel, Parquet, and more

*Labels*: #Go #SQL #JSON #CSV

*Examples*:

```bash
$ cat testdata.csv | dsq csv "SELECT * FROM {} LIMIT 1"
$ dsq testdata.json "SELECT * FROM {} WHERE x > 10"
```
