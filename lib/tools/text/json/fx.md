[fx](https://github.com/antonmedv/fx)

*Description*: Command-line tool and terminal JSON viewer

*Labels*: #JSON #JavaScript

*Docs*: https://github.com/antonmedv/fx/blob/master/DOCS.md

*Examples*:

```bash
$ echo '{"foo": [{"bar": "value"}]}' | fx 'x => x.foo[0].bar'
$ echo '{"foo": "bar"}' | fx .
$ echo '{"foo": 1, "bar": 2}' | fx .[]
$ echo '{"count": 0}' | fx '{...this, count: 1}'
$ echo '123' | fx 'x => x * 2'
```
