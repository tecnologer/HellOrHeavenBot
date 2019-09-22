package core

import (
	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/HellOrHeavenBot/db"
	bot "github.com/yanzay/tbot"
)

//Cancel cancels the current task
func Cancel(msg *bot.Message) {
	if HasUserIncompleteRes(msg.From) {
		identifier := getResponseIdentifier(msg.From)
		delete(incompleteResponse, identifier)
	}

	res, err := db.GetResponseByCommand(AcceptedCommands.GetID("cancel"), msg.From.LanguageCode)
	if err != nil {
		log.WithError(err).Info("error when try get response for cancel command")
		sendText(msg, cLang["genericCancel"])
	} else {
		SendResponse(msg, res)
	}
}
