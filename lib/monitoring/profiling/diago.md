[diago](https://github.com/remeh/diago)

*Description*: Diago is a visualization tool for CPU profiles and heap snapshots generated with `pprof`.

*Labels*: #Go #Profiling

*Links*:
  - https://remy.io/blog/how-to-use-diago-to-diagnose-cpu-and-memory-usage-in-go-programs/

*Examples*:

```bash
$ curl "http://localhost:6060/debug/pprof/profile?seconds=30" > profile
$ diago -file profile
```
