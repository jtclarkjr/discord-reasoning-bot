# Discord reasoning bot

Discord bot utilizing OpenAI api for reasoning

## Installation

```bash
go mod tidy
```

## Usage

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

Add a `config.json` file
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
