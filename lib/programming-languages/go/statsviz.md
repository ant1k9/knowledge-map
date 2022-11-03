[statsviz](https://github.com/arl/statsviz)

*Description*: :rocket: Visualise Go program runtime metrics in real time in your browser

*Labels*: #Go #Monitoring

*Examples*:

```go
mux := http.NewServeMux()
statsviz.Register(mux) // http://localhost:6060/debug/statsviz/
```
