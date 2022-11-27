[gobreaker](https://github.com/sony/gobreaker)

*Description*: Circuit Breaker implemented in Go

*Labels*: #Go #CircuitBreaker

*Examples*:

```go
var cb *breaker.CircuitBreaker

func Get(url string) ([]byte, error) {
	body, err := cb.Execute(func() (interface{}, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	})
	if err != nil {
		return nil, err
	}

	return body.([]byte), nil
}
```
