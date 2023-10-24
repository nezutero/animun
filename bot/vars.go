package bot

import (
	"github.com/enescakir/emoji"
)

const (
	creatorChatID int64 = 5785150199
)

var (
	isBotRunning    bool
	infoEmoji       = emoji.Sprintf("%v", emoji.Information)
	infinityEmoji   = emoji.Sprintf("%v", emoji.Infinity)
	stopEmoji       = emoji.Sprintf("%v", emoji.RedCircle)
	okEmoji         = emoji.Sprintf("%v", emoji.GreenCircle)
	cactusEmoji     = emoji.Sprintf("%v", emoji.Cactus)
	idkEmoji        = emoji.Sprintf("%v", emoji.OpenHands)
	GreenHeartEmoji = emoji.Sprintf("%v", emoji.GreenHeart)
)

var (
	alreadyRunningMsg     = okEmoji + " animun is already running"
	startMsg              = okEmoji + " animun is started"
	stopMsg               = okEmoji + " animun is stopped"
	helpMsg               = infoEmoji + " animun hints\n\n+ /help - to get all commands\n+ /start - to start animun\n+ /stop - to stop animun\n+ /schedule - to see schedule\n+ /support - to tell about bugs you found"
	thxForBugReportMsg    = GreenHeartEmoji + " thanks for your bug report!"
	selectDayMsg          = infinityEmoji + " select day you're interested in"
	describeTheProblemMsg = cactusEmoji + " please describe the problem:"
	idkMsg                = idkEmoji + " i don't know what you mean\n/help - to get all commands"
	supportMsg            = " bug report from user %s:\n%s"
)
