package serve

import (
	"strings"
)

type (
	File struct {
		Name             string
		Link             string
		Description      string
		Docs             string
		Labels           []string
		Examples         string
		ExamplesLanguage string
		ExtraLinks       []string

		Path string
	}
)

func (f *File) hasLabel(label string) bool {
	for _, l := range f.Labels {
		if strings.ToLower(l) == label {
			return true
		}
	}
	return false
}

func (f *File) matchDescription(q string) bool {
	for _, part := range strings.Split(q, " ") {
		if !strings.Contains(strings.ToLower(f.Description), part) &&
			!strings.Contains(strings.ToLower(f.Name), part) {
			return false
		}
	}
	return true
}
