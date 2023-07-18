package main

import "fmt"

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
	fmt.Println(resultado(8.94))
	ponteiros()

	// vars()
	// tv50, tv32, sorvete := comprar(true, false)
	// fmt.Printf("TV 50: %t, Tv 32: %t, Sorvete: %t, Ficar sem sorvete: %t \n",
	// 	tv50, tv32, sorvete, !sorvete)

	// %t = representa boolen em print, o resto é igual a C
}
