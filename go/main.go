package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/jtclarkjr/router-go"
	"github.com/jtclarkjr/router-go/middleware"
	"github.com/openai/openai-go/v2"
)

var sess *discordgo.Session

// BotHandler struct to hold dependencies
type BotHandler struct {
	client *openai.Client
}

// NewBotHandler initializes a new bot handle
func NewBotHandler(client *openai.Client) *BotHandler {
	return &BotHandler{client: client}
}

func main() {
	r := router.NewRouter()
	r.Use(middleware.Logger)

	apiKey := os.Getenv("OPENAI_API_KEY")
	botToken := os.Getenv("DISCORD_BOT_TOKEN")

	r.Use(middleware.EnvVarChecker("OPENAI_API_KEY", "DISCORD_BOT_TOKEN"))

	client := initOpenAIClient(apiKey)
	botHandler := NewBotHandler(&client)

	// Initialize Discord session
	var err error
	sess, err = discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal(err)
	}

	// Register message handler immediately, independent of /bot/on calls
	sess.AddHandler(botHandler.messageHandler)

	// Set intents to receive message events
	sess.Identify.Intents = discordgo.IntentsGuildMessages

	r.Post("/bot/on", botHandler.startBotHandler)
	r.Post("/bot/off", botHandler.stopBotHandler)

	// Start HTTP server in a separate goroutine
	go func() {
		fmt.Println("Starting HTTP server on :8080")
		if err := http.ListenAndServe(":8080", r); err != nil {
			log.Fatal(err)
		}
	}()

	// Wait for termination signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

// messageHandler continuously monitors messages while the bot is running
func (b *BotHandler) messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check if the message is offensive
	offensive, err := isMessageOffensive(b.client, m.Content)
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
	}
}

// startBotHandler starts the bot connection
func (b *BotHandler) startBotHandler(w http.ResponseWriter, r *http.Request) {
	err := sess.Open()
	if err != nil {
		http.Error(w, "Failed to start bot: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Bot started successfully.")
}

// stopBotHandler stops the bot connection
func (b *BotHandler) stopBotHandler(w http.ResponseWriter, r *http.Request) {
	err := sess.Close()
	if err != nil {
		http.Error(w, "Failed to stop bot: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Bot stopped successfully.")
}
