[semgrep](https://github.com/returntocorp/semgrep)

*Description*: Lightweight static analysis for many languages. Find bug variants with patterns that look like source code.

*Labels*: #OCaml #Linters

*Docs*: https://semgrep.dev/docs/

*Examples*:

```bash
$ semgrep -e '$X == $X' --lang=py path/to/src
$ semgrep --config=p/r2c-ci path/to/src
``
