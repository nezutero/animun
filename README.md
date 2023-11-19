<h3 align="center">Tg bot for self use, was created coz I want to track anime release schedule</h3>

###

<img align="right" height="200" src="https://media.tenor.com/6VJldkd3beMAAAAC/kaori-miyazono-kousei-arima.gif"  />

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
  <img width="15" />
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" height="200" alt="docker logo"  />
</div>

###

## Project structure

```go
animun
│
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

## Installation

```sh
git clone https://github.com/kenjitheman/animun
```

## Usage

- Create .env file and inside you should create env variable with your api key:

```.env
TELEGRAM_API_TOKEN=YOUR_TOKEN
```

- You need to uncomment these lines in bot.go if you are going to run it using go run or go build:

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

- To run it:

```sh
go run main.go
```

- Or build and run:

```sh
go build
```

```sh
./animun
```

#### Run it using Docker:

- You need to paste your api key in dockerfile:

```dockerfile
ENV TELEGRAM_API_TOKEN=YOUR_API_TOKEN
```

- Run it:

```sh
docker build -t your_image_name .
docker run -d -p 8080:80 your_image_name
```

## Contributing

- Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

- Please make sure to update tests as appropriate.

## License

- [MIT](./LICENSE)
