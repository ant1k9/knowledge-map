[spago](https://github.com/nlpodyssey/spago)

*Description*: Self-contained Machine Learning and Natural Language Processing library in Go

*Labels*: #Go #MachineLearning #NLP

*Examples*:

```bash
$ ./bert-server server --repo=~/.spago --model=deepset/bert-base-cased-squad2 --tls-disable
$ PASSAGE="BERT is a technique for NLP developed by Google. BERT was created and published in 2018 by Jacob Devlin and his colleagues from Google." \
    QUESTION1="Who is the author of BERT?" \
    QUESTION2="When was BERT created?" \
    curl -k -d '{"question": "'"$QUESTION1"'", "passage": "'"$PASSAGE"'"}' \
		-H "Content-Type: application/json" \
		"http://127.0.0.1:1987/answer?pretty"
```
