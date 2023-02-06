package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go func() {
		exampleGoroutine("Hello World")
		wg.Done()
	}()

	go func() {
		exampleGoroutine("Olá mundo")
		wg.Done()
	}()

	wg.Wait()
}

func exampleGoroutine(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text, ":", i+1)
		time.Sleep(time.Second)
	}
}
