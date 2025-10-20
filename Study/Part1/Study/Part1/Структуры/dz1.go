/*package main

import "fmt"

type Manga struct {
	Name  string
	Avtor string
	Style string
	Size  int
	Anime bool
}

func main() {
	HxH := Manga{
		Name:  "Hanter x Hanter",
		Avtor: "Ёсихиро Тагаси",
		Style: "Сёнен",
		Size:  410,
		Anime: true,
	}
	fmt.Println("До изменения")
	fmt.Println("Manga:", HxH)
	fmt.Println("----------------")
	HxH.Size = 666
	fmt.Println("После изменения")
	fmt.Println("Manga:", HxH)
}
