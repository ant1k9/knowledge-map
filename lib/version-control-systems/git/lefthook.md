[lefthook](https://github.com/evilmartians/lefthook)

*Description*: Fast and powerful Git hooks manager for any type of projects.

*Labels*: #Go #Git

*Examples*:

```bash
$ lefthook install
$ echo '---
pre-push:
  scripts:
    "audit.sh":
      runner: bash' > lefthook.yml
$ lefthook add --dirs pre-push
```
