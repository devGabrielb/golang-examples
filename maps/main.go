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

}

func Frequency(words []string) map[string]int {
	freq := make(map[string]int)

	for _, w := range words {
		freq[w]++

	}
	return freq

}
