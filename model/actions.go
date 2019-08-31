package model

import bot "gopkg.in/tucnak/telebot.v2"

//ActionReturn is the default return of the actions for the commands
type ActionReturn struct {
	Error error
}

//Action is executed when the command is called
type Action func(*bot.Message) *ActionReturn
