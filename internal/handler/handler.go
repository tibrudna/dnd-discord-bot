//Package handler provides functions for handling events.
package handler

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/tibrudna/dnd-discord-bot/internal/function"
)

//MessageCreate handles a MessageCreateEvent.
func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if session.State.User == message.Author {
		return
	}

	if !strings.HasPrefix(message.Content, os.Getenv("PREFIX")) {
		return
	}

	command := strings.Split(message.Content, " ")[1]
	function, ok := function.Functions[command]

	if !ok {
		session.ChannelMessageSend(message.ChannelID, "Command not known")
	}

	result, err := function(message)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "there was a problem")
	}

	session.ChannelMessageSend(message.ChannelID, result)
}
