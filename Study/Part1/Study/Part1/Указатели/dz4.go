package main

import "fmt"

func main() {
	boobs := 1.5
	fmt.Println("Do:", boobs)
	boobsUp(&boobs)
	fmt.Println("Posle:", boobs)
}

func boobsUp(ptr *float64) {
	if ptr == nil {
		return
	}
	*ptr += 1.5
}
