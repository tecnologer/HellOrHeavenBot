package core

import (
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/HellOrHeavenBot/db"
	"github.com/tecnologer/HellOrHeavenBot/model"
	"github.com/tecnologer/HellOrHeavenBot/resources"
	bot "github.com/yanzay/tbot"
)

var incompleteResponse map[uint32]*model.Response

func init() {
	incompleteResponse = make(map[uint32]*model.Response)
}

//NewResponse starts process to insert new response
func NewResponse(msg *bot.Message) {
	cmdID, content := getResponseParams(msg.Text)

	incomplete := false
	newResponse := &model.Response{
		Type:      model.Text,
		CommandID: cmdID,
		Content:   content,
	}

	defer func() {
		if incomplete {
			identifier := getResponseIdentifier(msg.From)
			incompleteResponse[identifier] = newResponse
		} else {
			addResponse(msg, newResponse)
		}
	}()

	if cmdID == 0 {
		incomplete = true
		go sendResponsesCommandsButtons(msg)
		return
	}

	if content == "" {
		incomplete = true
		sendText(msg, cLang["requestResponseContent"])
	}
}

//HasUserIncompleteRes returns true if the user of the messages has incomplete responses
func HasUserIncompleteRes(from *bot.User) bool {

	identifier := getResponseIdentifier(from)
	_, exists := incompleteResponse[identifier]

	return exists
}

func getResponseParams(text string) (byte, string) {
	tokens := strings.SplitAfterN(text, " ", 3)

	//only command
	if len(tokens) == 1 {
		return 0, ""
	}

	cmdIDString := ""
	if len(tokens) > 1 {
		cmdIDString = strings.Trim(tokens[1], " ")
	}

	//command and command id for response
	if len(tokens) == 2 {
		cmdID, err := strconv.Atoi(cmdIDString)
		if err != nil {
			log.WithError(err).Error("error when try parse command id for new response")
			return 0, ""
		}
		if cmdID > 255 {
			return 0, ""
		}

		return byte(cmdID), ""
	}

	cmdID, err := strconv.Atoi(cmdIDString)
	content := tokens[2]
	if err != nil {
		log.WithError(err).Error("error when try parse command id for new response")
		return 0, content
	}

	if cmdID > 255 {
		return 0, content
	}

	return byte(cmdID), content
}

func getResponseButtons() *bot.InlineKeyboardMarkup {
	var buttons []bot.InlineKeyboardButton

	for _, c := range AcceptedCommands {
		if !c.AcceptsResponse {
			continue
		}

		button := bot.InlineKeyboardButton{
			Text:         c.Aliases[0],
			CallbackData: fmt.Sprintf("type: %d", c.ID),
		}
		buttons = append(buttons, button)
	}

	return &bot.InlineKeyboardMarkup{
		InlineKeyboard: [][]bot.InlineKeyboardButton{
			buttons,
		},
	}
}

func sendResponsesCommandsButtons(msg *bot.Message) {
	buttons := getResponseButtons()
	_, err := Client.SendMessage(msg.Chat.ID, "Selecciona un tipo", bot.OptInlineKeyboardMarkup(buttons))
	if err != nil {
		log.Println(err)
	}
}

func setContentToIncomplete(from *bot.User, msg *bot.Message) {
	if !HasUserIncompleteRes(from) {
		return
	}

	identifier := getResponseIdentifier(from)
	res, _ := incompleteResponse[identifier]

	if msg.Sticker != nil {
		res.Type = model.Sticker
		res.Content = msg.Sticker.FileID
	} else if msg.Document != nil {
		res.Type = model.Gif
		res.Content = msg.Document.FileID
	} else {
		res.Type = model.Text
		res.Content = msg.Text
	}

	isResponseComplete(from, msg)
}

func setCmdIDToIncomplete(from *bot.User, msg *bot.Message, cmdID byte) {
	if !HasUserIncompleteRes(from) {
		return
	}

	identifier := getResponseIdentifier(from)
	incompleteResponse[identifier].CommandID = cmdID

	isResponseComplete(from, msg)
}

func isResponseComplete(from *bot.User, msg *bot.Message) {
	identifier := getResponseIdentifier(from)
	res, _ := incompleteResponse[identifier]

	if res.CommandID == 0 {
		sendResponsesCommandsButtons(msg)
		return
	}

	if res.Content == "" {
		sendText(msg, cLang["requestResponseContent"])
		return
	}

	err := addResponse(msg, res)
	if err != nil {
		log.WithError(err).Info("error when validate if response is completed")
	}

	delete(incompleteResponse, identifier)
}

func addResponse(msg *bot.Message, newResponse *model.Response) error {
	if newResponse.Language == "" {
		newResponse.Language = msg.From.LanguageCode
	}

	err := db.InsertResponse(newResponse)
	if err != nil {
		return err
	}

	sendText(msg, cLang["responseStored"])
	return nil
}

func getResponseIdentifier(from *bot.User) uint32 {
	name := resources.GetName(from)
	ident := resources.GetHash(from.ID, name)
	log.WithField("identifier", ident).Info("identifier calculated")
	return ident
}
