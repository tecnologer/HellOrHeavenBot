package core

import (
	"github.com/tecnologer/HellOrHeavenBot/model"
	bot "github.com/yanzay/tbot"
)

//CommandList type for list of commands
type CommandList []*model.Command

//AcceptedCommands has the list of acepted commands
var AcceptedCommands = CommandList{
	&model.Command{
		ID:          1,
		Aliases:     []string{"hell", "infierno"},
		Description: "Agrega al usuario un boleto al infierno",
		Params:      []string{"<username>"},
		Timeout:     model.DefaultTimeout,
		Action:      Hell,
	},
	&model.Command{
		ID:          1,
		Aliases:     []string{"heaven", "cielo"},
		Description: "Agrega al usuario un boleto al cielo",
		Params:      []string{"<username>"},
		Timeout:     model.DefaultTimeout,
		Action:      Heaven,
	},
	&model.Command{
		ID:          1,
		Aliases:     []string{"stats", "estadisticas"},
		Description: "Retorna la cantidad de tickets al cielo y al infierno",
		Timeout:     model.DefaultTimeout,
		Action:      GetStats,
	},
	&model.Command{
		ID:          1,
		Aliases:     []string{"start"},
		Description: "Inicia el bot",
		Timeout:     model.DefaultTimeout,
		Action:      Start,
	},
}

//Call execute the function using the name of command
func (l CommandList) Call(cmd string, msg *bot.Message) {
	for _, c := range l {
		if c.HasAlias(cmd) {
			c.Action(msg)
		}
	}
}
