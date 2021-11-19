[hugo](https://github.com/gohugoio/hugo)

*Description*: Hugo is a static HTML and CSS website generator written in Go. It is optimized for speed, ease of use, and configurability. Hugo takes a directory with content and templates and renders them into a full HTML website.

*Labels*: #Go #HTML

*Docs*: https://gohugo.io/documentation/

*Examples*:

```bash
$ hugo new site quickstart
$ cd quickstart
$ git init
$ git submodule add https://github.com/theNewDynamic/gohugo-theme-ananke.git themes/ananke
$ echo theme = \"ananke\" >> config.toml
$ hugo new posts/my-first-post.md
$ hugo server -D
```
