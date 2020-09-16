package function

import (
	"errors"
	"math/rand"
	"regexp"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

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
