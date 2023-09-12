package tg

import (
	"fmt"
	"log"
	"os"

	"github.com/darenliang/jikan-go"
	"github.com/enescakir/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	// "github.com/joho/godotenv"

	"main.go/api"
)

var (
	isBotRunning  bool
	creatorChatID int64
)

func Start() {
	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	fmt.Println("[ERROR] error loading .env file")
	// 	log.Panic(err)
	// }
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	isBotRunning = false

	generalKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/help"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/schedule"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/support"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/stop"),
		),
	)

	startKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/start"),
		),
	)

	weekdaysKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("monday"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("tuesday"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("wednesday"),
		),

		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("thursday"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("friday"),
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
			okEmoji := emoji.Sprintf("%v", emoji.GreenCircle)
			if isBotRunning {
				msg.Text = okEmoji + " animun is already running"
			}
			msg.Text = okEmoji + " animun has been started"
			msg.ReplyMarkup = generalKeyboard
		case "help":
			if isBotRunning {
				infoEmoji := emoji.Sprintf("%v", emoji.Information)
				msg.Text = infoEmoji + " animun hints\n\n+ /help - to get all commands\n+ /start - to start animun\n+ /stop - to stop animun\n+ /schedule - to see schedule\n+ /support - to tell about bugs you found"
				msg.ReplyMarkup = generalKeyboard
			}
		case "schedule":
			if isBotRunning {
				msg.ReplyMarkup = weekdaysKeyboard
				infinityEmoji := emoji.Sprintf("%v", emoji.Infinity)
				msg.Text = infinityEmoji + " select day you're interested in"
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
		case "stop":
			if isBotRunning {
				stopEmoji := emoji.Sprintf("%v", emoji.RedCircle)
				msg.Text = stopEmoji + " animun has been stopped"
				msg.ReplyMarkup = startKeyboard
				isBotRunning = false
			}
			isBotRunning = false
		case "support":
			if isBotRunning {
				cactusEmoji := emoji.Sprintf("%v", emoji.Cactus)
				creatorChatID = 5785150199
				msg.Text = cactusEmoji + " please describe the problem:"
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
					GreenHeartEmoji := emoji.Sprintf("%v", emoji.GreenHeart)
					msg.Text = GreenHeartEmoji + " thanks for your bug report!"
					bot.Send(msg)
					supportMsg := tgbotapi.NewMessage(
						creatorChatID,
						fmt.Sprintf(
							" bug report from user %s:\n%s",
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
				idkEmoji := emoji.Sprintf("%v", emoji.OpenHands)
				msg.Text = idkEmoji + " i don't understand you\n/help"
			}
		}
		if _, err := bot.Send(msg); err != nil {
			fmt.Printf("[ERROR] error sending message")
		}
	}
}
