[soft-serve](https://github.com/charmbracelet/soft-serve)

*Description*: A tasty, self-hostable Git server for the command line🍦

*Labels*: #Go #Git #GitServer

*Examples*:

```bash
$ cat > serve.config << EOF
name: Soft Serve
host: localhost
port: 23231
anon-access: read-write
allow-keyless: false

repos:
  - name: Home
    repo: config
    private: true
    note: "Configuration and content repo for this server"

users:
  - name: Beatrice
    admin: true
    public-keys:
      - ssh-rsa AAAAB3Nz...   # redacted
      - ssh-ed25519 AAAA...   # redacted
EOF
$ soft serve
```
