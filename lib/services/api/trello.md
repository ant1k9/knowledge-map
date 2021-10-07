[Trello Webhooks](https://developer.atlassian.com/cloud/trello/guides/rest-api/webhooks/)

*Description*: Webhooks provide a way for application developers to receive notifications when a model changes.

*Labels*: #Trello

*Examples*:

```
$ curl -X POST -H "Content-Type: application/json" \
	https://api.trello.com/1/tokens/{APIToken}/webhooks/ \
	-d '{
	  "key": "{APIKey}",
	  "callbackURL": "http://www.mywebsite.com/trelloCallback",
	  "idModel":"4d5ea62fd76aa1136000000c",
	  "description": "My first webhook"
	}'
$ curl "https://trello.com/1/boards/<board_id>/lists"  # get lists on a board
$ curl "https://api.trello.com/1/boards/P3bsnKz1/customFields?key=$TRELLO_API_KEY&token=$TRELLO_TOKEN"
```
