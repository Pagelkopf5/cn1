package main

import (
	"sort"
)

type Estado struct {
	nome    string
	tamanho int
}

func main() {
	listaEstados := []Estado{
		Estado{"Pará", 1247},
		Estado{"Amazonas", 1570},
		Estado{"Bahia", 564},
		Estado{"Mato Grosso", 903},
		Estado{"Minas Gerais", 587},
		Estado{"Mato Grosso do Sul", 358},
		Estado{"Goiás", 340},
		Estado{"Maranhão", 332},
		Estado{"Rio Grande do Sul", 281},
		Estado{"Tocantins", 277},
		Estado{"Piauí", 251},
		Estado{"São Paulo", 248},
		Estado{"Rondônia", 237},
		Estado{"Roraima", 224},
		Estado{"Paraná", 199},
		Estado{"Acre", 152},
		Estado{"Ceará", 148},
		Estado{"Amapá", 142},
		Estado{"Pernambuco", 98},
		Estado{"Santa Catarina", 95},
		Estado{"Paraíba", 56},
		Estado{"Rio Grande do Norte", 52},
		Estado{"Espírito Santo", 46},
		Estado{"Rio de Janeiro", 43},
		Estado{"Alagoas", 27},
		Estado{"Sergipe", 21},
		Estado{"Distrito Federal", 5},
	}

	sort.Slice(listaEstados, func(i, j int) bool {
		return listaEstados[i].tamanho > listaEstados[j].tamanho
	})

	maioresEstados := listaEstados[0:10]

	for key := range maioresEstados {
		println(maioresEstados[key].nome)
	}

}
