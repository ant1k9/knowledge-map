[glab](https://gitlab.com/gitlab-org/cli)

*Description*: An open-source GitLab command line tool bringing GitLab's cool features to your command line

*Docs*: https://glab.readthedocs.io/

*Labels*: #Gitlab #Git #Go

*Examples*:

```bash
# create mr from command line
$ glab mr create \
    --title "MESSAGE" \
    --description "DESCRIPTION" \
    --remove-source-branch \
    --target-branch "BRANCH"

# check gitlab pipeline status
$ glab ci status
```
