package core

import (
	"regexp"

	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/HellOrHeavenBot/db"
	"github.com/tecnologer/HellOrHeavenBot/model"
	"github.com/tecnologer/HellOrHeavenBot/resources"
	bot "github.com/yanzay/tbot"
)

//hasCustomResponse checks if the text match with some regex
func hasCustomResponse(msg *bot.Message) {
	if msg.Text == "" {
		return
	}

	readyResponse := make([]*model.Response, 0)
	defer func() {
		if r := recover(); r != nil {
			log.WithFields(log.Fields{
				"Responses": readyResponse,
				"text":      msg.Text,
				"chatID":    msg.Chat.ID,
			}).Debug(r)
		}
	}()

	responses, err := db.RetrieveCustomResponses(msg.Chat.ID)

	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"text":   msg.Text,
			"chatID": msg.Chat.ID,
		}).Error("error when try retrive custom responses")
		return
	}

	if len(responses) == 0 {
		return
	}

	for _, res := range responses {
		regex, err := regexp.Compile(res.Regex)
		if err != nil {
			log.WithError(err).WithFields(log.Fields{
				"text":   msg.Text,
				"chatID": msg.Chat.ID,
				"regex":  res.Regex,
			}).Error("error when try compile the regex")
			continue
		}

		if regex.Match([]byte(msg.Text)) {
			readyResponse = append(readyResponse, &model.Response{Content: res.Response, Type: res.ResponseType})
		}
	}

	if len(readyResponse) == 0 {
		return
	}

	if len(readyResponse) == 1 {
		SendResponse(msg, readyResponse[0])
		return
	}

	selection := resources.GetRandomIntFromRange(0, len(readyResponse)-1)

	SendResponse(msg, readyResponse[selection])
}

func AddCustomResponse(cRes *model.CustomResponse) error {
	_, err := regexp.Compile(cRes.Regex)

	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"CustomResponse": cRes,
		}).Error("error when try compile the regex")
		return err
	}

	return db.InsertCustomResponse(cRes)
}
