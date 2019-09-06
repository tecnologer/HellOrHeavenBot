package core

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/HellOrHeavenBot/db"
	m "github.com/tecnologer/HellOrHeavenBot/model"
	bot "github.com/yanzay/tbot"
)

//Hell is the function in charge of registering the hell tickets
func Hell(msg *bot.Message) {
	doomedName := getDoomedName(msg.Text)

	if doomedName == "" {
		sendText(msg, cLang["ticketsNameRequired"])
		return
	}
	err := db.InsertStat(doomedName, m.StatsHell)

	if err != nil {
		log.Println(err)
		sendText(msg, cLang["genericFail"])
		return
	}
	SendResponse(msg, &m.Response{Content: "CAADAwADcQADJaHuBOXuxozHxyQrAg", Type: m.Sticker})
}

//Heaven is the function in charge of registering the hell tickets
func Heaven(msg *bot.Message) {
	doomedName := getDoomedName(msg.Text)

	if doomedName == "" {
		sendText(msg, cLang["ticketsNameRequired"])
		return
	}
	err := db.InsertStat(doomedName, m.StatsHeaven)

	if err != nil {
		log.Println(err)
		sendText(msg, cLang["genericFail"])
		return
	}

	SendResponse(msg, &m.Response{Content: "CAADAwADcQADJaHuBOXuxozHxyQrAg", Type: m.Sticker})
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

func getDoomedName(text string) string {
	tokens := strings.Split(text, " ")
	if len(tokens) < 2 {
		return ""
	}

	if strings.HasPrefix(tokens[1], "@") {
		return strings.ToLower(tokens[1][1:])
	}

	return strings.ToLower(tokens[1])
}
