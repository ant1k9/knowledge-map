[jq](https://gist.github.com/olih/f7437fb6962fb3ee9fe95bda8d2c8fa4)

*Description*: jq is useful to slice, filter, map and transform structured json data.

*Labels*: #C #JSON

*Docs*: https://stedolan.github.io/jq/

*Examples*:

```bash
$ jq 'map(select(. >= 2))'
$ jq '[range(2;4)]'
$ jq 'group_by(.foo)'
$ jq 'sort_by(.foo)'
```
