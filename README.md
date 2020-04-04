# Rest Chat
## Application Set-up
* Run the application.

`cd rest-chat/src`

`godep go install ./…`

`go run main.go`

* Database Set-up

Script to create the tables `rest-chat/src/db/v1.sql`

## Requests

* Health Check

```
curl --location --request GET '127.0.0.1:8080/check'
```

* New User

```
curl --location --request POST '127.0.0.1:8080/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "test",
    "password": "test"
}'
```

* Log In

```
curl --location --request POST '127.0.0.1:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "test",
    "password": "test"
}'
```

Save the Auth Token because you’ll need it.

* Get Messages

```
curl --location --request GET '127.0.0.1:8080/messages?recipient=2&start=1' \
--header 'Content-Type: application/json' \
--header 'Authorization: {{auth_token}}’ \
--data-raw ''
```

1. Recipient: user ID of recipient.
2. Start: starting message ID.

* Post New Message

```
curl --location --request POST '127.0.0.1:8080/messages' \
--header 'Content-Type: application/json' \
--header 'Authorization: {{auth_token}}’ \
--data-raw '{
    "sender": 1,
    "recipient": 2,
    "content": {
        "type": "string",
        "test": "message text"
    }
}'
```
