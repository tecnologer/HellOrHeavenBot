package commands

import (
	"github.com/derekparker/delve/pkg/proc/core"
	"github.com/tecnologer/HellOrHeavenBot/model"
)

//CommandList has the list of the supported commands for the bot
var CommandList = map[string]*model.Command{
	"/hell": &model.Command{
		Name:        "hell",
		Description: "",
		Params:      []string{"[@]<username>"},
		Action:      core.Hell,
	},
}
