[Curl Performance mertics](https://speedtestdemon.com/a-guide-to-curls-performance-metrics-how-to-analyze-a-speed-test-result/)

*Description*: Cheat Sheet on Curl Performance Metrics: how to benchmark server latency with curl

*Labels*: #Curl #Networking #Performance

*Examples*:

```bash
$ echo '
time_namelookup: %{time_namelookup}s\n
time_connect: %{time_connect}s\n
time_appconnect: %{time_appconnect}s\n
time_pretransfer: %{time_pretransfer}s\n
time_redirect: %{time_redirect}s\n
time_starttransfer: %{time_starttransfer}s\n
----------\n
time_total: %{time_total}s\n
' > /tmp/curl-format.txt
$ curl -L -w "@/tmp/curl-format.txt" -o tmp -s "https://www.google.com/"
```
