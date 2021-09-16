[mpb](https://github.com/vbauerster/mpb)

*Description*: multi progress bar for Go cli applications

*Labels*: #Go #ProgressBar

*Examples*:

```go
package main

import (
    "math/rand"
    "time"

    "github.com/vbauerster/mpb/v7"
    "github.com/vbauerster/mpb/v7/decor"
)

func main() {
    p := mpb.New(mpb.WithWidth(64))

    total := 100
    name := "Single Bar:"
    bar := p.Add(int64(total),
        mpb.NewBarFiller(mpb.BarStyle().Lbound("╢").Filler("▌").Tip("▌").Padding("░").Rbound("╟")),
        mpb.PrependDecorators(
            decor.Name(name, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
            decor.OnComplete(
                decor.AverageETA(decor.ET_STYLE_GO, decor.WC{W: 4}), "done",
            ),
        ),
        mpb.AppendDecorators(decor.Percentage()),
    )
    max := 100 * time.Millisecond
    for i := 0; i < total; i++ {
        time.Sleep(time.Duration(rand.Intn(10)+1) * max / 10)
        bar.Increment()
    }
    p.Wait()
}
```
