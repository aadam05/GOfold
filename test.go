package main

import (
	"fmt"
)

type Area interface {
	getArea() int
}

type Sides struct {
	lateralSide int
	mainSide    int
}

func (side *Sides) getArea() int {
	area := side.lateralSide * side.mainSide
	return area
}

func main() {
	fmt.Print("Вычисляем площадь квадрата\n")
	total := Sides{
		lateralSide: 5,
		mainSide:    7,
	}
	fmt.Println("Площадь равна", total.getArea())
}
