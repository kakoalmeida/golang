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
	//não um operador ao estilo das outra linguagens mas pode ser feito de tal forma
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

	// nome := "Almeida"
	// tamanhoString(nome)

	// //fmt.Printf("%v %v %v %q %v \n", a, b, c, d, *e)

	// aritmetica()

	// operador ternário
	//fmt.Println(resultado(8.94))
	ponteiros()
	fmt.Println(notasConceito(8, 8))

	if numero := randNumber(); numero >= 70 {
		fmt.Println(numero)
	}

	repeticao()

	// vars()
	// tv50, tv32, sorvete := comprar(true, false)
	// fmt.Printf("TV 50: %t, Tv 32: %t, Sorvete: %t, Ficar sem sorvete: %t \n",
	// 	tv50, tv32, sorvete, !sorvete)

	// %t = representa boolen em print, o resto é igual a C
}
