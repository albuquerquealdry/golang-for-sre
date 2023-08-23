package main

import "fmt"

func pointFunction(a *int) {
	*a = 400
}

func main() {
	var x = 100
	fmt.Print("O valor da variavel x antes da funcao:", x)
	var pa *int = &x

	pointFunction(pa)
	fmt.Print(" O valor de x depois da funcao:", x)
}
