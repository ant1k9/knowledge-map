[bats](https://github.com/bats-core/bats-core)

*Description*: Bats is a TAP-compliant testing framework for Bash. It provides a simple way to verify that the UNIX programs you write behave as expected.

*Labels*: #Bash #Testing

*Examples*:

```bash
$ cat > test << EOF
#!/usr/bin/env bats

@test "addition using bc" {
    result="$(echo 2+2 | bc)"
    [ "$result" -eq 4  ]
}
EOF
$ bats --tap test
```
