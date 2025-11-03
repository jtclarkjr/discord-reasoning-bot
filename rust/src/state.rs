use async_openai::Client as OpenAIClient;
use serenity::prelude::Client;
use std::sync::Arc;
use tokio::sync::RwLock;

#[derive(Clone)]
pub struct AppState {
    pub bot_client: Arc<RwLock<Option<Client>>>,
    pub openai_client: OpenAIClient<async_openai::config::OpenAIConfig>,
    pub bot_token: String,
}
