[process-compose](https://github.com/F1bonacc1/process-compose)

*Description*: Process Compose is like docker-compose, but for orchestrating a suite of processes, not containers.

*Labels*: #Go #Watchdog

*Examples*:

```bash
$ echo 'version: "0.5"

processes:
  ls:
    command: "ls"
    availability:
      restart: "always"

  date:
    command: "date"
    availability:
      restart: "always"
' > process-compose.yaml
$ process-compose
```
