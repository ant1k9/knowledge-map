[MeiliSearch](https://github.com/meilisearch/MeiliSearch)

*Description*: Powerful, fast, and an easy to use search engine

*Labels*: #Rust #SearchEngine

*Docs*: https://docs.meilisearch.com/learn/getting_started/quick_start.html#search

*Examples*:

```bash
$ curl -L 'https://bit.ly/2PAcw9l' -o movies.json
$ curl -i -X POST 'http://127.0.0.1:7700/indexes/movies/documents' \
    --header 'content-type: application/json' \
    --data-binary @movies.json
$ curl 'http://127.0.0.1:7700/indexes/movies/search?q=botman+robin&limit=2' | jq
```
