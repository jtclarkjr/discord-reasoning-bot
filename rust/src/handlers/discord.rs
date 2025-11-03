use async_openai::Client as OpenAIClient;
use serenity::{
    async_trait,
    model::{channel::Message, gateway::Ready},
    prelude::*,
};
use tracing::{error, info};

use crate::services::moderation::is_message_offensive;

pub struct Handler {
    pub openai_client: OpenAIClient<async_openai::config::OpenAIConfig>,
}

#[async_trait]
impl EventHandler for Handler {
    async fn message(&self, ctx: Context, msg: Message) {
        if msg.author.bot {
            return;
        }

        if msg.content.trim().is_empty() {
            return;
        }

        match is_message_offensive(&self.openai_client, &msg.content).await {
            Ok(true) => {
                if let Err(e) = msg.delete(&ctx.http).await {
                    error!("Failed to delete message: {:?}", e);
                    return;
                }

                let warning = format!(
                    "<@{}>, your message was blocked because it contained offensive content.",
                    msg.author.id
                );

                if let Err(e) = msg.channel_id.say(&ctx.http, warning).await {
                    error!("Failed to send warning message: {:?}", e);
                }
            }
            Ok(false) => {}
            Err(e) => {
                error!("Error checking message: {:?}", e);
            }
        }
    }

    async fn ready(&self, _: Context, ready: Ready) {
        info!("{} is connected!", ready.user.name);
    }
}
