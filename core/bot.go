package core

import (
	"strconv"
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
	Bot.HandleCallback(callBackHandler)

	log.Info("Listening...")
	StartupTime = time.Now()
	Bot.Start()
	return nil
}

func messagesHandle(msg *bot.Message) {
	// log.WithFields(log.Fields{
	// 	"msg": msg,
	// }).Info("new message")

	if msg.From.IsBot || msg.EditDate != 0 {
		return
	}

	cLang = lang.GetMessagesByLanguage(msg.From.LanguageCode)

	cmd := getCommand(msg.Text)

	if cmd != "" {
		AcceptedCommands.Call(cmd, msg)
		return
	}

	if HasUserIncompleteRes(msg.From) {
		setContentToIncomplete(msg.From, msg)
		return
	}

	if HasUserIncompleteCustomResponse(msg.From) {
		setContentToIncompleteCustomResponse(msg.From, msg)
		return
	}

	hasCustomResponse(msg)

}

func callBackHandler(cq *bot.CallbackQuery) {
	// log.WithFields(log.Fields{
	// 	"query": cq,
	// }).Info("new call back query")

	if strings.HasPrefix(cq.Data, "type:") && HasUserIncompleteRes(cq.From) {
		cmdIDString := strings.ReplaceAll(cq.Data, "type: ", "")
		cmdID, err := strconv.Atoi(cmdIDString)
		if err != nil {
			log.WithError(err).Error("callback query command id invalid")
			return
		}
		setCmdIDToIncomplete(cq.From, cq.Message, byte(cmdID))
	} else if strings.HasPrefix(cq.Data, "alias:") {
		cmdIDString := strings.ReplaceAll(cq.Data, "alias: ", "")
		cmdID, err := strconv.Atoi(cmdIDString)
		if err != nil {
			log.WithError(err).Error("callback query command id invalid")
			return
		}

		cmd, err := AcceptedCommands.GetCmdByID(cmdID)
		if err != nil {
			log.WithError(err)
			return
		}

		GetAliasOfCmd(cq.Message, cmd)
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
