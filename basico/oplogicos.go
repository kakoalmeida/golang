package main

func comprar(trab1 bool, trab2 bool) (bool, bool, bool) { //indica multiplo retorno

	comprarTv50 := trab1 && trab2    // And
	comprarTv32 := trab1 != trab2    // xor
	comprarSorvete := trab1 || trab2 // or

	return comprarTv50, comprarTv32, comprarSorvete
}
