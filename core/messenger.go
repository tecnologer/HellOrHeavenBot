package core

import (
	m "github.com/tecnologer/HellOrHeavenBot/model"
	bot "github.com/yanzay/tbot"
)

//SendResponse is in charge of send the data of the response (text, sticker, gif, etc)
func SendResponse(msg *bot.Message, res *m.Response) {
	switch res.Type {
	case m.Text:
		Client.SendMessage(msg.Chat.ID, res.Content)
	case m.Sticker:
		Client.SendSticker(msg.Chat.ID, res.Content)
	}
}

func sendText(msg *bot.Message, text string) {
	Client.SendMessage(msg.Chat.ID, text)
}
