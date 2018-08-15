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
			results, err := teleBot.MakeTuringResult(update.Message.Text, MessageText)
			if err != nil {
				log.Panic(err)
			}
			for _, result := range results {
				if result.ResultType == "image" {
					msg := teleBot.NewMessage(update.Message.Chat.ID, MessageImage)
					msg.Photo = result.Values.Image
					teleBot.Send(msg)
				} else if result.ResultType == "text" {
					msg := teleBot.NewMessage(update.Message.Chat.ID, MessageText)
					msg.Text = result.Values.Text
					teleBot.Send(msg)
				}
			}
		case <-quitChan:
			break
		}
	}

}
