package bot

import (
	"fmt"
	"github.com/darenliang/jikan-go"
	"github.com/joho/godotenv"
	"github.com/kenjitheman/animun/api"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	isBotRunning := false

	log.Printf("[SUCCESS] authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	lastUserMessageTime := time.Now()

	for update := range updates {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			userInput := update.Message.Text

			if autoOff != nil {
				autoOff.Stop()
			}

			if time.Since(lastUserMessageTime) > 5*time.Minute {
				if isBotRunning {
					isBotRunning = false
					msg.Text = autoOffMsg
					msg.ReplyMarkup = StartKeyboard
					bot.Send(msg)
				}
			}

			lastUserMessageTime = time.Now()

			autoOff := time.NewTimer(5 * time.Minute)
			go func() {
				<-autoOff.C
				if isBotRunning {
					isBotRunning = false
					msg.Text = autoOffMsg
					msg.ReplyMarkup = StartKeyboard
					bot.Send(msg)
				}
			}()

			switch userInput {
			case "/start", "start":
				if !isBotRunning {
					isBotRunning = true
					msg.Text = startMsg
					msg.ReplyMarkup = GeneralKeyboard
				} else {
					msg.Text = alreadyRunningMsg
				}

			case "/help", "help":
				if isBotRunning {
					msg.Text = helpMsg
					msg.ReplyMarkup = GeneralKeyboard
				}

			case "schedule", "/schedule":
				if isBotRunning {
					msg.ReplyMarkup = WeekdaysKeyboard
					msg.Text = selectDayMsg
					bot.Send(msg)
					for {
						response := <-updates
						if response.Message == nil {
							continue
						}
						if response.Message.Chat.ID != update.Message.Chat.ID {
							continue
						}
						if response.Message.Text == "back to menu" {
							msg.Text = backToMenuMsg
							msg.ReplyMarkup = GeneralKeyboard
							break
						}
						selectedDay := response.Message.Text
						result, err := api.GetData(jikan.ScheduleFilter(selectedDay))
						if err != nil {
							msg.Text = fmt.Sprintf("[ERROR] %s", err)
							log.Println(err)
						} else {
							msg.Text = result
						}
						msg.ReplyMarkup = GeneralKeyboard
						break
					}
				}
			case "/stop", "stop":
				if isBotRunning {
					isBotRunning = false
					msg.Text = stopMsg
					msg.ReplyMarkup = StartKeyboard
				} else {
					msg.Text = alreadyStoppedMsg
				}

			case "/bug_report", "bug_report", "bug report":
				if isBotRunning {
					msg.ReplyMarkup = BackKeyboard
					msg.Text = describeTheProblemMsg
					bot.Send(msg)

					response := <-updates

					if response.Message != nil {
						if response.Message.Chat.ID != update.Message.Chat.ID {
							continue
						}
						description := response.Message.Text

						if description == "back to menu" {
							msg.Text = backToMenuMsg
							msg.ReplyMarkup = GeneralKeyboard
						} else {
							supportMsg = tgbotapi.NewMessage(
								creatorChatID,
								fmt.Sprintf(
									bugReportMsg,
									update.Message.From.UserName,
									description,
								),
							)
							msg.Text = thxForBugReportMsg
						}
						bot.Send(supportMsg)
						msg.ReplyMarkup = GeneralKeyboard
					}

				} else {
					msg.Text = alreadyStoppedMsg
				}
			default:
				if isBotRunning {
					msg.Text = idkMsg
				}
			}
			lastUserMessageTime = time.Now()
			if _, err := bot.Send(msg); err != nil {
				fmt.Println("[ERROR] error sending message")
			}
		}
	}

	if autoOff != nil {
		autoOff.Stop()
	}
}

