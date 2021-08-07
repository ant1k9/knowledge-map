[A tcpdump Tutorial with Examples](https://danielmiessler.com/study/tcpdump/)

*Description*: Tutorial with 50 Ways to isolate traffic

*Labels*: #tcpdump #networking

*Examples*:

```bash
$ tcpdump -nnSX port 443  # https traffic
$ tcpdump -i eth0         # listen to interface
$ tcpdump host 1.1.1.1    # sniff traffic to host
$ tcpdump src 1.1.1.1
$ tcpdump dst 1.0.0.1
$ tcpdump icmp            # only icmp requests
```
