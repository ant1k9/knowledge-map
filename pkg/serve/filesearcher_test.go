package serve

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fileSearcher_parseFileInfo(t *testing.T) {
	type args struct {
		content string
		path    string
	}

	tests := []struct {
		name string
		args args
		want File
	}{
		{
			name: "parse hugo information",
			args: args{
				content: `
[hugo](https://github.com/gohugoio/hugo)

*Description*: Hugo is a static HTML and CSS website generator written in Go. It is optimized for speed, ease of use, and configurability. Hugo takes a directory with content and templates and renders them into a full HTML website.

*Labels*: #Go #HTML

*Docs*: https://gohugo.io/documentation/

*Links*:
  - https://gohugo.io/content-management/
  - https://gohugo.io/getting-started/

*Examples*:

 ` + "```bash" + `
$ hugo new site quickstart
$ cd quickstart
$ git init
$ git submodule add https://github.com/theNewDynamic/gohugo-theme-ananke.git themes/ananke
$ echo theme = \"ananke\" >> config.toml
$ hugo new posts/my-first-post.md
$ hugo server -D
` + "```",
				path: "analytics/web/countly-server.md",
			},
			want: File{
				Name:        "hugo",
				Link:        "https://github.com/gohugoio/hugo",
				Description: "Hugo is a static HTML and CSS website generator written in Go. It is optimized for speed, ease of use, and configurability. Hugo takes a directory with content and templates and renders them into a full HTML website.",
				Docs:        "https://gohugo.io/documentation/",
				Labels:      []string{"Go", "HTML"},
				Examples: `$ hugo new site quickstart
$ cd quickstart
$ git init
$ git submodule add https://github.com/theNewDynamic/gohugo-theme-ananke.git themes/ananke
$ echo theme = \"ananke\" >> config.toml
$ hugo new posts/my-first-post.md
$ hugo server -D
`,
				ExamplesLanguage: "bash",
				ExtraLinks: []string{
					"https://gohugo.io/content-management/",
					"https://gohugo.io/getting-started/",
				},
				Path: "/analytics/web/countly-server.md",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewFileSearcher()
			assert.EqualValues(t, tt.want, s.parseFileInfo(tt.args.content, tt.args.path))
		})
	}
}
