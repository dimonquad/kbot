package main

import (
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

var (
	teleToken = os.Getenv("Tele_Token")
)

func main() {
	pref := tele.Settings{
		Token:  teleToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatalf("Check systemEnv: %v", err)
		return
	}

	b.Handle(tele.OnText, func(c tele.Context) error {
		log.Print(c.Message().Payload, c.Text())
		payload := c.Message().Payload
		switch payload {
		case "hello":
			return c.Send("Hello! How can I help you?")
		}

		return err
	})

	b.Start()
}
