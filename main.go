package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.SetPrefix("discord-bot: ")
	log.Println("Starting bot...")

	token := os.Getenv("token")

	log.Println("Creating Discord session...")
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
		return
	}

	session.Identify.Intents = discordgo.IntentsAll

	err = session.Open()
	if err != nil {
		log.Fatal("Error opening connection: ", err)
		return
	}

	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "ping" {
			_, err := s.ChannelMessageSend(m.ChannelID, "pong")
			if err != nil {
				log.Println("Error sending message: ", err)
				return
			}
		}
	})

	err = session.UpdateGameStatus(0, "ping pong")
	if err != nil {
		log.Println("Error setting game status: ", err)
		return
	}

	log.Println("Bot is now running. Press CTRL-C to exit.")

	defer func(session *discordgo.Session) {
		log.Println("Closing connection...")

		err := session.Close()
		if err != nil {
			log.Println("Error closing connection: ", err)
			return
		}
	}(session)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
