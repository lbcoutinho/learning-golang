package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://facebook.com",
		"http://amazon.com",
		"http://apple.com",
		"http://netflix.com",
		"http://google.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkStatus(url, c)
	}

	for u := range c {
		go func(url string) {
			time.Sleep(5 * time.Second)
			checkStatus(url, c)
		}(u)
	}

}

func checkStatus(url string, c chan string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		c <- url
		return
	}
	fmt.Println(url, "-", res.Status)
	c <- url
}
