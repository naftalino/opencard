package main

import (
	"opencard/bot"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		color.Red("Error loading variables: ", err)
	}

	greenBackground := color.New(color.FgBlack, color.BgGreen).SprintFunc()

	println(greenBackground("[!] Opencard system — Created by " + os.Getenv("DEV_USERNAME")))

	go bot.StartBot()
}
