package main

import "fmt"

type Bot interface {
	getReply() string
}

func getReply(bot Bot) {
	fmt.Println(bot.getReply())
}

type Writter interface {
	write()
}
type BotWritter interface {
	Bot
	Writter
}

func write(botWritter BotWritter) {
	botWritter.write()
}

type ChatGPT struct{}

func (ChatGPT) getReply() string {
	return "Hi I'am ChatGPT"
}
func (bot ChatGPT) write() {
	fmt.Println("Hi I am bot Writter and you are working with ChatGPT")
}

type Bard struct{}

func (Bard) getReply() string {
	return "Hi I'm Bard"
}
func main() {
	bot1 := ChatGPT{}
	write(bot1)

	bot2 := Bard{}
	getReply(bot2)

}
