mod api;
mod handlers;
mod services;
mod state;

use async_openai::Client as OpenAIClient;
use axum::{routing::post, Router};
use std::{env, sync::Arc};
use tokio::sync::RwLock;
use tracing::info;

use api::routes::{start_bot, stop_bot};
use state::AppState;

#[tokio::main]
async fn main() {
    dotenv::dotenv().ok();

    tracing_subscriber::fmt::init();

    let openai_api_key = env::var("OPENAI_API_KEY").expect("OPENAI_API_KEY must be set");
    let bot_token = env::var("DISCORD_BOT_TOKEN").expect("DISCORD_BOT_TOKEN must be set");

    let openai_client = OpenAIClient::with_config(
        async_openai::config::OpenAIConfig::default().with_api_key(openai_api_key),
    );

    let app_state = AppState {
        bot_client: Arc::new(RwLock::new(None)),
        openai_client,
        bot_token,
    };

    let app = Router::new()
        .route("/bot/on", post(start_bot))
        .route("/bot/off", post(stop_bot))
        .with_state(app_state);

    let listener = tokio::net::TcpListener::bind("0.0.0.0:8080")
        .await
        .expect("Failed to bind to port 8080");

    info!("Starting HTTP server on :8080");

    axum::serve(listener, app)
        .await
        .expect("Failed to start server");
}
