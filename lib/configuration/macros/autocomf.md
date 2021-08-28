[autocomf](https://github.com/athas/autocomf)

*Description*: AUTOCOMF is a programming language that presents a configuration UI based on annotations in a source file. The configuration specified by the user is then incorporated by automatically modifying the source file.

*Labels*: #Haskell #Configuration

*Examples*:

```
# AUTOCOMF VAR myForegroundColour REGEX red|green|blue|purple|cyan|white|yellow
# AUTOCOMF USE fg=$¤myForegroundColour¤
fg=$red
```
