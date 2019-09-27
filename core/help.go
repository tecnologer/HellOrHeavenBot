package core

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/HellOrHeavenBot/model"
	bot "github.com/yanzay/tbot"
)

//SendHelp creates a string with all commands and send them in a message
func SendHelp(msg *bot.Message) {
	const helpRow = "- /%s: `%s`\n"

	help := ""

	for _, c := range AcceptedCommands {
		help += fmt.Sprintf(helpRow, c.Aliases[0], c.Description)
	}

	sendText(msg, help)
}

//GetAlias sends a list of buttons for commands who has aliases
func GetAlias(msg *bot.Message) {
	buttons := getAliasButtons()
	_, err := Client.SendMessage(msg.Chat.ID, "Alias de que comando?", bot.OptInlineKeyboardMarkup(buttons))
	if err != nil {
		log.Println(err)
	}
}

//GetAliasOfCmd sends in a message the aliases for the selected command
func GetAliasOfCmd(msg *bot.Message, cmd *model.Command) {
	list := fmt.Sprintf("Alias del comando /%s:\n\n", cmd.Aliases[0])
	for _, alias := range cmd.Aliases {
		list += fmt.Sprintf("/%s\n", alias)
	}

	sendText(msg, list)
}

//getAliasButtons creates the list of buttons for the commands who has aliases
func getAliasButtons() *bot.InlineKeyboardMarkup {
	var buttons []bot.InlineKeyboardButton

	for _, c := range AcceptedCommands {
		if len(c.Aliases) <= 1 {
			continue
		}

		button := bot.InlineKeyboardButton{
			Text:         c.Aliases[0],
			CallbackData: fmt.Sprintf("alias: %d", c.ID),
		}
		buttons = append(buttons, button)
	}

	return &bot.InlineKeyboardMarkup{
		InlineKeyboard: [][]bot.InlineKeyboardButton{
			buttons,
		},
	}
}
