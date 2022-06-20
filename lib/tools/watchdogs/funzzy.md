[funzzy](https://github.com/cristianoliveira/funzzy)

*Description*: Yet another fancy watcher. (Rust)

*Labels*: #Rust #Watchdog

*Examples*:

```bash
$ echo "
- name: run my tests
  run: make test
  change: 'tests/**'
  ignore: 'tests/integration/**'

- name: Starwars
  run: telnet towel.blinkenlights.nl
  change: '.watch.yaml'
" > .watch.yaml
$ funzzy watch
```
