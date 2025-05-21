package main

import (
	"log"
	"math/rand"
	"os"
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var isThisTrue *regexp.Regexp = regexp.MustCompile(`(?i)gork\s+is\s+this\s+true`)
var isSummarize *regexp.Regexp = regexp.MustCompile(`(?i)(перескажи|summarize)`)

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
		if update.Message == nil {
			continue
		}

		responce := ""
		if isThisTrue.MatchString(update.Message.Text) {
			log.Printf("isThisTrue: [%s] %s\n", update.Message.From.UserName, update.Message.Text)

			if update.Message.ReplyToMessage != nil && update.Message.ReplyToMessage.From.UserName == "g0rk_bot" {
				responce = "this is absolutely true"
			} else {
				responce = calculateTruthness()
			}
		} else if isSummarize.MatchString(update.Message.Text) {
			log.Printf("isSummarize: [%s] %s\n", update.Message.From.UserName, update.Message.Text)

			responce = generateSummary()
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, responce)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}

}

func calculateTruthness() string {
	if rand.Intn(2) == 0 {
		return "no this is false"
	}

	return "this is absolutely true"
}

func generateSummary() string {
	switch rand.Intn(5) {
	case 0:
		return "Ну там это... Ведро.... Колодец.... И вода...."
	case 1:
		return "Все проблемы бекоз оф вок"
	case 2:
		return "Экономике России осталось 3 дня"
	case 3:
		return "В этом, как в зеркале отразилась вся суть России за последние 20 лет"
	default:
		return "Это критика капитализма"
	}
}
