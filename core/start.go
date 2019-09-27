package core

import bot "github.com/yanzay/tbot"

//Start sends welcome to the user
func Start(msg *bot.Message) {
	sendText(msg, "hola wuap@ :kissing_heart::kissing_heart:")
}
