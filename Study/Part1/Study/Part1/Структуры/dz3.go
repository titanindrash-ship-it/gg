package main

import "fmt"

type Apartment struct {
	Address string
	Area    int
	Floor   int
	Price   int
}

func NewApartment(
	address string,
	area int,
	floor int,
	price int,
) Apartment {
	if address == "" {
		fmt.Println("Пустой адрес(")
		return Apartment{}
	}
	if area < 10 {
		fmt.Println("Неверная площадь квартиры")
		return Apartment{}
	}
	if floor <= 0 || floor >= 100 {
		fmt.Println("Неверно указан этаж")
		return Apartment{}
	}
	if price < 0 {
		fmt.Println("Неверная цена")
		return Apartment{}
	}
	return Apartment{
		Address: address,
		Area:    area,
		Floor:   floor,
		Price:   price,
	}
}

func (a Apartment) giveAddress() {
	fmt.Println("Адрес квартиры:", a.Address)
}
func (a Apartment) giveArea() {
	fmt.Println("Площадь квартиры:", a.Area)
}
func (a *Apartment) changePrice(price int) {
	if a.Price < 0 {
		fmt.Println("Указана неверная цена")
		return
	}
	a.Price = price
}

func main() {
	Apartment := NewApartment(
		"ул. Московская",
		69,
		13,
		10000,
	)
	fmt.Println("До изменения", Apartment)
	Apartment.changePrice(1000)
	fmt.Println("После изменения", Apartment)
	Apartment.giveAddress()
	Apartment.giveArea()

}
