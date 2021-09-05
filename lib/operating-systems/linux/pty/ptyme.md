[ptyme](https://github.com/iximiuz/ptyme)

*Description*: Have you ever been wondering how docker attach command is implemented under the hood? This article covers the basics of Linux pseudoterminal capabilities and continuously shows how attach-like feature can be implemented in a ridiculously small amount of code.

*Labels*: #C #PTY #Linux #Docker

*Docs*: https://iximiuz.com/en/posts/linux-pty-what-powers-docker-attach-functionality/

*Examples*:

```bash
$ make shim
$ ./build/shim 43210 /bin/bash
$ go mod init ptyme
$ go mod tidy
$ make attach SOCK=localhost:43210
```
