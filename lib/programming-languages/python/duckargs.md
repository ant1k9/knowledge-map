[duckargs](https://github.com/eriknyquist/duckargs)

*Description*: Productivity tool for quickly creating python programs that parse command-line arguments. Stop writing argparse boilerplate code!

*Labels*: #Python #CLI

*Examples*:

```bash
$ python -m duckargs \
    positional_arg1 positional_arg2 \
    -i --int-val 4 \
    -f 3.3 \
    -F --file file_that_exists \
    -a -b -c > program.py
```
