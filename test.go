package main

import "strconv"

type Carro struct {
	modelo string
	ano    int
	cor    string
}

func eachBit(value string) {
	println("Cada bit")
	for key := range value {
		println(value[key])
	}
	println("------------------")
}

func eachLetter(value string) {
	println("Cada letra")
	for key := range value {
		println(string(value[key]))
	}
	println("------------------")
}

func firstTwoLetters(value string) {
	println("Primeiras 2 letras")
	println(string(value[0:2]))
	println("------------------")
}

func listaModelos(value []Carro) {
	println("Modelos disponiveis")
	for key := range value {
		println(value[key].modelo)
	}
	println("------------------")
}

func listaModelosAno(ano int, value []Carro) {
	println("Modelos do ano " + strconv.Itoa(ano))
	for key := range value {
		if value[key].ano == ano {
			println(value[key].modelo)
		}
	}
	println("------------------")
}

func main() {
	//foo := "Hello"
	carros := []Carro{
		Carro{"Gol", 2018, "roxo"},
		Carro{"Civic", 2016, "roxo"},
		Carro{"Ka", 2019, "roxo"},
		Carro{"308", 2018, "roxo"},
		Carro{"saveiro", 2020, "roxo"},
	}

	println("------------------")

	//eachLetter(foo)
	//eachBit(foo)
	//firstTwoLetters(foo)

	listaModelos(carros)
	listaModelosAno(2019, carros)
}
