package main

import "fmt"

func main() {
	fmt.Println("main 1")
	defer func() {
		fmt.Println("main 1 end")
	}()
	def1()
	fmt.Println("main 2")
	def2()
	fmt.Println("main 3")
	def3()
}

func def1() {
	fmt.Println("defer1")
	defer func() {
		fmt.Println("defer1(1)")
	}()
}
func def2() {
	fmt.Println("defer2")
	defer func() {
		fmt.Println("defer2(1)")
	}()
}
func def3() {
	fmt.Println("defer3")
	defer func() {
		fmt.Println("defer3(1)")
	}()
}
