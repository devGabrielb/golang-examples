package main

import (
	"fmt"
	"time"
)

func main() {

	go exampleGoroutine("Hello World goroutine 1")
	exampleGoroutine("Hello World goroutine 2")
}

func exampleGoroutine(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
