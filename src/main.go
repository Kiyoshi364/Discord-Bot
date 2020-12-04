package main

import (
	"fmt"
	db "./discordbot"
)

func main() {
	err := db.Init()
	if err != nil {
		fmt.Println("Main: ", err.Error())
		return
	}

	fmt.Println("Bot is running")

	<-make(chan struct{})
	return
}
