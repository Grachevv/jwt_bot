# JWT_bot
Simple telegram bot for decoding JWT.

Bot lives here:  https://t.me/ssg_jwt_bot

## How to start

First, you need to register a new bot. You can read about this here: https://core.telegram.org/bots#3-how-do-i-create-a-bot.
Finally, you will get API token.

Run docker container with a bot:
```shell
docker build -t jwt-bot:latest . 
docker run --name jwt-bot -d -e BOT_TOKEN={your api token} jwt-bot
```
