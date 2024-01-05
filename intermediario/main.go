package main

import (
	"fmt"
	"math/rand"
	"time"
) // Utilização de arrays

func main() {
	// arrays homogeneos e estaticos
	var numeros [6]int

	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	// for i := 0; i < len(numeros); i++ {
	// 	numeros[i] = r.Intn(61)
	// }

	for i := range numeros {
		numeros[i] = r.Intn(61)
	}

	fmt.Println(numeros)

	dias := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, dia := range dias {
		fmt.Println(dia)
	}
}
