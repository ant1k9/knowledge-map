[modd](https://github.com/cortesi/modd)

*Description*: A flexible developer tool that runs processes and responds to filesystem changes

*Labels*: #Go #DevTools #HotReload #Watchdog

*Examples*:

```
$ cat << EOF > mod.conf
**/*.go {
    prep: go test @dirmods
}
EOF
$ modd
```
