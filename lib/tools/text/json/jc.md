[jc](https://github.com/kellyjonbrazil/jc)

*Description*: CLI tool and python library that converts the output of popular command-line tools and file-types to JSON or Dictionaries. This allows piping of output to tools like jq and simplifying automation scripts.

*Labels*: #Python #JSON

*Examples*:

```bash
$ upower -d | jc --upower -p
$ dig example.com | jc --dig -p
```
