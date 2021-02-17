package main

import (
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"net/http"
)

const (
	BotToken   = "1659894835:AAFMGTrQgxWRc8Qh8rlGRiziuZ1mxMBs7iA"
	WebhookURL = "http://3.134.80.221/"
)



func main() {
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		panic(err)
	}

	// bot.Debug = true
	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(WebhookURL))
	if err != nil {
		panic(err)
	}

	updates := bot.ListenForWebhook("/")

	go http.ListenAndServe(":8080", nil)
	fmt.Println("start listen :8080")

	// получаем все обновления из канала updates
	for update := range updates {
		if update.Message.Text == "Hello" {
			bot.Send(tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"poshel nahui",
			))
		} else if update.Message.Text == "privet" && update.Message.From.UserName == "mukhametkaly" {
			bot.Send(tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"krasava",
			))
		}

	}
}