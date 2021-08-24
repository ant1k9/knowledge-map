[rlite](https://github.com/seppo0010/rlite)

*Description*: self-contained, serverless, zero-configuration, transactional redis-compatible database engine. rlite is to Redis what SQLite is to SQL.

*Labels*: #Redis #C

*Examples*:

```go
package main

import "github.com/seppo0010/rlite-go"
import "fmt"

func main () {
    db, _ := rlite.Open(":memory:")
    rlite.Command(db, []string{"SET", "key", "value"})

    reply, err := rlite.Command(db, []string{"GET", "key"})
    if err != nil {
        // ...
    }
    if reply != "value" {
        // ...
    }
    fmt.Println(reply)
}
```
