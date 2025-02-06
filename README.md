# Discord reasoning bot

Discord bot utilizing OpenAI api for reasoning

## Installation

```bash
go mod tidy
```

## Usage

Base example uses reasoning model to determine if the input from the user is offensive. Then from discord side the bot removes (Admin access) the message and notifies the user
```bash
go run .
```

## Source code
- Go
- Typescript (add later)

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
