[htmlq](https://github.com/mgdm/htmlq)

*Description*: Like jq, but for HTML. Uses CSS selectors to extract bits of content from HTML files. Mozilla's MDN has a good reference for CSS selector syntax.

*Labels*: #Rust #HTML

*Examples*:

```bash
$ curl --silent https://www.rust-lang.org/ | htmlq '#get-help'
$ curl --silent https://www.rust-lang.org/ | htmlq --attribute href a
```
