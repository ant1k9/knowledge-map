[dasel](https://github.com/TomWright/dasel)

*Description*: Select, put and delete data from JSON, TOML, YAML, XML and CSV files with a single tool. Supports conversion between formats and can be used as a Go package.

*Docs*: https://daseldocs.tomwright.me/

*Labels*: #YAML #Go

*Examples*:

```bash
# convert json to yaml
$ echo '{"name": "Tom"}' | dasel -r json -w yaml

# update field
$ echo 'name: Tom' | dasel put string -r yaml '.email' 'contact@tomwright.me'

# update field in file in-place
$ dasel put string -f example.yaml '.email' 'contact@tomwright.me'
```
