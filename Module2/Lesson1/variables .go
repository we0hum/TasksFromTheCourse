package Lesson1

import "fmt"

func Main() {
	//задача 1
	city := "Moscow"
	name := "Alex"
	fmt.Println("Hello,", name, "from", city)

	//задача 2
	temperature := 21.5
	fmt.Println()
	fmt.Printf("Current temperature is %.1f°C\n", temperature)

	//задача 3
	birthYear := 1986
	currentYear := 2000
	age := currentYear - birthYear
	fmt.Println()
	fmt.Printf("You are %d years old.\n", age)

	//задача 4
	word := "Привет"
	fmt.Println()
	fmt.Println("Runes:", len([]rune(word)))
	fmt.Println("Bytes:", len(word))

	//задача 5
	letter := 'A'
	fmt.Println()
	fmt.Printf("Code: %d\n", letter)
	fmt.Printf("Char: %c\n", letter)

	//задача 6
	const Pi = 3.1415
	r := 10.0
	fmt.Println()
	fmt.Println("Circle length:", 2*Pi*r)

	//задача 7
	first := "Go"
	second := "lang"
	fmt.Println()
	fmt.Println(first + second)

	//задача 8
	lang := "Go"
	fmt.Println()
	fmt.Println(lang + lang + lang)

	//задача 9
	nameNew := "Alexander"
	fmt.Println()
	fmt.Printf("Hello, %s! Your name has %d letters.\n", nameNew, len(nameNew))

	//задача 10
	price := 199.9
	count := 3
	total := float64(count) * price
	fmt.Println()
	fmt.Println("Total:", total)

	//задача 11
	age1 := 17
	canDrive := age1 >= 18
	fmt.Println()
	fmt.Println("Can drive:", canDrive)

	//задача 12 Мини-проект «Профиль пользователя»
	name2 := "Bob"
	city2 := "Moscow"
	age2 := 20
	isStudent := true
	fmt.Println()
	fmt.Printf("--- Profile ---\nName: %s\nAge: %d\nCity: %s\nStudent: %t\n", name2, age2, city2, isStudent)

	//Pro-задачи
	//задача 13
	usd := 50
	rate := 80.46
	rub := float64(usd) * rate
	fmt.Println()
	fmt.Printf("%d USD = %.2f RUB\n", usd, rub)

	//задача 14
	a := "Go"
	b := "Golang"
	result1 := a == b
	fmt.Println()
	fmt.Println("Equal:", result1)

	//задача 15
	n := 12345
	last := n % 10

	first1 := n
	for first1 >= 10 {
		first1 /= 10
	}

	fmt.Println()
	fmt.Println("First:", first1)
	fmt.Println("Last:", last)
	fmt.Println()
}
