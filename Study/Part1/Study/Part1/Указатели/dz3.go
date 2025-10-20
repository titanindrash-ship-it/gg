/*package main

import "fmt"

func main() {
	// number := 10
	text := "stroka"
	boolean := false
	drob := 10.5

	var ptrNum2 *int = nil
	// var ptrText2 *string = nil
	// var ptrBol2 *bool = nil
	// var ptrDrob2 *float64 = nil

	// ptrInt := &number
	ptrString := &text
	ptrBool := &boolean
	ptrFloat64 := &drob

	num(ptrNum2)
	str(ptrString)
	bol(ptrBool)
	num64(ptrFloat64)

}

func num(ptr *int) {
	fmt.Println("int pnt:", ptr)
	if ptr == nil {
		fmt.Println("int ptr является нил указателем(")
		return
	}
	*ptr = 5
	fmt.Println("int pnt:", *ptr)
}

func str(ptr *string) {
	fmt.Println("int pnt:", ptr)
	if ptr == nil {
		fmt.Println("int ptr является нил указателем(")
		return
	}
	*ptr = "change"
	fmt.Println("int pnt:", *ptr)
}

func bol(ptr *bool) {
	fmt.Println("int pnt:", ptr)
	if ptr == nil {
		fmt.Println("int ptr является нил указателем(")
		return
	}
	*ptr = true
	fmt.Println("int pnt:", *ptr)
}

func num64(ptr *float64) {
	fmt.Println("int pnt:", ptr)
	if ptr == nil {
		fmt.Println("int ptr является нил указателем(")
		return
	}
	*ptr = 9.4
	fmt.Println("int pnt:", *ptr)
}
