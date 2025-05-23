import 'dotenv/config'
import express from 'express'
import morgan from 'morgan'
import OpenAI from 'openai'
import { Client, GatewayIntentBits, Message, TextChannel } from 'discord.js'

// Initialize OpenAI client
const openai = new OpenAI({
  apiKey: process.env.OPENAI_API_KEY
})

// Global Discord client variable
let client: Client | null = null

/**
 * Checks if a message is offensive using OpenAI's API.
 * @param messageContent - The content of the message to check.
 * @returns A promise that resolves to a boolean indicating if the message is offensive.
 */
async function isMessageOffensive(messageContent: string): Promise<boolean> {
  try {
    const response = await openai.chat.completions.create({
      model: 'o3-mini', // Ensure this model is available to your account or will error
      messages: [
        {
          role: 'user',
          content: `Is the following message offensive? Answer with "true" or "false" and no period:\n\n"${messageContent}"`
        }
      ]
    })

    const answer = response.choices[0]?.message?.content?.trim().toLowerCase()
    return answer === 'true'
  } catch (error) {
    console.error('Error checking message:', error)
    return false
  }
}

/**
 * Starts the Discord bot.
 */
async function startBot(): Promise<string> {
  if (client) {
    return 'Bot is already running.'
  }

  client = new Client({
    intents: [
      GatewayIntentBits.Guilds,
      GatewayIntentBits.GuildMessages,
      GatewayIntentBits.MessageContent
    ]
  })

  // Listen for messages
  client.on('messageCreate', async (message: Message) => {
    // Ignore messages from bots
    if (message.author.bot) return

    const offensive = await isMessageOffensive(message.content)
    if (offensive) {
      try {
        await message.delete()
        if (message.channel.isTextBased()) {
          const textChannel = message.channel as TextChannel
          await textChannel.send(
            `<@${message.author.id}>, your message was removed due to offensive content.`
          )
        }
      } catch (error) {
        console.error('Failed to delete message or notify user:', error)
      }
    }
  })

  client.once('ready', () => {
    console.log(`Logged in as ${client?.user?.tag}!`)
  })

  try {
    await client.login(process.env.DISCORD_BOT_TOKEN)
    return 'Bot started successfully.'
  } catch (error) {
    console.error('Error logging in:', error)
    client = null
    return 'Failed to start bot.'
  }
}

/**
 * Stops the Discord bot.
 */
async function stopBot(): Promise<string> {
  if (!client) {
    return 'Bot is not running.'
  }
  try {
    await client.destroy()
    client = null
    return 'Bot stopped successfully.'
  } catch (error) {
    console.error('Error stopping bot:', error)
    return 'Failed to stop bot.'
  }
}

/**
 * Express REST endpoints
 */

const app = express()

// HTTP request logging
app.use(morgan('combined'))

const port = process.env.PORT || 8080

// Endpoint to start the bot
app.post('/bot/on', async (_req, res) => {
  const result = await startBot()
  res.status(200).send(result)
})

// Endpoint to stop the bot
app.post('/bot/off', async (_req, res) => {
  const result = await stopBot()
  res.status(200).send(result)
})

app.listen(port, () => {
  console.log(`Server listening on port ${port}`)
})
