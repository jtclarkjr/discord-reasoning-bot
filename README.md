# Discord reasoning bot

Discord bot utilizing OpenAI api for reasoning

## Installation

Go (1.24.0)

```bash
go mod tidy
```

Node (24.0.0)

```bash
brew install nvm
pnpm i
```

Bun (lastest)

```bash
brew install bun
bun i
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

docker compose up --build
```

Turn on/off

```curl
curl -X POST http://localhost:8080/bot/on
curl -X POST http://localhost:8080/bot/off
```

### Node

#### Local

```bash
pnpm dev
```

#### Docker

```bash
export OPENAI_API_KEY=
export DISCORD_BOT_TOKEN=
eval "$(direnv hook zsh)"

docker compose up --build
```

### Bun

#### Local

```bash
bun dev
```

#### Docker

```bash
export OPENAI_API_KEY=
export DISCORD_BOT_TOKEN=
eval "$(direnv hook zsh)"

docker compose up --build
```

Turn on/off

```curl
curl -X POST http://localhost:8080/bot/on
curl -X POST http://localhost:8080/bot/off
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
- Bun (TypeScript)

## OpenAI

You will need to add a secret key to use and have access to reasoning models
- Code uses o3 reasoning model but needs verified account or 4-5 tier
- Tier 1-3 is reccommend to use o4‑mini reason model instead since it's the available option

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
