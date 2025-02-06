package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// Load OpenAI API key from environment variable
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OpenAI API key is not set")
	}

	// Initialize OpenAI client
	client := initOpenAIClient(apiKey)

	// Load Discord bot token from environment variable
	botToken := os.Getenv("DISCORD_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("Discord bot token is not set")
	}

	// Initialize Discord session
	sess, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal(err)
	}

	// Add message handler
	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		// Check if the message is offensive
		offensive, err := isMessageOffensive(client, m.Content)
		if err != nil {
			log.Println("Error checking message:", err)
			return
		}

		if offensive {
			// Delete the offensive message
			err := s.ChannelMessageDelete(m.ChannelID, m.ID)
			if err != nil {
				log.Println("Failed to delete message:", err)
			}

			// Notify the user about the blocked message
			warningMessage := fmt.Sprintf("<@%s>, your message was blocked because it contained offensive content.", m.Author.ID)
			_, err = s.ChannelMessageSend(m.ChannelID, warningMessage)
			if err != nil {
				log.Println("Failed to send warning message:", err)
			}

			return
		}

		// Additional logic for non-offensive messages
		fmt.Println(m.Content)
	})

	// Set intents to receive message events
	sess.Identify.Intents = discordgo.IntentsGuildMessages

	// Open Discord session
	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("The bot is online!")

	// Wait for a termination signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
