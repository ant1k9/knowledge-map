[act](https://github.com/nektos/act)

*Description*: Run your GitHub Actions locally 🚀

*Labels*: #Github #GithubActions #Go

*Examples*:

```bash
act -l                   # List the actions for the default event:
act workflow_dispatch -l # List the actions for a specific event:
act                      # Run the default (`push`) event:
act pull_request         # Run a specific event:
act -j test              # Run a specific job:
act -n                   # Run in dry-run mode:
```
