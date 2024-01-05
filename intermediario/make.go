package main

import "fmt"

//similar com malloc

func main() {
	a := make([]int, 10) //array vazio de tamanho 10
	fmt.Println(len(a), cap(a))
	b := make([]float64, 0, 10) //tamanho 0 e capacidade 10
	fmt.Println(len(b), cap(b))

	a[0] = 1
	a[9] = 10
	fmt.Println(a)

	b = append(b, 1.1, 2.2, 3.3, 4.4)
	fmt.Println(b)
	fmt.Println(b[0])
	fmt.Println(len(b))
	b = append(b, 5.5, 6.6, 7.7, 8.8, 9.9, 10.0, 11.1, 5.5)
	fmt.Println(b)
	fmt.Println(len(b), cap(b))

}
