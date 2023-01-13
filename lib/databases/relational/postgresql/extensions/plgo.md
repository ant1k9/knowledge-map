[plgo](https://gitlab.com/microo8/plgo)

*Description*: easily create postgresql extensions in golang

*Labels*: #Go #PostgreSQL

*Examples*:

```go
//must be main package

package main

import (
	"log"
	"strings"

	"gitlab.com/microo8/plgo"
)

//from every exported function will be generated a stored procedure
//functions can take (and return) any golang builtin type (like string, int, float64, []int, ...)

func Meh() {
    //NoticeLogger for printing notice messages to elog
    logger := plgo.NewNoticeLogger("", log.Ltime|log.Lshortfile)
    logger.Println("meh")
}

//ConcatArray concatenates an array of strings
//function arguments (and return values) can be also array types of the golang builtin types
func ConcatArray(strs []string) string {
    return strings.Join(strs, "")
}
```
