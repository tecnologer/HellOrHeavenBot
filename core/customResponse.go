package core

import (
	"encoding/json"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/HellOrHeavenBot/db"
	"github.com/tecnologer/HellOrHeavenBot/model"
	"github.com/tecnologer/HellOrHeavenBot/resources"
	bot "github.com/yanzay/tbot"
)

var allahID = 10244644
var allahChats = []json.Number{"10244644"}

var incompleteCResponse map[uint32]*model.CustomResponse

func init() {
	incompleteCResponse = make(map[uint32]*model.CustomResponse)
}

//hasCustomResponse checks if the text match with some regex
func hasCustomResponse(msg *bot.Message) {
	if msg.Text == "" {
		return
	}

	readyResponse := make([]*model.Response, 0)
	defer func() {
		if r := recover(); r != nil {
			log.WithFields(log.Fields{
				"Responses": readyResponse,
				"text":      msg.Text,
				"chatID":    msg.Chat.ID,
			}).Debug(r)
		}
	}()

	responses, err := db.RetrieveCustomResponses(msg.Chat.ID)

	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"text":   msg.Text,
			"chatID": msg.Chat.ID,
		}).Error("error when try retrive custom responses")
		return
	}

	if len(responses) == 0 {
		return
	}

	for _, res := range responses {
		regex, err := regexp.Compile(res.Regex)
		if err != nil {
			log.WithError(err).WithFields(log.Fields{
				"text":   msg.Text,
				"chatID": msg.Chat.ID,
				"regex":  res.Regex,
			}).Error("error when try compile the regex")
			continue
		}

		if regex.Match([]byte(msg.Text)) {
			readyResponse = append(readyResponse, &model.Response{Content: res.Response, Type: res.ResponseType})
		}
	}

	if len(readyResponse) == 0 {
		return
	}

	if len(readyResponse) == 1 {
		SendResponse(msg, readyResponse[0])
		return
	}

	selection := resources.GetRandomIntFromRange(0, len(readyResponse)-1)

	SendResponse(msg, readyResponse[selection])
}

//NewCustomResponse starts process to insert new custom response
func NewCustomResponse(msg *bot.Message) {
	chunks := strings.SplitAfterN(msg.Text, " ", 2)

	if len(chunks) < 2 {
		sendText(msg, cLang["cResponseRegexReq"])
		return
	}

	regex := chunks[1]

	newResponse := &model.CustomResponse{
		ResponseType: model.Text,
		Response:     "",
		Regex:        regex,
		ChatID:       getChatID(msg.Chat),
		Author:       msg.From.ID,
	}

	identifier := getResponseIdentifier(msg.From)
	incompleteCResponse[identifier] = newResponse

	sendText(msg, cLang["requestResponseContent"])
}

func getChatID(chat bot.Chat) json.Number {
	chatID := json.Number(chat.ID)
	for _, chID := range allahChats {
		if chID == chatID {
			return ""
		}
	}

	return chatID
}

//AddCustomResponse validate the regex and insert into db
func AddCustomResponse(cRes *model.CustomResponse) error {
	_, err := regexp.Compile(cRes.Regex)

	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"CustomResponse": cRes,
		}).Error("error when try compile the regex")
		return err
	}

	return db.InsertCustomResponse(cRes)
}

//HasUserIncompleteCustomResponse returns true if the user of the message has incomplete custom responses
func HasUserIncompleteCustomResponse(from *bot.User) bool {

	identifier := getResponseIdentifier(from)
	_, exists := incompleteCResponse[identifier]

	return exists
}

func setContentToIncompleteCustomResponse(from *bot.User, msg *bot.Message) {
	if !HasUserIncompleteCustomResponse(from) {
		return
	}

	identifier := getResponseIdentifier(from)
	res, _ := incompleteCResponse[identifier]

	if msg.Sticker != nil {
		res.ResponseType = model.Sticker
		res.Response = msg.Sticker.FileID
	} else if msg.Document != nil {
		res.ResponseType = model.Gif
		res.Response = msg.Document.FileID
	} else if strings.HasPrefix(msg.Text, "/") || msg.Text == "" {
		return
	} else {
		res.ResponseType = model.Text
		res.Response = msg.Text
	}

	addCustomResponse(msg, res)
	delete(incompleteCResponse, identifier)
}

func addCustomResponse(msg *bot.Message, newResponse *model.CustomResponse) {

	err := AddCustomResponse(newResponse)
	if err != nil {
		sendText(msg, cLang["customResponseStoredFailed"])
		log.WithField("Custom Response", newResponse).
			WithError(err).Debug("error when try store a incomplete custom response")
		return
	}

	sendText(msg, cLang["customResponseStored"])
}
