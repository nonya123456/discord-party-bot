package main

import (
	"fmt"

	"github.com/nonya123456/discord-party-bot/internal/config"
)

func main() {
	conf := config.New()
	fmt.Println(conf.Token)
}
