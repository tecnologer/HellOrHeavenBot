package core

import (
	"fmt"
	"time"

	bot "github.com/yanzay/tbot"
)

//Uptime returns time running
func Uptime(msg *bot.Message) {
	uptime := time.Since(StartupTime)

	sendText(msg, fmt.Sprintf("Time running: %v", uptime))
}
