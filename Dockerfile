# syntax=docker/dockerfile:1

FROM golang:1.19.5-alpine3.17

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o /discord-party-bot ./cmd

CMD [ "/discord-party-bot" ]