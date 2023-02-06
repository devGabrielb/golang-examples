package main

import "fmt"

var words = []string{
	"Amigo",
	"Casa",
	"Jardim",
	"Sorvete",
	"Computador",
	"Livro",
	"Pássaro",
	"Carro",
	"Cachorro",
	"Gato",
	"Bicicleta",
	"Céu",
	"Mar",
	"Árvore",
	"Flores",
	"Nuvem",
	"Estrela",
	"Lua",
	"Sol",
	"Montanha",
	"Rio",
	"Praia",
	"Parque",
	"Cidade",
	"Campo",
	"Lago",
	"Rio",
	"Pássaro",
	"Jardim",
	"Amigo"}

func main() {

	fmt.Println(Frequency(words))
	ExampleCommaOk()

}

func Frequency(words []string) map[string]int {
	freq := make(map[string]int)

	for _, w := range words {
		freq[w]++

	}
	return freq

}

func ExampleCommaOk() {
	prices := map[string]int{
		"banana": 2,
	}

	price, ok := prices["banana"]

	if ok {
		fmt.Printf("The price of banana is $%d\n", price)
	} else {

		fmt.Printf("We don't have bananas")
	}

	apple, ok := prices["apple"]

	if ok {
		fmt.Printf("The price of apple is $%d\n", apple)
	} else {

		fmt.Printf("We don't have apples")
	}

}
