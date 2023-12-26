package main

import (
	"fmt"
	"math/rand"
	"time"
)

// m "math" importa lib math como 'm'

func tamanhoString(word string) {
	fmt.Println(len(word))
}

func resultado(nota float64) string {
	//não tem um operador ao estilo das outra linguagens mas pode ser feito de tal forma
	if nota >= 7 { //utiliza parenteses apenas em precedencia ou expressões
		return "Aprovado"
	}
	return "Reprovado"
}

func notasConceito(notaUm float64, notaDois float64) string {
	//media := (notaUm + notaDois) / 2 // inicialização pode ser na condição, igual a um for

	if media := (notaUm + notaDois) / 2; media >= 8 { // só existe no bloco
		return "A+"
	} else if media >= 6 && media < 8 {
		return "A"
	} else if media >= 4 && media < 6 {
		return "B"
	} else {
		return "C"
	}

}

func randNumber() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(100)
}

func ponteiros() {
	i := 1

	var p *int = nil // nome do ponteiro e tipo; nil == null

	p = &i // acessa endereço
	//*p++   // acessa valor

	i += 3

	fmt.Printf("Endereço de p = %v\nEndereço de i = %v\nValor de *p = %d\nValor de i = %d\n", p, &i, *p, i)
}

func repeticao() {

	// não existe while em Go, apenas for
	// o for pode ser utilizado das seguintes maneiras abaixo

	x := 0
	for x < 10 {
		fmt.Print(x)
		x++
	}
	println("")
	for i := x; i > 0; i-- {
		fmt.Print(i)
	}

}

func main() {

	// name := "kako"

	// const PI float64 = 3.1415
	// var raio = 3.8
	// area := PI * m.Pow(raio, 2)
	// fmt.Printf("Area %.2f\n", area)

	// var a int
	// var b float32
	// var c bool
	// var d string
	// var e *int

	// a = 10

	// e = &a

	nome := "Almeida"
	tamanhoString(nome)

	ponteiros()
	fmt.Println(notasConceito(3.25, 7.3))

	if numero := randNumber(); numero >= 70 {
		fmt.Println(numero)
	}

	repeticao()

}
