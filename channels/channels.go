package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go exampleGoroutine("Hello World", channel)
	fmt.Println("Inicío do programa")

	// for {

	// Recebendo mensagem do canal
	// 	msg, ok := <-channel;
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	for msg := range channel {
		fmt.Println(msg)
	}

	fmt.Println("Fim do programa")

}

func exampleGoroutine(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		//enviando mensagem para o canal
		channel <- fmt.Sprintf("%s: %d", text, i)
		time.Sleep(time.Second)
	}

	defer close(channel)
}
