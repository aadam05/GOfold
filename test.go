package main

import (
	"fmt"
)

//Создаём интерфейс и определяем в нем метод
type Area interface {
	getArea() int
}

//Создаём структуру с 2 переменными
type Sides struct {
	lateralSide int
	mainSide    int
}

//Реализуем метод интерфейса
func (side *Sides) getArea() int {
	area := side.lateralSide * side.mainSide
	return area
}

//Вычисление и вывод программы
func main() {
	fmt.Print("Вычисляем площадь квадрата\n")
	total := Sides{
		lateralSide: 5,
		mainSide:    7,
	}
	fmt.Println("Площадь равна", total.getArea())
}
