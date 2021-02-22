package main

import (
	"fmt"
	"net/http"
)

func main() {
	webSites := []string{
		"https://www.facebook.com",
		"https://www.google.com",
		"https://www.amazon.in",
		"https://twitter.com",
		"https://www.beartobull.in",
	}
	for _, link := range webSites {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("I think %s is down :/ Error: %d \n", link, err)
		return
	}
	fmt.Printf("%s is up & running fine :)\n", link)
}
