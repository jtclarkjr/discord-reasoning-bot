use async_openai::{types::CreateChatCompletionRequestArgs, Client as OpenAIClient};

pub async fn is_message_offensive(
    client: &OpenAIClient<async_openai::config::OpenAIConfig>,
    message: &str,
) -> Result<bool, Box<dyn std::error::Error + Send + Sync>> {
    let request = CreateChatCompletionRequestArgs::default()
        .model("gpt-5")
        .messages(vec![async_openai::types::ChatCompletionRequestMessage::User(
            async_openai::types::ChatCompletionRequestUserMessageArgs::default()
                .content(format!(
                    "Is the following message offensive? Answer with \"true\" or \"false\" and no period:\n\n\"{}\"",
                    message
                ))
                .build()?,
        )])
        .build()?;

    let response = client.chat().create(request).await?;

    if let Some(choice) = response.choices.first() {
        if let Some(content) = &choice.message.content {
            return Ok(content.trim().to_lowercase() == "true");
        }
    }

    Ok(false)
}
