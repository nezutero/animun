package bot

import (
	"fmt"
	"log"
	"os"

	"github.com/darenliang/jikan-go"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	"github.com/kenjitheman/animun/api"
)

func Start() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("[ERROR] error loading .env file")
		log.Panic(err)
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	isBotRunning = false

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
		case "start", "/start":
			isBotRunning = true
			if isBotRunning {
				msg.Text = alreadyRunningMsg
			}
			msg.Text = startMsg
			msg.ReplyMarkup = generalKeyboard
		case "help", "/help":
			if isBotRunning {
				msg.Text = helpMsg
				msg.ReplyMarkup = generalKeyboard
			}
		case "schedule", "/schedule":
			if isBotRunning {
				msg.ReplyMarkup = weekdaysKeyboard
				msg.Text = selectDayMsg
				for {
					response := <-updates
					if response.Message == nil {
						continue
					}
					if response.Message.Chat.ID != update.Message.Chat.ID {
						continue
					}
					selectedDay := response.Message.Text
					result, err := api.GetData(jikan.ScheduleFilter(selectedDay))
					if err != nil {
						msg.Text = fmt.Sprintf("[ERROR] %s", err)
					} else {
						msg.Text = result
					}
					msg.ReplyMarkup = generalKeyboard
					break
				}
			}
		case "stop", "/stop":
			if isBotRunning {
				msg.Text = stopMsg
				msg.ReplyMarkup = startKeyboard
				isBotRunning = false
			}
			isBotRunning = false
		case "bug_report", "/bug_report": //TODO: back button
			if isBotRunning {
				msg.Text = describeTheProblemMsg
				bot.Send(msg)
				for {
					response := <-updates
					if response.Message == nil {
						continue
					}
					if response.Message.Chat.ID != update.Message.Chat.ID {
						continue
					}
					description := response.Message.Text
					msg.Text = thxForBugReportMsg
					bot.Send(msg)
					supportMsg := tgbotapi.NewMessage(
						creatorChatID,
						fmt.Sprintf(
							supportMsg,
							update.Message.From.UserName,
							description,
						),
					)
					bot.Send(supportMsg)
					break
				}
				continue
			}
			isBotRunning = false

		default:
			if isBotRunning {
				msg.Text = idkMsg
			}
		}
		if _, err := bot.Send(msg); err != nil {
			fmt.Printf("[ERROR] error sending message")
		}
	}
}
