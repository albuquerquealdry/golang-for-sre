package main

import "fmt"

type Empregado struct {
	Name string
	id   int
}

func main() {
	emp := Empregado{"Joao", 123}
	pts := &emp

	fmt.Println(pts)
	pts.Name = "Paulo"
	fmt.Println("valor pts:", pts)
}
