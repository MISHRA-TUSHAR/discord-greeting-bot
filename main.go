package main

import (
	"fmt"

	"github.com/MISHRA-TUSHAR/discord-greeting-bot/bot"
	"github.com/MISHRA-TUSHAR/discord-greeting-bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println("Error reading config file")
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
