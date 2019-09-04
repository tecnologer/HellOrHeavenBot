package model

import bot "github.com/yanzay/tbot"

//ActionReturn is the default return of the actions for the commands
type ActionReturn struct {
	Response ActionResponse
	Error    error
}

//ActionResponse is the struct for tha data returned by actions
type ActionResponse struct {
	Answer string
	Type   ResponseType
}

//Action is executed when the command is called
type Action func(*bot.Message)
