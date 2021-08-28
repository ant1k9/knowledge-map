[devd](https://github.com/cortesi/devd#routes)

*Description*: A local webserver for developers

*Labels*: #Go #WebServer

*Examples*:

```bash
$ devd -lo .
$ devd -l \
    -w ./src/ \
    /=http://localhost:8888 \
    /api/=http://localhost:8889 \
    /static/=./assets
```
