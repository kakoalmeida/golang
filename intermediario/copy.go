package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3}
	var s1 []int

	s1 = append(s1, 4, 5, 6)

	fmt.Println(arr1, s1)

	s2 := make([]int, 2)
	copy(s2, s1)

	fmt.Println(s2)
}
