package main

import (
	"fmt"
	"time"
)

func main() {
	name("Василий")
	name("Дмитрий")
	name("Антон")
	name("Евгений")
}
func name(name string) {
	fmt.Println("Добрый день, уважаемый", name)
	fmt.Println("Проходите за данный столик")
	fmt.Println("Вот ваше меню, как будете готоы сделать заказ, нажмите на кнопку")
	fmt.Println("----------------------------------------")
	time.Sleep(500 * time.Millisecond)
}
