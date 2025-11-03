use axum::{extract::State, http::StatusCode, response::IntoResponse};
use serenity::prelude::{Client, GatewayIntents};
use tracing::error;

use crate::handlers::discord::Handler;
use crate::state::AppState;

pub async fn start_bot(State(state): State<AppState>) -> impl IntoResponse {
    let mut bot_guard = state.bot_client.write().await;

    if bot_guard.is_some() {
        return (StatusCode::OK, "Bot is already running.");
    }

    let intents = GatewayIntents::GUILD_MESSAGES | GatewayIntents::MESSAGE_CONTENT;

    let handler = Handler {
        openai_client: state.openai_client.clone(),
    };

    match Client::builder(&state.bot_token, intents)
        .event_handler(handler)
        .await
    {
        Ok(mut client) => {
            if let Err(e) = client.start().await {
                error!("Failed to start bot: {:?}", e);
                return (StatusCode::INTERNAL_SERVER_ERROR, "Failed to start bot.");
            }
            *bot_guard = Some(client);
            (StatusCode::OK, "Bot started successfully.")
        }
        Err(e) => {
            error!("Failed to build bot client: {:?}", e);
            (StatusCode::INTERNAL_SERVER_ERROR, "Failed to start bot.")
        }
    }
}

pub async fn stop_bot(State(state): State<AppState>) -> impl IntoResponse {
    let mut bot_guard = state.bot_client.write().await;

    if let Some(client) = bot_guard.take() {
        client.shard_manager.shutdown_all().await;
        (StatusCode::OK, "Bot stopped successfully.")
    } else {
        (StatusCode::OK, "Bot is not running.")
    }
}
