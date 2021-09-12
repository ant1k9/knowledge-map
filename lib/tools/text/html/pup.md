[pup](https://github.com/ericchiang/pup)

*Description*: pup is a command line tool for processing HTML. It reads from stdin, prints to stdout, and allows the user to filter parts of the page using CSS selectors.

*Labels*: #Go #HTML

*Examples*:

```bash
$ curl -s https://news.ycombinator.com/ | pup 'table table tr:nth-last-of-type(n+2) td.title a'
$ curl -s https://news.ycombinator.com/ | pup 'table table tr:nth-last-of-type(n+2) td.title a attr{href}'
$ curl -s https://news.ycombinator.com/ | pup 'table table tr:nth-last-of-type(n+2) td.title a json{}'
```
