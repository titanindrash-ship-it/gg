package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	score := 0
	i := 1
	fmt.Println("Get ready?")
	for {

		fmt.Println("----------------------")
		fmt.Println("Вы подлетели к трубе", i)
		fmt.Println("Вы пролетаете трубу", i)
		fmt.Println("Вы пролетели трубу", i)
		i++
		score++
		fmt.Println("Ваш счет:", score)
		time.Sleep(200 * time.Millisecond)
		if rand.IntN(20) == 1 {
			break
		}
	}
}
