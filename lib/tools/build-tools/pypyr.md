[pypyr](https://github.com/pypyr/pypyr)

*Description*: pypyr task-runner cli & api for automation pipelines. Automate anything by combining commands, different scripts in different languages & applications into one pipeline process.

*Labels*: #Python #BuildTools

*Docs*: https://pypyr.io

*Examples*:

```bash
$ echo 'context_parser: pypyr.parser.keyvaluepairs
steps:
  - name: pypyr.steps.echo
    in:
      echoMe: o hai!
  - name: pypyr.steps.cmd
    in:
      cmd: echo any cmd you like' > task.yaml
$ pypyr task
o hai!
any cmd you like
```
