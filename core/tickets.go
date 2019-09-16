package core

import (
	"fmt"
	"strings"

	"github.com/tecnologer/HellOrHeavenBot/resources"

	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/HellOrHeavenBot/db"
	m "github.com/tecnologer/HellOrHeavenBot/model"
	bot "github.com/yanzay/tbot"
)

//Hell is the function in charge of registering the hell tickets
func Hell(msg *bot.Message) {
	go updateStats(msg, m.StatsHell)
}

//Heaven is the function in charge of registering the hell tickets
func Heaven(msg *bot.Message) {
	go updateStats(msg, m.StatsHeaven)
}

//GetStats gets the count of tickets for the user who requested
func GetStats(msg *bot.Message) {
	stats := db.GetStats(strings.ToLower(msg.From.Username))

	if stats == nil {
		sendText(msg, "vas al infierno de todas maneras")
		return
	}

	sendText(msg, fmt.Sprintf("Hell: %d, Heaven: %d", stats.Hell, stats.Heaven))
}

//updateStats calls the functions to update database
func updateStats(msg *bot.Message, t m.StatsType) {
	doomedName := getDoomedName(msg.Text)

	if doomedName == "" {
		sendText(msg, cLang["ticketsNameRequired"])
		return
	}
	err := db.InsertStat(doomedName, t)

	if err != nil {
		log.Println(err)
		sendText(msg, cLang["genericFail"])
		return
	}

	//command ID
	var cmdID int
	if t == m.StatsHeaven {
		cmdID = AcceptedCommands.GetID("heaven")
	} else {
		cmdID = AcceptedCommands.GetID("hell")
	}

	res, err := db.GetResponseByCommand(cmdID, msg.From.LanguageCode)
	if err != nil {
		log.Println(err)
		sendText(msg, cLang["genericResponse"])
		return
	}

	SendResponse(msg, res)
}

func getDoomedName(text string) (name string) {
	tokens := strings.Split(text, " ")
	if len(tokens) < 2 {
		return
	}
	name = strings.ToLower(tokens[1])
	name = resources.LeftTrimAtSign(name)

	return
}
