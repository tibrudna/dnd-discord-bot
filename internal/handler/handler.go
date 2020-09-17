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
	action, ok := function.Functions[command]

	if !ok {
		action, _ = function.Functions["help"]
		returnString, _ := action(message)
		session.ChannelMessageSend(message.ChannelID, "Command not known\n\n"+returnString)
	}

	result, err := action(message)
	if err != nil {
		action, _ = function.Functions["help"]
		returnString, _ := action(message)
		session.ChannelMessageSend(message.ChannelID, returnString)
	}

	session.ChannelMessageSend(message.ChannelID, result)
}
