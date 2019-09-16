package core

import (
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/HellOrHeavenBot/lang"
	"github.com/tecnologer/HellOrHeavenBot/resources"
	bot "github.com/yanzay/tbot"
)

//StartupTime time when the bot is started
var StartupTime time.Time

//Bot is the instance of the bot
var Bot *bot.Server

//Client of the bot
var Client *bot.Client

//cLang is the group of messages for the language of the current message
var cLang map[string]string

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
	StartupTime = time.Now()
	Bot.Start()
	return nil
}

func messagesHandle(msg *bot.Message) {
	if msg.From.IsBot || msg.EditDate != 0 {
		return
	}

	cLang = lang.GetMessagesByLanguage(msg.From.LanguageCode)

	cmd := getCommand(msg.Text)

	if cmd != "" {
		AcceptedCommands.Call(cmd, msg)
		return
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
