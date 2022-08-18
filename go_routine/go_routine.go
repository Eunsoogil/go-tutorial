package main

import (
	"fmt"
	"net/http"
	"time"
)

var links = []string{
	"https://google.com",
	"https://facebook.com",
	"https://stackoverflow.com",
	"https://golang.org",
	"https://amazon.com",
}

func main() {
	// case1()
	// case2()
	// case3()
	// case4()
	// case5()
	final()
}

func case1() {
	for _, link := range links {
		checkLink(link)
	}
}

func case2() {
	for _, link := range links {
		go checkLink(link)
	}
}

func case3() {

	c := make(chan string)

	for _, link := range links {
		go checkLinkWithChannel(link, c)
	}

	fmt.Println(<-c)
}

func case4() {

	c := make(chan string)

	for _, link := range links {
		go checkLinkWithChannel(link, c)
	}

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println(<-c)
}

func case5() {

	c := make(chan string)

	for _, link := range links {
		go checkLinkWithChannel(link, c)
	}

	// for {
	// 	go checkLinkWithChannel(<-c, c)
	// }

	// 위와 동일 코드
	for l := range c {
		go checkLinkWithChannel(l, c)
	}
}

func final() {
	c := make(chan string)

	for _, link := range links {
		go checkLinkWithChannel(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(3 * time.Second)
			checkLinkWithChannel(link, c)
		}(l)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is not connected!")
	}

	fmt.Println(link, "is connected")

}

func checkLinkWithChannel(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is not connected!")
		c <- link
	}

	fmt.Println(link, "is connected")
	c <- link

}
