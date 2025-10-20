/*package main

import "fmt"

func main() {
	number := 10
	strin := "stroka"
	bool := false
	number64 := 10.5
	fmt.Println(number, strin, bool, number64)
	fmt.Println("-----------------")
	num(&number)
	str(&strin)
	bol(&bool)
	num64(&number64)
	fmt.Println(number, strin, bool, number64)
	fmt.Println("-----------------")
}

func num(ptr *int) {
	*ptr = 5
}

func str(ptr *string) {
	*ptr = "change"
}

func bol(ptr *bool) {
	*ptr = true
}

func num64(ptr *float64) {
	*ptr = 9.4
}
