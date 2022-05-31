package main

import "fmt"

func main() {
	// инициализация переменной с указанием типа
	var num int = 5
	fmt.Printf("Инициализация переменной с указанием типа: %d \n\n", num)

	// инициализация переменной без указания типа
	var nextNum = 10.5
	fmt.Printf("инициализация переменной без указания типа: %g \n\n", nextNum)

	// короткая инициализация
	thirdNum := 10
	fmt.Printf("короткая инициализация: %d \n\n", thirdNum)

	// инициализация нескольких переменных, первый вариант
	var isMale, age = false, 16
	fmt.Printf("инициализация нескольких переменных, первый вариант: %t, %d \n\n", isMale, age)

	// инициализация нескольких переменных, второй вариант
	var (
		year  = 2022
		month = 5
	)
	fmt.Printf("инициализация нескольких переменных, второй вариант: %d, %d \n\n", year, month)

	// инициализация нескольких переменных с указанием типа
	var (
		isFemale bool   = true
		city     string = "Moscow"
	)
	fmt.Printf("инициализация нескольких переменных с указанием типа: %t, %s \n\n", isFemale, city)

	// отложенная инициализация
	var name string

	if isFemale {
		name = "Елизавета"
	}

	if isMale {
		name = "Андрей"
	}
	fmt.Printf("отложенная инициализация: %s", name)
}
