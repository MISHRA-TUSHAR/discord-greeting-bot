package main

import (
	"fmt"
	"io/ioutil"
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
