package function

import (
	"github.com/bwmarrin/discordgo"
)

func Help(message *discordgo.MessageCreate) (string, error) {
	rollHelp := "***roll*** performs a roll. **Example:** roll 2d4+1 1d10-1\n"
	spellHelp := "***spell*** returns information about a spell. **Example:** spell tensers-floating-disk\n"
	helpHelp := "***help*** returns this information.\n\n"
	github := "For more information see https://github.com/tibrudna/dnd-discord-bot"

	return rollHelp + spellHelp + helpHelp + github, nil
}
