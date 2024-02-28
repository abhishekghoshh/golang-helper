package others

import (
	"fmt"
	"net/http"
	"time"
)

func ChannelLinks() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// type1(links)

	// type2(links)

	type3(links)
}

func type3(links []string) {
	c := make(chan string, len(links))

	for _, link := range links {
		go checkLink(link, c)
	}
	for l := range c {
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func type2(links []string) {
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for l := range c {
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func type1(links []string) {
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		time.Sleep(2 * time.Second)
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
