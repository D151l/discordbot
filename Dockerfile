FROM golang:1.18-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go build -o discordbot

CMD ["./discordbot"]
