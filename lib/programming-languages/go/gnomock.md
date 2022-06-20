[gnomock](https://github.com/orlangure/gnomock)

*Description*: Test your code without writing mocks with ephemeral Docker containers 📦 Setup popular services with just a couple lines of code ⏱️ No bash, no yaml, only code 💻

*Labels*: #Go #Docker #Testing

*Examples*:

```go
package db_test

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq" // postgres driver
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
)

func TestDB(t *testing.T) {
	p := postgres.Preset(
		postgres.WithUser("gnomock", "gnomick"),
		postgres.WithDatabase("mydb"),
		postgres.WithQueriesFile("/var/project/db/schema.sql"),
	)
	container, _ := gnomock.Start(p)
	t.Cleanup(func() { _ = gnomock.Stop(container) })

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
		container.Host, container.DefaultPort(),
		"gnomock", "gnomick", "mydb",
	)
	db, _ := sql.Open("postgres", connStr)
}
```
