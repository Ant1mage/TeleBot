package main

import (
	"TeleBot/teleBot"
	"log"
)

const (
	MessageText = iota
	MessageImage
	MessageImageAndText
)

var quitChan = make(chan bool)

func main() {
	user, err := teleBot.NewBotApi()
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", user.FirstName)

	u := teleBot.NewUpdate(0)
	u.Timeout = 60

	updates, err := teleBot.ReceiveUpdateChan(u)

	for {
		select {
		case update := <-updates:
			if update.Message.Text == "胡婷" {
				msg := teleBot.NewMessage(update.Message.Chat.ID, MessageText)
				msg.Text = "你怎么知道我爱她!?"
				teleBot.Send(msg)
			} else {
				msg := teleBot.NewMessage(update.Message.Chat.ID, MessageText)
				msg.Text = update.Message.Text
				teleBot.Send(msg)
			}
		case <-quitChan:
			break
		}
	}

}
