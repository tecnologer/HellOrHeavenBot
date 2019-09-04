package core

import "github.com/tecnologer/HellOrHeavenBot/model"

//CommandList has the list of acepted commands
var CommandList = map[string]*model.Command{
	"hell": &model.Command{
		ID:          1,
		Aliases:     []string{"hell", "infierno"},
		Description: "Agrega al usuario un boleto al infierno",
		Params:      []string{"<username>"},
		Timeout:     model.DefaultTimeout,
		Action:      Hell,
	},
	"stats": &model.Command{
		ID:          1,
		Aliases:     []string{"stats", "estadisticas"},
		Description: "Retorna la cantidad de tickets al cielo y al infierno",
		Timeout:     model.DefaultTimeout,
		Action:      GetStats,
	},
}
