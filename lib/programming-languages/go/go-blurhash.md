[go-blurhash](https://github.com/buckket/go-blurhash)

*Description*: A Blurhash implementation in pure Go (Decode/Encode)

*Labels*: #Go #Hash #Images

*Docs*: https://blurha.sh/

*Links*:
  - https://github.com/woltapp/blurhash

*Examples*:

```go
package main

import (
	"fmt"
	"log"
	"github.com/buckket/go-blurhash"
	"image/png"
	"os"
)

func main() {
	imageFile, _ := os.Open("test.png")
	loadedImage, err := png.Decode(imageFile)
	str, _ := blurhash.Encode(4, 3, loadedImage)
	fmt.Printf("Hash: %s\n", str)

	img, _ := blurhash.Decode(str, 300, 500, 1)
	f, _ := os.Create("test_blur.png")
	png.Encode(f, img)

	x, y, _ := blurhash.Components("LFE.@D9F01_2%L%MIVD*9Goe-;WB")
	fmt.Printf("xComponents: %d, yComponents: %d", x, y)
}
```
