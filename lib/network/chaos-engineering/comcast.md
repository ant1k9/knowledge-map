[comcast](https://github.com/tylertreat/comcast)

*Description*: Simulating shitty network connections so you can build better systems.

*Labels*: #Go #Networking #ChaosEngineering

*Examples*:

```bash
$ comcast --device=eth0 --latency=250 --target-bw=1000 --default-bw=1000000 --packet-loss=10% --target-addr=8.8.8.8,10.0.0.0/24 --target-proto=tcp,udp,icmp --target-port=80,22,1000:2000
$ comcast --device=eth0 --latency=250 --target-bw=1000 --packet-loss=10%
$ comcast stop
```
