import 'dotenv/config'
import { Client, GatewayIntentBits, Message, TextChannel } from 'discord.js'
import OpenAI from 'openai'

// Initialize OpenAI client
const openai = new OpenAI({
  apiKey: process.env.OPENAI_API_KEY
})

// Initialize Discord client
const client = new Client({
  intents: [
    GatewayIntentBits.Guilds,
    GatewayIntentBits.GuildMessages,
    GatewayIntentBits.MessageContent
  ]
})

/**
 * Checks if a message is offensive using OpenAI's API.
 * @param messageContent - The content of the message to check.
 * @returns A promise that resolves to a boolean indicating if the message is offensive.
 */
async function isMessageOffensive(messageContent: string): Promise<boolean> {
  try {
    const response = await openai.chat.completions.create({
      model: 'o3-mini', // Ensure this model is available to your account or will get error
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

// Event listener for new messages
client.on('messageCreate', async (message: Message) => {
  if (message.author.bot) return // Ignore bot messages

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
    } catch (err) {
      console.error('Failed to delete message or notify user:', err)
    }
  }
})

// Bot login
client.once('ready', () => {
  console.log(`Logged in as ${client.user?.tag}!`)
})

client.login(process.env.DISCORD_BOT_TOKEN)
