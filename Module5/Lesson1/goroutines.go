package Lesson1

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	Name     string
	Messages int
}

func Main() {
	// задача 1
	go printMessage()
	fmt.Println("Привет из  main")
	time.Sleep(1 * time.Millisecond)

	// задача 2
	go func() {
		fmt.Println("Привет из первой")
	}()
	go func() {
		fmt.Println("Привет из второй")
	}()
	go func() {
		fmt.Println("Привет из третьей")
	}()

	// задача 3
	go func(n int) {
		fmt.Println("Число:", n)
	}(5)

	// задача 4
	for i := 1; i <= 5; i++ {
		fmt.Println("Горутина номер:", i)
	}

	// задача 5
	go doTask("логирование")
	go doTask("кеширование")
	go doTask("загрузка")

	// задача 6
	go fmt.Println("hi")
	time.Sleep(3 * time.Millisecond)

	// задача 7
	go fastTask("1")
	go slowTask("1")
	go fastTask("2")
	go slowTask("2")
	time.Sleep(2 * time.Second)

	// задача 16
	m := make(map[string]int)
	m["Телефон"] = 1000
	m["Чехол"] = 200
	m["Кабель"] = 150
	for s, i := range m {
		go fmt.Printf("Товар %s стоит %d\n", s, i)
	}
	time.Sleep(1 * time.Second)

	// задача 17
	buyers := []string{"Vasya", "Masha", "Petr", "Olya"}
	products := []string{"Телефон", "Чехол", "Кабель", "Наушники"}

	for _, b := range buyers {
		go buyer(b, products)
	}
	time.Sleep(2 * time.Second)

	// задача 18
	go sumRange(1, 10)
	go sumRange(11, 20)
	go sumRange(21, 30)
	time.Sleep(1 * time.Second)

	// задача 19
	go checkWord("Apple")
	go checkWord("Go")
	go checkWord("Admin")
	time.Sleep(1 * time.Second)

	// задача 20
	u := []User{
		{"Vasya", 3},
		{"Masha", 5},
		{"Petr", 2},
	}
	for _, user := range u {
		fmt.Printf("%s отправил %d сообщений\n", user.Name, user.Messages)
	}
	time.Sleep(1 * time.Second)
}

// задача 1
func printMessage() {
	fmt.Println("Привет из горутины")
}

// задача 5
func doTask(task string) {
	fmt.Println("Начинаю задачу:", task)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Завершаю задачу:", task)
}

// задача 7
func fastTask(name string) {
	fmt.Println("Быстрая задача", name)
}
func slowTask(name string) {
	fmt.Println("Медленная задача", name)
	time.Sleep(1 * time.Second)
	fmt.Println("Медленная задача", name, "завершена")
}

// задача 17
func buyer(name string, products []string) {
	product := products[rand.Intn(len(products))]

	fmt.Printf("Покупатель %s купил %s\n", name, product)
}

// задача 18
func sumRange(start, end int) {
	allSum := 0
	for i := start; i <= end; i++ {
		allSum += i
	}
	fmt.Printf("Диапозон %d-%d: сумма %d\n", start, end, allSum)
}

// задача 19
func checkWord(word string) {
	s := []rune(word)
	if s[0] == 'A' {
		fmt.Println(word, "- Да")
	} else {
		fmt.Println(word, "- Нет")
	}
}
