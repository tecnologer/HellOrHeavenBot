package main

import (
	"bufio"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/HellOrHeavenBot/core"
	"github.com/tecnologer/HellOrHeavenBot/db"
	"github.com/tecnologer/HellOrHeavenBot/test"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	file, err := os.OpenFile("HellOrHeavenBot.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err == nil {
		// Output to stdout instead of the default stderr
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	// Only log the warning severity or above.
	log.SetLevel(log.TraceLevel)
}

func main() {
	log.Println("************************** new instance running **************************")
	err := db.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	test.RestoreData()

	err = core.StartBot()

	if err != nil {
		log.Fatal(err)
	}

	// waitExit()
}

func waitExit() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	in := input.Text()

	for in != "c" {
		in = input.Text()
	}
}