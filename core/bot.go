package core

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/HellOrHeavenBot/resources"
	bot "github.com/yanzay/tbot"
)

//Bot is the instance of the bot
var Bot *bot.Server

//Client of the bot
var Client *bot.Client

//StartBot runs the bot
func StartBot() error {
	var err error
	Bot = bot.New(resources.BotToken)
	Client = Bot.Client()

	if err != nil {
		return err
	}

	// registerHandlers()
	Bot.HandleMessage(".*", messagesHandle)

	log.Println("Listening...")
	Bot.Start()
	return nil
}

func messagesHandle(msg *bot.Message) {
	cmd := getCommand(msg.Text)

	if cmd != "" {
		CommandList[cmd].Action(msg)
	}
}

func getCommand(text string) string {
	if strings.HasPrefix(text, "/") {
		tokens := strings.Split(text, " ")
		cmd := tokens[0]
		return strings.ToLower(cmd[1:])
	}

	return ""
}

func registerHandlers() {
	for _, cmd := range CommandList {
		for _, alias := range cmd.Aliases {
			Bot.HandleMessage("/"+alias, cmd.Action)
		}
	}
}
