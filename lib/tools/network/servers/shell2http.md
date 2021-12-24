[shell2http](https://github.com/msoap/shell2http)

*Description*: Executing shell commands via HTTP server

*Labels*: #Go #WebServer

*Examples*:

```bash
$ shell2http /top "top -l 1 | head -10"
$ shell2http /date date /ps "ps aux"
$ shell2http -export-vars=GOPATH /get 'echo $GOPATH'
$ shell2http /cal_html 'echo "<html><body><h1>Calendar</h1>Date: <b>$(date)</b><br><pre>$(cal $(date +%Y))</pre></body></html>"'
```
