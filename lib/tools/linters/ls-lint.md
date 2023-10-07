[ls-lint](https://github.com/loeffel-io/ls-lint)

*Description*: An extremely fast directory and filename linter - Bring some structure to your project filesystem

*Labels*: #Go #Linters

*Docs*: https://ls-lint.org

*Examples*:

```bash
$ echo '
ls:
  .js: snake_case
  .ts: snake_case | camelCase
  .d.ts: PascalCase
  .html: regex:[a-z0-9]+

ignore:
  - node_modules' > .ls-lint.yml
$ ls-lint
```
