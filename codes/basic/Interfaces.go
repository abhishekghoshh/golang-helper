package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// https://gobyexample.com/interfaces
func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	bot1 := ChatGPT{}
	write(bot1)

	bot2 := Bard{}
	getReply(bot2)
}

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
