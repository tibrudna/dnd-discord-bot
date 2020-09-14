//Package function provides functions to for common tasks
package function

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

//Functions maps all exported functions to the name as a string
var Functions = map[string]func(message *discordgo.MessageCreate) (string, error){
	"spell": Spell,
	"roll":  Roll,
}

//Spell returns information about a spell as a string.
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

//Roll returns the result of one or multiple rolls
func Roll(message *discordgo.MessageCreate) (string, error) {
	contentParts := strings.Split(message.Content, " ")
	if len(contentParts) < 3 {
		return "", errors.New("Not enough arguments")
	}

	sum := 0
	text := "**" + message.Author.Username + "** rolled:\n"
	argExpression, _ := regexp.Compile("^[1-9]\\d*d[1-9]\\d*((\\+|-)\\d+){0,1}")
	submatchExpression, _ := regexp.Compile("(\\+|-){0,1}\\d+")

	for _, part := range contentParts[2:] {
		if !argExpression.MatchString(part) {
			return "", errors.New("Bad dice expression format")
		}

		argSubmatches := submatchExpression.FindAllStringSubmatch(part, -1)
		count, _ := strconv.Atoi(argSubmatches[0][0])
		point, _ := strconv.Atoi(argSubmatches[1][0])

		for i := 0; i < count; i++ {
			result := rand.Intn(point) + 1
			sum += result
			text += strconv.Itoa(result) + "\n"
		}

		if len(argSubmatches) == 3 {
			plus, _ := strconv.Atoi(argSubmatches[2][0])
			sum += plus
			text += argSubmatches[2][0] + "\n"
		}

		text += "-----\n"
	}

	text += strconv.Itoa(sum)

	return text, nil
}
