[benchdraw](https://github.com/cep21/benchdraw)

*Description*: A CLI to turn Go's benchmark output into pictures

*Labels*: #Go #Benchmarks #Charts

*Examples*:

```bash
$ echo > add_test.go << EOF
package main

import "testing"

func BenchmarkAddSimple(b *testing.B) {
	b.Run("some=abc", func(b *testing.B) {
		var s int
		for i := 0; i < b.N; i++ {
			s += i
		}
		_ = s
	})

	b.Run("some=def", func(b *testing.B) {
		var s int
		for i := 0; i < b.N; i++ {
			func() {
				s += i
			}()
		}
		_ = s
	})
}
EOF
$ go test -benchmem -bench=. -v > benchmark.txt
$ benchdraw --x=some < benchmark.txt > benchmark.svg
```
