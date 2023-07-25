package tg

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var isBotRunning bool

func Start() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("[ERROR] error loading .env file")
		log.Panic(err)
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	isBotRunning = true

	generalKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("help"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("schedule"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("support"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("stop"),
		),
	)

	startKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("start"),
		),
	)

	log.Printf("[SUCCESS] authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "start":
      isBotRunning = true
			if isBotRunning {
				msg.Text = "animun is already running"
			}
      msg.Text = "animun has been started"
      msg.ReplyMarkup = generalKeyboard
      isBotRunning = true
		case "help":
			if isBotRunning {
				msg.Text = "animun hints\n\n /help - to get all commands\n /start - to start animun\n /stop - to stop animun\n /schedule - to see schedule\n /support - to tell about bugs you found"
				msg.ReplyMarkup = generalKeyboard
			}
		case "schedule":
			if isBotRunning {
        msg.Text = "schedule days of week keaboard must be here"
				// msg.ReplyMarkup = "schedule days of week keaboard must be here"
			}
		case "stop":
			if isBotRunning {
				msg.Text = "animun has been stopped"
				msg.ReplyMarkup = startKeyboard
				isBotRunning = false
			}
    case "support":
			if isBotRunning {
				msg.Text = "please, specify problems/bugs you found"
			}
		default:
			if isBotRunning {
				msg.Text = "i don't understand you\n/help"
			}
		}
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
			fmt.Printf("[ERROR] error sending message")
			log.Panic(err)
		}
	}
}
