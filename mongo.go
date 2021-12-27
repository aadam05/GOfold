package main

import (
	//"errors"
	"fmt"
)

var USD = "USD"
var EUR = "EUR"
var RUB = "RUB"

var amountUSD float32
var amountEUR float32
var amountRUB float32

var firstValue string
var secondValue string

// Проверка на неккоректные значения доллара
func MistakeChekerUSD(float32) {
	if amountUSD <= 0 {
		fmt.Println("Вы ввели нулевое или отрицательное значение")
		Convertation()
	} else {
		return
	}
}

// Проверка на неккоректные значения евро
func MistakeChekerEUR(float32) {
	if amountEUR <= 0 {
		fmt.Println("Вы ввели нулевое или отрицательное значение")
		Convertation()
	} else {
		return
	}
}

// Проверка на неккоректные значения рубля
func MistakeChekerRUB(float32) {
	if amountRUB <= 0 {
		fmt.Println("Вы ввели нулевое или отрицательное значение")
		Convertation()
	} else {
		return
	}
}

// Конвертирование введенного числа с консоли в переменную
func Convertation() {
	if firstValue == "USD" {
		fmt.Print("Сколько вы отдаёте долларов\n")
		fmt.Scan(&amountUSD)
		MistakeChekerUSD(amountUSD)
	} else if firstValue == "EUR" {
		fmt.Print("Сколько вы отдаёте евро\n")
		fmt.Scan(&amountEUR)
		MistakeChekerEUR(amountEUR)
	} else if firstValue == "RUB" {
		fmt.Print("Сколько вы отдаёте рублей\n")
		fmt.Scan(&amountRUB)
		MistakeChekerRUB(amountRUB)
	}

}

func main() {

	fmt.Print("Приветсвуем вас! Напишите из какой валюты вы хотите перевести\n")
	//Проверяем ввел ли пользователь наименование одной из валют
	//Также записываем введенное значение в переменную т.к. еще понадобится нам
	if USD == "USD" {
		fmt.Scanf("%s\n", &firstValue)
	} else if EUR == "EUR" {
		fmt.Scanf("%s\n", &firstValue)
	} else if RUB == "RUB" {
		fmt.Scanf("%s\n", &firstValue)
	}

	fmt.Print("В какую валюту вы переводите?\n")
	//Проверяем ввел ли пользователь наименование одной из валют
	//Также записываем введенное значение в переменную т.к. еще понадобится нам
	if USD == "USD" {
		fmt.Scanf("%s\n", &secondValue)
	} else if EUR == "EUR" {
		fmt.Scanf("%s\n", &secondValue)
	} else if RUB == "RUB" {
		fmt.Scanf("%s\n", &secondValue)
	}

	//Проверка валют и последовательная конвертация
	if firstValue == "USD" && secondValue == "EUR" {
		Convertation()
		amountUSD *= 0.88676
		fmt.Println("У вас теперь: ", amountUSD, "евро")
	} else if firstValue == "USD" && secondValue == "RUB" {
		Convertation()
		amountUSD *= 73.41
		fmt.Println("У вас теперь: ", amountUSD, "рублей")
	}

	//Проверка валют и последовательная конвертация
	if firstValue == "EUR" && secondValue == "USD" {
		Convertation()
		amountEUR *= 1.13
		fmt.Println("У вас теперь: ", amountEUR, "долларов")
	} else if firstValue == "EUR" && secondValue == "RUB" {
		Convertation()
		amountEUR *= 82.89
		fmt.Println("У вас теперь: ", amountEUR, "рублей")
	}

	//Проверка валют и последовательная конвертация
	if firstValue == "RUB" && secondValue == "EUR" {
		Convertation()
		amountRUB *= 0.012025
		fmt.Println("У вас теперь: ", amountRUB, "евро")
	} else if firstValue == "RUB" && secondValue == "USD" {
		Convertation()
		amountRUB *= 0.013548
		fmt.Println("У вас теперь: ", amountRUB, "долларов")
	}

	main()
}
