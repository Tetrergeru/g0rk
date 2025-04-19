package main

import (
	"log"
	"math/rand"
	"os"
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var isThisTrue *regexp.Regexp = regexp.MustCompile(`(?i)gork\s+is\s+this\s+true`)

func main() {
	token := os.Getenv("GORK_TELEGRAM_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && isThisTrue.MatchString(update.Message.Text) {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, calculateTruthness())
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}

}

func calculateTruthness() string {
	if rand.Intn(2) == 0 {
		return "no this is false"
	}

	return "this is absolutely true"
}
