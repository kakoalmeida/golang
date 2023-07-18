package main

import "fmt"

func ponteiros() {
	i := 1

	var p *int = nil // nome do ponteiro e tipo; nil == null

	p = &i // acessa endereço
	//*p++   // acessa valor

	i += 3

	fmt.Printf("Endereço de p = %v\n Endereço de i = %v\n Valor de *p = %d\n Valor de i = %d\n", p, &i, *p, i)
}
