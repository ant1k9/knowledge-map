[panicparse](https://github.com/maruel/panicparse)

*Description*: Crash your app in style (Golang). Parses panic stack traces, densifies and deduplicates goroutines with similar stack traces. Helps debugging crashes and deadlocks in heavily parallelized process.

*Labels*: #Go

*Examples*:

```bash
$ go get github.com/maruel/panicparse/v2/cmd/pp
$ go run main.go |& pp
```
