//Package function provides functions to for common tasks
package function

import (
	"github.com/bwmarrin/discordgo"
)

//Functions maps all exported functions to the name as a string
var Functions = map[string]func(message *discordgo.MessageCreate) (string, error){
	"spell": Spell,
	"roll":  Roll,
	"help":  Help,
}
