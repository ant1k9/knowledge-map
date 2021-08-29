package main

import (
	"github.com/ant1k9/knowledge-map/pkg/serve"
)

func main() {
	serve.Serve(serve.NewFileSearcher())
}
