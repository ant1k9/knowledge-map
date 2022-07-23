[go-mysql-server](https://github.com/dolthub/go-mysql-server)

*Description*: A MySQL-compatible relational database with a storage agnostic query engine. Implemented in pure Go.

*Labels*: #Go #MySQL

*Examples*:

```go
package main

import (
	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/information_schema"
)

// Example of how to implement a MySQL server based on a Engine:
//
// ```
// > mysql --host=127.0.0.1 --port=3306 -u root mydb -e "SELECT * FROM mytable"
// +----------+-------------------+-------------------------------+---------------------+
// | name     | email             | phone_numbers                 | created_at          |
// +----------+-------------------+-------------------------------+---------------------+
// | John Doe | john@doe.com      | ["555-555-555"]               | 2018-04-18 09:41:13 |
// | John Doe | johnalt@doe.com   | []                            | 2018-04-18 09:41:13 |
// | Jane Doe | jane@doe.com      | []                            | 2018-04-18 09:41:13 |
// | Evil Bob | evilbob@gmail.com | ["555-666-555","666-666-666"] | 2018-04-18 09:41:13 |
// +----------+-------------------+-------------------------------+---------------------+
// ```
func main() {
	engine := sqle.NewDefault(
		sql.NewDatabaseProvider(
			memory.NewDatabase("mydb"),
			information_schema.NewInformationSchemaDatabase(),
		))
	config := server.Config{
		Protocol: "tcp",
		Address:  "localhost:3306",
	}
	s, err := server.NewDefaultServer(config, engine)
	if err != nil {
		panic(err)
	}
	s.Start()
}
```
