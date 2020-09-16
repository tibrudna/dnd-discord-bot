package function

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Spell(message *discordgo.MessageCreate) (string, error) {
	contentParts := strings.Split(message.Content, " ")
	if len(contentParts) != 3 {
		return "", errors.New("To many Arguments")
	}

	expression, _ := regexp.Compile("^([[:alpha:]]+-)*[[:alpha:]]+$")
	if !expression.MatchString(contentParts[2]) {
		return "", errors.New("False argument format")
	}

	url := "http://dnd5e.wikidot.com/spell:" + contentParts[2]
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	if response.StatusCode == 404 {
		return "Spell not known", nil
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	expression, _ = regexp.Compile("(?s)<p>.*<\\/p>")
	subString := expression.FindAllStringSubmatch(string(body), -1)[0][0]

	subString = strings.Replace(subString, "<strong>", "**", -1)
	subString = strings.Replace(subString, "</strong>", "**", -1)

	expression, _ = regexp.Compile("(<.{0,5}>)*(<a.*\">)*")
	subString = expression.ReplaceAllString(subString, "")

	header := strings.Title(contentParts[2])
	header = strings.Replace(header, "-", " ", -1)

	subString = "**" + header + "**\n\n" + subString

	return subString, nil
}
