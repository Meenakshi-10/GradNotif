package main

import (
	"fmt"
	"gradnotif/notify"
	"gradnotif/scrape"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Load environment variables from .env file
func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	loadEnv()
	mentions := scrape.Scrape()
	fmt.Println("MENTIONS-----------------------------")
	fmt.Println(mentions)
	token := os.Getenv("PUSHOVER_TOKEN")
	user := os.Getenv("PUSHOVER_USER_KEY")

	for mention := range mentions {
		message := mentions[mention]
		err := notify.SendNotification(token, user, message)
		if err != nil {
			println("Error:", err.Error())
		}
	}
}
