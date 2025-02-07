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

```bash
go run .
```

Turn on/off

```curl
curl -X POST http://localhost:3000/bot/on
curl -X POST http://localhost:3000/bot/off
```

### Node

```bash
pnpm start
```

Turn on/off

```curl
curl -X POST http://localhost:3000/bot/on
curl -X POST http://localhost:3000/bot/off
```

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
