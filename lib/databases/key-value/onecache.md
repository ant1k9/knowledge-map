[onecache](https://github.com/dadgar/onecache)

*Description*: OneCache is a best-effort, replicated KV store accessible via the memcached protocol

*Labels*: #Go #Gossip

*Examples*:

```bash
$ ./onecache -gossip_port=7946 -port=11211
$ ./onecache -gossip_port=7947 -port=11212 -join="127.0.0.1:7946"
```
