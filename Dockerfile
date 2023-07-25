FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/main .

ENTRYPOINT ["./main"]

