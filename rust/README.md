# Discord Reasoning Bot - Rust Implementation

A Discord bot written in Rust that monitors messages for offensive content using OpenAI's API and automatically removes them.

## Features

- HTTP API server using Axum
- Discord integration via Serenity
- OpenAI API integration for content moderation
- Async runtime with Tokio
- Start/stop bot via REST endpoints

## Prerequisites

- Rust (latest stable version)
- Discord bot token
- OpenAI API key

## Setup

1. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```

2. Edit `.env` and add your credentials:
   ```
   OPENAI_API_KEY=your_openai_api_key_here
   DISCORD_BOT_TOKEN=your_discord_bot_token_here
   ```

3. Build the project:
   ```bash
   cargo build --release
   ```

## Running

Start the server:
```bash
cargo run --release
```

The HTTP server will start on port 8080.

## API Endpoints

### Start the bot
```bash
curl -X POST http://localhost:8080/bot/on
```

### Stop the bot
```bash
curl -X POST http://localhost:8080/bot/off
```

## How it works

1. The bot connects to Discord when started via the `/bot/on` endpoint
2. It monitors all messages in channels it has access to
3. Each message is analyzed by OpenAI to determine if it contains offensive content
4. If offensive content is detected, the message is deleted and a warning is sent to the user
5. The bot can be stopped via the `/bot/off` endpoint

## Dependencies

- `serenity` - Discord API client
- `axum` - HTTP web framework
- `tokio` - Async runtime
- `async-openai` - OpenAI API client
- `dotenv` - Environment variable management
- `tracing` - Logging framework
