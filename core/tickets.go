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
		sendText(msg, "El nombre del condenado es requerido")
		return
	}
	err := db.InsertStat(doomedName, m.StatsHell)

	if err != nil {
		log.Println(err)
		sendText(msg, "falio ferga, no funco")
		return
	}
	SendResponse(msg, &m.Response{Content: "CAADAwADcQADJaHuBOXuxozHxyQrAg", Type: m.Sticker})
}

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
	// if msg.From.Username != "" {
	// 	return msg.From.Username
	// }
	// return strings.Join([]string{msg.From.FirstName, msg.From.LastName}, "")
}
