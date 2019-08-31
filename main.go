package main

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/HellOrHeavenBot/resources"
	bot "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := bot.NewBot(bot.Settings{
		Token:  resources.BotToken,
		Poller: &bot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *bot.Message) {
		b.Send(m.Sender, "hello world")
	})

	log.Println("Listening...")
	b.Start()
}
