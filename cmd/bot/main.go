package main

import (
	"github.com/Dimadetected/pocket-tg-bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1687025157:AAGbegSVrFWHwAOHp5eCSiqriC-4E18RkvQ")
	if err != nil {
		log.Fatal(err)
	}

	pocketClient, err := pocket.NewClient("96523-8840334b96dcabb25386a179")
	if err != nil {
		log.Fatal(err)
	}

	tgBot := telegram.NewBot(bot, pocketClient, "http://localhost")
	err = tgBot.Start()
	if err != nil {
		log.Fatal(err)
	}

}
