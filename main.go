package main

import (
	tgClient "TestGOBot/clients/telegram"
	"TestGOBot/consumer/eventconsumer"
	"TestGOBot/events/telegram"
	"TestGOBot/storage/files"
	"flag"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)
	log.Printf("starting telegram bot")

	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}

}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot")

	flag.Parse()
	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
