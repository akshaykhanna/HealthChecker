package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	webSites := []string{
		"https://www.facebook.com",
		"https://www.google.com",
		"https://www.amazon.in",
		"https://twitter.com",
		"https://www.beartobull.in",
	}
	channel := make(chan string)
	for _, link := range webSites {
		go checkLink(link, channel)
	}
	for channelOutput := range channel {
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLink(link, channel)
		}(channelOutput)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("I think %s is down :/ Error: %d \n", link, err)
		c <- link
		return
	}
	fmt.Printf("%s is up & running fine :)\n", link)
	c <- link
}
