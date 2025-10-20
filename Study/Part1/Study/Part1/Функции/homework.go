package main

import "fmt"

var global int64 = 190

func kor() {
	increment()
	fmt.Println("Global:", global)
}

func increment() {
	global += 10
}
