[task](https://github.com/go-task/task)

*Description*: A task runner / simpler Make alternative written in Go

*Labels*: #Go #BuildTools #Make

*Docs*: https://taskfile.dev/#/usage?id=dynamic-variables

*Examples*:

```bash
$ cat > Taskfile.yml << EOF
version: '3'

tasks:
  greet:
    cmds:
      - echo $GREETING
  env:
    GREETING: Hey, there!
EOF
$ task greet
```
