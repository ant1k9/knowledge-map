[websocketd](https://github.com/joewalnes/websocketd)

*Description*: Turn any program that uses STDIN/STDOUT into a WebSocket server. Like inetd, but for WebSockets.

*Docs*: https://github.com/joewalnes/websocketd/wiki

*Labels*: #Go #Networking #Websockets #Server

*Examples*:

```bash
$ websocketd --port 8080 my-program
$ websocketd --port 8080 ipython
$ websocketd --port=8080 sh -c 'sh 2>&1'  # https://vcs.rowanthorpe.com/rowan/ws-repl
```
