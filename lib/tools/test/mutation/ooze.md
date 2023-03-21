[ooze](https://github.com/gtramontina/ooze)

*Description*: 🧬 Go Mutation Testing

*Labels*: #Go #MutationTesting #Testing

*Examples*:

```bash
$ cat > mutation_test.go <<EOF
//go:build mutation

package main_test

import (
    "testing"

    "github.com/gtramontina/ooze"
)

func TestMutation(t *testing.T) {
    ooze.Release(t)
}
EOF
$ go test -v -tags=mutation
```
