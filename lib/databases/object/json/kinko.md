[kinko](https://github.com/Kinto/kinto)

*Description*: A generic JSON document store with sharing and synchronisation capabilities.

*Labels*: #Python #JSON #Database

*Docs*: https://docs.kinto-storage.org/en/latest/

*Examples*:

```bash
$ echo '{"data": {"password": "s3cr3t"}}' | \
    http PUT https://kinto.dev.mozaws.net/v1/accounts/bob -v
$ echo '{"data": {"description": "Write a tutorial explaining Kinto", "status": "todo"}}' | \
    http POST https://kinto.dev.mozaws.net/v1/buckets/default/collections/tasks/records \
        -v --auth 'bob:s3cr3t'
$ http GET https://kinto.dev.mozaws.net/v1/buckets/default/collections/tasks/records \
    -v --auth 'bob:s3cr3t'
```
