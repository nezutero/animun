package bot

import (
	"github.com/enescakir/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

const (
	creatorChatID int64 = 5785150199
)

var (
	isBotRunning        bool
	lastUserMessageTime time.Time
	autoOff             *time.Timer
	supportMsg          tgbotapi.MessageConfig
)

var (
	infoEmoji       = emoji.Sprintf("%v", emoji.Information)
	infinityEmoji   = emoji.Sprintf("%v", emoji.Infinity)
	stopEmoji       = emoji.Sprintf("%v", emoji.RedCircle)
	okEmoji         = emoji.Sprintf("%v", emoji.GreenCircle)
	cactusEmoji     = emoji.Sprintf("%v", emoji.Cactus)
	idkEmoji        = emoji.Sprintf("%v", emoji.OpenHands)
	GreenHeartEmoji = emoji.Sprintf("%v", emoji.GreenHeart)
)

var (
	alreadyRunningMsg     = okEmoji + " animun is already running\n[ ? ] /stop - to stop animun"
	startMsg              = okEmoji + " animun is started\n[ ? ] /help - to get all commands"
	stopMsg               = stopEmoji + " animun is stopped\n[ ? ] /start - to start animun"
	helpMsg               = infoEmoji + " animun hints\n\n+ /help - to get all commands\n+ /start - to start animun\n+ /stop - to stop animun\n+ /schedule - to see schedule\n+ /bug_report - to tell about bugs you found"
	thxForBugReportMsg    = GreenHeartEmoji + " thanks for your bug report!"
	selectDayMsg          = infinityEmoji + " select day you're interested in"
	describeTheProblemMsg = cactusEmoji + " please describe the problem:"
	idkMsg                = idkEmoji + " i don't know what you mean\n[ ? ] /help - to get all commands"
	bugReportMsg          = "[ ! ] bug report from user @%s\n[ ! ] msg: %s"
	alreadyStoppedMsg     = stopEmoji + " animun is already stopped\n[ ? ] /start - to start animun"
	backToMenuMsg         = okEmoji + " back to menu"
	autoOffMsg            = stopEmoji + " animun is stopped due to inactivity\n[ ? ] /start - to start animun"
)
