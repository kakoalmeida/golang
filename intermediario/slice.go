package main

import "fmt"

func main() {
	// slice eh um "array dinamico"
	names := [4]string{"Frodo", "Sam", "Merry", "Pippin"}
	fmt.Println(names)

	// slices sao referencias aos arrays
	a := names[:]   // todos indices do array
	b := names[0:2] // 0, 1
	fmt.Println(a)

	b[0] = "The Bearer" // altera o indice no array e slice
	fmt.Println(b)
	c := names[2:]
	fmt.Println(c)
	c[1] = "Tuk"
	fmt.Println(names)
	d := names[:3]
	fmt.Println(d)
}
