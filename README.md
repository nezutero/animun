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

## project structure

```go
.
├── api
│   ├── api.go
│   └── api_test.go
├── bot
│   ├── bot.go
│   ├── keyboards.go
│   └── vars.go
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE
├── main.go
└── README.md
```

## installation

```shell
git clone https://github.com/kenjitheman/animun
```

## usage

- create .env file and inside you should create env variable with your api key

```.env
TELEGRAM_API_TOKEN=YOUR_TOKEN
```

- then you should uncomment these lines in bot.go 
    - **( ! you need uncomment commented lines only if you using this way !)**

```go
// "github.com/joho/godotenv"
```

```go
// err := godotenv.Load("../.env")
// if err != nil {
// 	fmt.Println("[ERROR] error loading .env file")
// 	log.Panic(err)
// }
```

#### run it using docker:

- you need to paste your api key in dockerfile:

```dockerfile
ENV TELEGRAM_API_TOKEN=YOUR_API_TOKEN
```

- run it:

```sh
docker build -t your_image_name .
docker run -d -p 8080:80 your_image_name
```

## contributing

- pull requests are welcome, for major changes, please open an issue first
to discuss what you would like to change

- please make sure to update tests as appropriate

## license

- [MIT](https://choosealicense.com/licenses/mit/)
