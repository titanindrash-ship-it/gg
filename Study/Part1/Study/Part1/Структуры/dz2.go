/*package main

import "fmt"

type Manga struct {
	Name  string
	Avtor string
	Style string
	Size  int
	Anime bool
}

func NewManga(
	Name string,
	Avtor string,
	Style string,
	Size int,
	Anime bool,
) Manga {
	if Name == "" {
		fmt.Println("Name не прошел валидацию(")
		return Manga{}
	}
	if Avtor == "" {
		fmt.Println("Avtor не прошел валидацию(")
		return Manga{}
	}
	if Style == "" {
		fmt.Println("Style не прошел валидацию(")
		return Manga{}
	}
	if Size < 10 || Size > 1500 {
		fmt.Println("Size не прошел валидацию(")
		return Manga{}
	}
	return Manga{
		Name:  Name,
		Avtor: Avtor,
		Style: Style,
		Size:  Size,
		Anime: Anime,
	}
}

func (m *Manga) changeManga(Name string) {
	if m.Name == "" {
		fmt.Println("Передано пустое название манги(")
		return
	}
	m.Name = Name
}
func (m *Manga) changeAvtor(Avtor string) {
	if m.Avtor == "" {
		fmt.Println("Передано пустое имя автора(")
		return
	}
	m.Avtor = Avtor
}
func (m *Manga) changeStyle(Style string) {
	if m.Style == "" {
		fmt.Println("Передан пустой стиль манги(")
		return
	}
	m.Style = Style
}
func (m *Manga) changeAnimeYes() {
	m.Anime = true
}
func (m *Manga) changeAnimeNo() {
	m.Anime = false
}
func (m *Manga) changeSize(Size int) {
	m.Size = Size
}

func main() {
	Manga := NewManga(
		"Hanter x Hanter",
		"Ёсихиро Тагаси",
		"Сёнен",
		410,
		true)

	fmt.Println("Manga до изменения", Manga)

	Manga.changeManga("Токийский гуль")

	Manga.changeAvtor("Исида")

	Manga.changeStyle("Психология")

	Manga.changeAnimeYes()

	Manga.changeSize(250)

	Manga.changeAnimeYes()

	fmt.Println("Manga после изменения", Manga)
}
