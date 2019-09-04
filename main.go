package main

import (
	"bufio"
	"log"
	"os"

	"github.com/tecnologer/HellOrHeavenBot/core"
)

func main() {
	err := core.StartBot()

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
