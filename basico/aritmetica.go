package main

import "fmt"

func aritmetica() {
	a := 2
	b := 2

	fmt.Printf("Sum %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Sub %d - %d = %d\n", a, b, a-b)
	fmt.Printf("Mult de %d x %d = %d\n", a, b, a*b)
	fmt.Printf("Div de %d / %d = %.2f\n", a, b, float32(a/b))
	fmt.Printf("Rest de %d e %d = %v\n", a, b, a%b)

	//bitwise

	fmt.Println("AND =>", a&b)
	fmt.Println("OR =>", a|b)
	fmt.Println("XOR =>", a^b)

	//math
	// math.Max() = maior valor; math.Min() = menor valor; math.Pow() = potencia
}
