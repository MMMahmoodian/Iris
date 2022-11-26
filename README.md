# Iris
<img src="images/Iris.jpeg" alt="Iris goddess of rainbow and messenger of gods" style="width:400px;"/>

## Intro
Iris was the goddess of the rainbow and the messenger of the Olympian gods. But now she is here to aid us in the journey of sending alarms to our customers.
## Setup
First clone this repository and cd into project root.<br>
```shell
git clone https://github.com/MMMahmoodian/alarm.git
cd alarm
```
Then we need to add `.env` file
```shell
cp src/.env.example src/.env
```
After that we should add telegram bot configs in `.env`. Open `.env` and add telegram token as `TELEGRAM_BOT_API_KEY`.
```
...

TELEGRAM_BOT_API_URL=https://api.telegram.org/bot
TELEGRAM_BOT_API_KEY={api_key}

...
```
Now we need to build and run the project:
```shell
docker compose build 
docker compose up -d
```

## Health check
`curl -f {YOUR_SERVER_IP}:8088/telegram/ping`

This should return response below with `200 OK` status code:
```json
{
    "errors": "",
    "messages": "pong"
}
```