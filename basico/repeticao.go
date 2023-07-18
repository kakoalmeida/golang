package main

import "fmt"

func repeticao() {

	// não existe while em Go, apenas for
	// o for pode ser utilizado das seguintes maneiras abaixo

	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	for i := x; i > 0; i-- {
		fmt.Println(i)
	}

}
