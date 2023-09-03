<h3 align="center">tg bot for self use, was created coz I want to track anime release schedule</h3>

###

<img align="right" height="200" src="https://media.tenor.com/6VJldkd3beMAAAAC/kaori-miyazono-kousei-arima.gif"  />

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
  <img width="15" />
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" height="200" alt="docker logo"  />
</div>

###

## installation

```
git clone https://github.com/kenjitheman/animun
```

## usage

- create .env file and inside you should create env variable with your api key ->

```
TELEGRAM_API_TOKEN=YOUR_TOKEN
```

- then you should uncomment commented lines in tg/tg.go ( ! you need uncomment commented lines only if you using this way !) ->

```
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
```

#### you can also run it using docker:

- you need to paste your api key in dockerfile:

```
ENV TELEGRAM_API_TOKEN=YOUR_API_TOKEN
```

- run it:

```
docker build -t your_image_name .
docker run -d -p 8080:80 your_image_name
```

## contributing

- pull requests are welcome, for major changes, please open an issue first
to discuss what you would like to change

## license

- [MIT](https://choosealicense.com/licenses/mit/)
