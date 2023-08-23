package main

import (
	"fmt"
	"os"
)

type Pessoa struct {
	name string
	idade int
	salario int
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Erro numero nao vailido")
		os.Exit(1)
	}
	fmt.Println("numero convertido", n)

	salarios := []int{}

	fmt.Println(salarios)
	salFunc := make(map[string]int)
	salFunc["snowden"] = 100000
	salFunc["Mitinik"] = 12222

	sal, exist := salFunc["snowden"]
	fmt.Println(sal, exist)

	name := modulation()
	name, coin, location := "Jesus Tricolor", 1000000, "Portugal"
	fmt.Println(name,coin, location )

	pessoa := Pessoa{
		name: "teste",
		idade: 10,
		salario: 1000,
	}

	fmt.Println(pessoa)
}

func modulation() string {
	return "GOLANG FOR SRE"
}

