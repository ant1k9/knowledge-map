[COLODEBUG](https://johannes.truschnigg.info/writing/2021-12_colodebug/)

*Description*: COLODEBUG: a simple way to improve bash script debugging

*Labels*: #Bash #Debug

*Examples*:

```bash
$ cat script.sh << EOF
if [[ -n ${COLODEBUG} && ${-} != *x*  ]]; then
    :() {
        [[ ${1:--} != ::*  ]] && return 0
        printf '%s\n' "${*}" >&2
    }
fi

: ::: make some computation
EOF
$ COLODEBUG=1 bash script.sh
::: make some computation
```
