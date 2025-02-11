# Discord reasoning bot

Discord bot utilizing OpenAI api for reasoning

## Installation

Go

```bash
go mod tidy
```

Node

```bash
brew install nvm
pnpm i
```

If don't have auto-install script for nvm can run

```bash
nvm install 23.7.0
nvm use 23.7.0
```

## Usage

Base example uses reasoning model to determine if the input from the user is offensive. Then from discord side the bot removes (Admin access) the message and notifies the user

### Go

#### Local

```bash
go run .
```

#### Docker

```bash
export OPENAI_API_KEY=
export DISCORD_BOT_TOKEN=
eval "$(direnv hook zsh)"

docker compose up
```

Turn on/off

```curl
curl -X POST http://localhost:3000/bot/on
curl -X POST http://localhost:3000/bot/off
```

### Node

#### Local

```bash
pnpm start
```

#### Docker

```bash
export OPENAI_API_KEY=
export DISCORD_BOT_TOKEN=
eval "$(direnv hook zsh)"

docker compose up
```

Turn on/off

```curl
curl -X POST http://localhost:3000/bot/on
curl -X POST http://localhost:3000/bot/off
```

If you want to do on/off by chat inputs the most clear and straightforward solution is to make another bot that controls the desired bot to on/off. However, this controller bot will need to be on. This is good for:

- If the controller bot is always on
- Have multiple bots needing to switch on/off
- Avoiding having to do a API request to turn on/off
- Can even set something like roleIds to have permissions only for `!on {bot.tag}` etc

Another solution would to make use of process manager (pm2). There are many ways to play around with the base code here

## Source code

- Go
- Node (TypeScript)

## OpenAI

You will need to add a secret key to use and have access to reasoning models

## Discord

You will need to set up a project and add a bot with the token and permissions to that project in the Discord developer portal

Add a `config.json` file to source code folder

```bash

{
  "token": "", // Bot token ("Bot TOKEN" format)
  "BotPrefix": "" // This is the permission code generated
}
```

## Env vars

- `OPENAI_API_KEY`
- `DISCORD_BOT_TOKEN`

## License

[MIT](https://choosealicense.com/licenses/mit/)
