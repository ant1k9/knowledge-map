[pspg](https://github.com/okbob/pspg)

*Description*: Unix pager (with very rich functionality) designed for work with tables. Designed for PostgreSQL, but MySQL is supported too. Works well with pgcli too. Can be used as CSV or TSV viewer too. It supports searching, selecting rows, columns, or block and export selected area to clipboard. 

*Labels*: #C #PostgreSQL

*Examples*:

```bash
$ psql -c "SELECT 1" | PSPG='--style=7' pspg
```
