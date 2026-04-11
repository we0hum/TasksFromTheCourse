package Lesson2

import (
	"errors"
	"fmt"
)

type Order struct {
	Items []string
	ID    int
}

type User struct {
	Name string
	Age  int
}

func Main() {
	// задача 1
	reportAction("Импорт данных")

	// задача 2
	runSteps(4)

	// задача 3
	sequenceDemo()

	// задача 4
	fmt.Println(calcBase())

	// задача 5
	initConfig(map[string]string{})
	initConfig(map[string]string{"host": "localhost"})

	// задача 6
	checkData([]int{10, 5, -3, 7})

	// задача 7
	loadUsers([]string{})
	loadUsers([]string{"Vasya", "Masha"})

	// задача 8
	setAge(25)
	setAge(250)

	// задача 9
	createOrder(Order{Items: []string{}})
	createOrder(Order{Items: []string{"Телефон"}})

	// задача 10
	safeRun()
	fmt.Println("Программа продолжает работу")

	// задача 11
	runValues([]int{10, 5, -3, 7})

	// задача 12
	safeCreate(Order{Items: []string{"Машина"}, ID: 1})
	safeCreate(Order{Items: []string{}})

	// задача 13
	d := []string{"db", "cache", "auth"}
	for _, v := range d {
		safeInit(v)
	}

	// задача 14
	safeTop()
	fmt.Println("программа продолжила выполнение")

	// задача 15
	safeInit1("db")
	safeInit1("cache")
	safeInit1("logger")
	safeInit1("api")

	// задача 16
	safeExecute("A", stepA)
	safeExecute("B", stepB)
	safeExecute("C", stepC)

	// задача 17
	u := User{Name: "", Age: 25}
	u1 := User{Name: "Vasya", Age: 200}
	u2 := User{Name: "Masha", Age: 30}

	safeValidate(u)
	safeValidate(u1)
	safeValidate(u2)

	// задача 18
	safeProcess([]int{1, 2, 3})
	safeProcess([]int{})

	// задача 19
	safeRun1("good")
	safeRun1("bad")
	safeRun1("crash")
	safeRun1("next")
}

// задача 1
func reportAction(action string) {
	fmt.Println("Старт:", action)
	defer fmt.Println("Завершение:", action)
	fmt.Println("В процессе:", action)
}

// задача 2
func runSteps(steps int) {
	done := 0
	defer fmt.Println("Итого выполнено шагов:", done)
	for i := 1; i <= steps; i++ {
		fmt.Println("Шаг", i)
		done++
	}
}

// задача 3
func sequenceDemo() {
	fmt.Println("Старт")
	defer fmt.Println("Завершение (A)")
	fmt.Println("Обработка")
	defer fmt.Println("Завершение (B)")
	fmt.Println("Выход")
}

// задача 4
func calcBase() int {
	base := 10
	defer func() {
		base += 5
	}()
	return base
}

// задача 5
func initConfig(config map[string]string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ошибка:", r)
		}
	}()
	if _, ok := config["host"]; ok {
		fmt.Println("конфигурация загружена")
	} else {
		panic("отсутствует параметр host в конфигурации")
	}
}

// задача 6
func checkData(data []int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ошибка:", r)
		}
	}()
	for _, datum := range data {
		if datum > 0 {
			fmt.Println("ОК:", datum)
		} else {
			panic("обнаружено отрицательное значение")
		}
	}
}

// задача 7
func loadUsers(users []string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ошибка:", r)
		}
	}()
	if len(users) == 0 {
		panic("список пользователей пуст")
	} else {
		fmt.Println("загружено пользователей:", len(users))
	}
}

// задача 8
func setAge(age int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ошибка:", r)
		}
	}()
	if age < 0 || age > 130 {
		panic("возраст вне допустимого диапазона")
	} else {
		fmt.Println("возраст утановлен:", age)
	}
}

// задача 9
func createOrder(o Order) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ошибка:", r)
		}
	}()
	if len(o.Items) == 0 {
		panic("невозможно создать заказ без товаров")
	} else {
		fmt.Println("заказ успешно создан")
	}
}

// задача 10
func criticalAction() {
	panic("критическая ошибка выполнения")
}
func safeRun() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Восстановлено после сбоя:", r)
		}
	}()
	criticalAction()
}

// задача 11
func processValue(v int) {
	if v < 0 {
		panic("обнаружено отрицательное значение")
	} else {
		fmt.Println("ОК", v)
	}
}
func runValues(data []int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ошибка при обработке:", r)
		}
	}()
	for _, datum := range data {
		processValue(datum)
	}
}

// задача 12
func createOrder1(o Order) {
	if len(o.Items) == 0 {
		panic("заказ без товаров")
	} else {
		fmt.Printf("заказ оформлен: #%d\n", o.ID)
	}
}
func safeCreate(o Order) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ошибка создания заказа:", r)
		}
	}()
	createOrder1(o)
}

// задача 13
func initComponent(name string) {
	if name == "cache" {
		panic("ошибка инициализации cache")
	} else {
		fmt.Println("компонент инициализирован:", name)
	}
}
func safeInit(name string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("восстановлено после сбоя в компоненте:", name)
		}
	}()
	initComponent(name)
}

// задача 14
func deepError() {
	panic("внутренняя ошибка")
}
func middle() {
	deepError()
	fmt.Println("выполняется middle")
}
func safeTop() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("восстановление после паники:", r)
		}
	}()
	middle()
}

// задача 15
func initModule(name string) error {
	switch name {
	case "logger":
		return fmt.Errorf("ошибка при инициализации")
	case "cache":
		panic("модуль cache не загрузился")
	default:
		fmt.Println("модуль инициализирован:", name)
	}
	return nil
}
func safeInit1(name string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("восстановлено после сбоя:", name)
		}
	}()
	err := initModule(name)
	if err != nil {
		fmt.Println(err, name)
	}
}

// задача 16
func safeExecute(name string, fn func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("восстановлено после сбоя в:", name)
		}
	}()
	fn()
}
func stepA() {
	fmt.Println("Шаг A выполнен")
}
func stepB() {
	panic("ошибка на шаге B")
}
func stepC() {
	fmt.Println("Шаг C выполнен")
}

// задача 17
func validateUser(u User) error {
	if u.Name == "" {
		return errors.New("Имя не указано")
	} else if u.Age < 0 || u.Age > 130 {
		panic("возраст вне диапазона")
	} else {
		fmt.Println("пользователь корректен")
	}
	return nil
}
func safeValidate(u User) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("фатальная ошибка при проверке:", r)
		}
	}()
	err := validateUser(u)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}

// задача 18
func processData(data []int) {
	fmt.Println("начало обработки")
	defer fmt.Println("обработка завершена")
	if len(data) == 0 {
		panic("нет данных для обработки")
	}
	for _, v := range data {
		fmt.Println("элемент:", v)
	}
}
func safeProcess(data []int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("восстановлено после сбоя:", r)
		}
	}()
	processData(data)
}

// задача 19
func runOperation(name string) error {
	switch name {
	case "bad":
		return fmt.Errorf("ошибка при выполнении")
	case "crash":
		panic("фатальный сбой")
	default:
		fmt.Println("операция успешно выполнена:", name)
	}
	return nil
}
func safeRun1(name string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("восстановление после паники в:", name)
		}
	}()
	err := runOperation(name)
	if err != nil {
		fmt.Println(err, name)
	}
}
