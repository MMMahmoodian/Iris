# Alarm
## Intro
This is a simple service that works as a middleware for sending messages to telegram.
## Setup
`git clone https://github.com/MMMahmoodian/alarm.git` 

`cd alarm`

`cp .env.example .env`

Open `.env` file and replace your bot token in `TELEGRAM_BOT_API_KEY`. Save `.env` and exit.

`docker compose up -d`

## Health check
`curl -f {YOUR_SERVER_IP}:8080/telegram/ping`

This should return:
```json
{
    "errors": "",
    "messages": "pong"
}
```