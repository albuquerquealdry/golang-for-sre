package main

import (
	"fmt"
)

type Pessoa struct {
	name string
	idade int
	salario int
}

func main() {
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
