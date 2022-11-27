[go-deadlock](https://github.com/sasha-s/go-deadlock)

*Description*: Online deadlock detection in go (golang)

*Labels*: #Go #Mutex

*Examples*:

```go
import "github.com/sasha-s/go-deadlock"

func main() {
    var mu deadlock.Mutex
    // Use normally, it works exactly like sync.Mutex does.
    mu.Lock()
    defer mu.Unlock()

    // Or
    var rw deadlock.RWMutex
    rw.RLock()
    defer rw.RUnlock() 
}
```
