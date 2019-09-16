package core

import (
	"github.com/tecnologer/HellOrHeavenBot/db"
	"github.com/tecnologer/HellOrHeavenBot/model"
)

//RegisterChat add to log the chat from the message arrives
func RegisterChat(chat *model.Chat) error {
	return db.InsertOrUpdateChat(chat)
}
