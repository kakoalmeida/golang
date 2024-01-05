package main

import "fmt"

func main() {

	phoneList := make(map[int]string)

	phoneList[55991418706] = "kaue"
	phoneList[55991781543] = "eduarda"
	phoneList[5533331790] = "zaida"

	for phone, name := range phoneList {
		fmt.Printf("%s - telefone %d\n", name, phone)
	}

	fmt.Println(phoneList[55991418706])
	delete(phoneList, 55991418706)
	fmt.Println(phoneList[55991418706])
}
