package main

import "fmt"

func vars() {

	var b byte // declaração
	b = 5      // atribuição padrao
	i := 0     // inicialização e declaração com inferencia

	var j = 2 // inferencia

	i += j * int(b) // cast

	fmt.Print(j, i, b)
}
