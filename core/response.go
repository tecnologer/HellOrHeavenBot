package core

import (
	log "github.com/sirupsen/logrus"
	bot "github.com/yanzay/tbot"
)

func NewResponse(msg *bot.Message) {
	resType, content := getResponseParams(msg.Text)

	if resType == -1 {
		buttons := getButtons()
		_, err := Client.SendMessage(msg.Chat.ID, "Selecciona un tipo", bot.OptInlineKeyboardMarkup(buttons))
		if err != nil {
			log.Println(err)
		}
	}

	if content == "" {

	}
}

func getResponseParams(text string) (resType int, contect string) {
	return -1, ""
}

func getButtons() *bot.InlineKeyboardMarkup {
	var buttons []bot.InlineKeyboardButton

	for _, c := range AcceptedCommands {
		if !c.AcceptsResponse {
			continue
		}

		button := bot.InlineKeyboardButton{
			Text:         c.Aliases[0],
			CallbackData: string(c.ID),
		}
		buttons = append(buttons, button)
	}

	return &bot.InlineKeyboardMarkup{
		InlineKeyboard: [][]bot.InlineKeyboardButton{
			buttons,
		},
	}
}
