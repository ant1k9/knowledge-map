[fgprof](https://github.com/felixge/fgprof)

*Description*: 🚀 fgprof is a sampling Go profiler that allows you to analyze On-CPU as well as Off-CPU (e.g. I/O) time together.

*Labels*: #Go #Profiling

*Links*:
  - https://github.com/felixge/fgtrace

*Examples*:

```go
// go tool pprof --http=:6061 http://localhost:6060/debug/fgprof?seconds=3

import(
	_ "net/http/pprof"
	"github.com/felixge/fgprof"
)

func main() {
	http.DefaultServeMux.Handle("/debug/fgprof", fgprof.Handler())
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	// <code to profile>
} 
```
