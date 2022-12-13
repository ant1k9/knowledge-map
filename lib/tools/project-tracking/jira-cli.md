[jira-cli](https://github.com/ankitpokhrel/jira-cli)

*Description*: 🔥 [WIP] Feature-rich interactive Jira command line.

*Labels*: #Go #JIRA

*Links*:
  - https://dev.to/konsole/jira-cli-the-missing-command-line-tool-for-atlassian-jira-2n2i

*Examples*:

```bash
# List issues that I am watching in the current board
$ jira issue list -w

# List issues assigned to me
$ jira issue list -a$(jira me)

# List issues created within an hour and updated in the last 30 minutes️
$ jira issue list --created -1h --updated -30m

# Give me issues that are of high priority, is in progress, was created this month, and has given labels 🔥
$ jira issue list -yHigh -s"In Progress" --created month -lbackend -l"high prio"

# Wait, what was that ticket I opened earlier today? 😫
$ jira issue list --history # What was the first issue I ever reported on the current board? 🤔
$ jira issue list -r$(jira me) --reverse

# What was the first bug I ever fixed in the current board? 🐞
$ jira issue list -a$(jira me) -tBug sDone -rFixed --reverse

# What issues did I report this week? 🤷‍♂️
$ jira issue list -r$(jira me) --created week
```
