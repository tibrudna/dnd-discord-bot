package function_test

import (
	"testing"

	"github.com/tibrudna/dnd-discord-bot/internal/function"

	"github.com/bwmarrin/discordgo"
)

var mockUser = discordgo.User{Username: "Bob"}
var mockMessage = discordgo.Message{
	Author:  &mockUser,
	Content: "!w roll 2d10",
}
var mockMessageCreate = discordgo.MessageCreate{Message: &mockMessage}

func TestRollParameter(t *testing.T) {
	mockMessage.Content = "!w roll"
	t.Run("Returns error when too few parameters", testParameterReturnsError)

	mockMessage.Content = "!w roll 2w10"
	t.Run("Returns Error when dice not match expression", testParameterReturnsError)

	mockMessage.Content = "!w roll 0d10"
	t.Run("Returns error when zero in front dice format", testParameterReturnsError)
	mockMessage.Content = "!w roll 1d0"
	t.Run("Returns error when zero in back dice format", testParameterReturnsError)

	mockMessage.Content = "!w roll 1d10"
	t.Run("Returns no error when correct parameter", testReturnsStringWhenCorrectParameter)
	mockMessage.Content = "!w roll 1d10+2"
	t.Run("Returns no error when plus parameter", testReturnsStringWhenCorrectParameter)
	mockMessage.Content = "!w roll 1d10+2 1d10+1"
	t.Run("Returns no error when multiple dices", testReturnsStringWhenCorrectParameter)
}

func testReturnsStringWhenCorrectParameter(t *testing.T) {
	result, err := function.Roll(&mockMessageCreate)
	if err != nil {
		t.Error("An error was returned")
	}
	if result == "" {
		t.Error("No result was returned")
	}
}

func testParameterReturnsError(t *testing.T) {
	_, err := function.Roll(&mockMessageCreate)
	if err == nil {
		t.Error("No error returned")
	}
}

func TestSpellParameter(t *testing.T) {
	mockMessage.Content = "!w spell tensers floating disk"
	t.Run("Returns error when too many parameter", testSpellParameterReturnsError)
	mockMessage.Content = "!w spell tensers_floating_disk"
	t.Run("Retruns error when wrong format", testSpellParameterReturnsError)
}

func testSpellParameterReturnsError(t *testing.T) {
	_, err := function.Spell(&mockMessageCreate)
	if err == nil {
		t.Error("No error returned")
	}
}
