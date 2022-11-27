[speedbump](https://github.com/kffl/speedbump)

*Description*: TCP proxy for simulating variable, yet predictable network latency

*Labels*: #Go #TCP #ChaosEngineering

*Examples*:

```bash
$ speedbump --latency=100ms --sine-amplitude=100ms --sine-period=1m --port=2000 localhost:80
```
