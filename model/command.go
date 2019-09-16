package model

import "time"

//DefaultTimeout is the default timeout between two commands
const DefaultTimeout time.Duration = 10 * time.Second

//Command is the base struct for bot commands
type Command struct {
	ID              byte
	Aliases         []string
	Description     string
	Params          []string
	Action          Action
	Timeout         time.Duration
	AcceptsResponse bool
}

//HasAlias returns true if the alias is in the list of aliases
func (cmd *Command) HasAlias(alias string) bool {
	for _, a := range cmd.Aliases {
		if a == alias {
			return true
		}
	}
	return false
}
