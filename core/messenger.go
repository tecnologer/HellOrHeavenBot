package core

import (
	"github.com/kyokomi/emoji"
	m "github.com/tecnologer/HellOrHeavenBot/model"
	bot "github.com/yanzay/tbot"
)

//SendResponse is in charge of send the data of the response (text, sticker, gif, etc)
func SendResponse(msg *bot.Message, res *m.Response) {
	switch res.Type {
	case m.Text:
		sendText(msg, res.Content)
	case m.Sticker:
		Client.SendSticker(msg.Chat.ID, res.Content)
	}
}

func sendText(msg *bot.Message, text string) {
	text = emoji.Sprint(text)
	Client.SendMessage(msg.Chat.ID, text)
}
