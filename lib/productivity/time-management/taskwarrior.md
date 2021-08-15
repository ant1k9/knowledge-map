[taskwarrior](https://github.com/GothenburgBitFactory/taskwarrior)

*Description*: Taskwarrior - Command line Task Management

*Labels*: #C++ #Productivity #TimeManagement

*Docs*: https://taskwarrior.org/docs/

*Examples*:

```bash
$ cp first_last.cert.pem ~/.task
$ cp first_last.key.pem  ~/.task
$ cp ca.cert.pem         ~/.task
$ task config taskd.certificate -- ~/.task/first_last.cert.pem
$ task config taskd.key         -- ~/.task/first_last.key.pem
$ task config taskd.ca          -- ~/.task/ca.cert.pem
$ task config taskd.server      -- host.domain:53589
$ task config taskd.credentials -- Public/First Last/cf31f287-ee9e-43a8-843e-e8bbd5de4294
$ task config taskd.trust       -- ignore hostname
$ task sync
```
