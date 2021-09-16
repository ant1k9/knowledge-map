[GlueSQL Sheets](https://sheets.gluesql.com/)

*Description*: Use SQL to manipulate Google Sheets!

*Labels*: #SQL #Rust

*Docs*: https://github.com/gluesql/gluesql

*Examples*:

```bash
$ curl https://sheets.gluesql.com/api/execute
  -H "Content-Type: application/json"
  -H "X-Api-Key: {API-KEY}"
  -d '{
    "spreadsheetId": "{SPREADSHEET-ID}",
    "sql": "{SQL-TO-EXECUTE}"
  }'
```
