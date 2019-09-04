package model

import "time"

const DefaultTimeout time.Duration = 10 * time.Second

//Command is the base struct for bot commands
type Command struct {
	ID          byte
	Aliases     []string
	Description string
	Params      []string
	Action      Action
	Timeout     time.Duration
}
