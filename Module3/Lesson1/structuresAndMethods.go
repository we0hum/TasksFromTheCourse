package Lesson1

import (
	"errors"
	"fmt"
)

type User struct {
	Name    string
	Age     int
	Email   string
	Address Address
	Contact
	Orders []Order
}

type Address struct {
	City, Street string
}

type Company struct {
	Name string
	HQ   Address
}

type Contact struct {
	Phone string
}

type BankAccount struct {
	balance            int
	transactionHistory []int
}

type Employee struct {
	Name   string
	Age    int
	Salary int
}

type Product struct {
	Name  string
	Price int
}

type Student struct {
	Name   string
	Grades []int
}

type Order struct {
	Product string
	Price   int
	Count   int
}

type Visit struct {
	Name  string
	Count int
}

func Main() {
	//задача 1
	u1 := User{
		Name: "Vasya",
		Age:  19,
		Address: Address{
			City: "Kazan",
		},
		Contact: Contact{
			Phone: "+7-900-000-00-00",
		},
	}
	u2 := User{
		Name: "Masha",
		Age:  16,
		Address: Address{
			City: "Moscow",
		},
	}
	u3 := User{
		Name: "Petr",
		Age:  25,
	}
	u4 := User{
		Name: "Olya",
		Age:  30,
	}
	users := []User{u1, u2, u3, u4}
	for _, user := range users {
		if user.Age > 18 {
			fmt.Println(user.Name)
		}
	}

	//задача 2
	users1 := []User{u1, u2, u3, u4}
	fmt.Println("Количество пользователей:", len(users1))
	users1 = append(users1, User{
		Name: "Ivan",
		Age:  23,
	})
	fmt.Println("Количество пользователей:", len(users1))

	//задача 3
	printUser(users, "Petr")
	printUser(users, "Anna")

	//задача 4
	u1.Hello()

	//задача 5
	u1.IsAdult()
	u2.IsAdult()

	//задача 6
	u4.Birthday()

	//задача 7
	u5 := User{
		Name:  "Vasya",
		Age:   25,
		Email: "v@mail.ru",
	}
	fmt.Println(u5.Card())

	//задача 8
	users2 := []User{u1, u2, u3, u4}
	for i, _ := range users2 {
		users[i].AllBirthday()
	}

	//задача 9
	for _, user := range users2 {
		if user.IsAdult1() {
			fmt.Println(user.Card())
			continue
		}
	}

	//задача 10
	fmt.Println("До:", u1.Age)
	u1.WrongBirthday()
	fmt.Println("После WrongBirthday:", u1.Age)
	u1.RealBirthday()
	fmt.Println("После RealBirthday:", u1.Age)

	//задача 11
	u6 := User{
		Name: "Olya",
		Age:  25,
		Address: Address{
			City:   "Moscow",
			Street: "Tverskaya",
		},
	}
	fmt.Printf("%s живет в %s на улице %s\n", u6.Name, u6.Address.City, u6.Address.Street)

	//задача 12
	users3 := []User{u1, u2, u6}
	res := GroupByCity(users3)
	for city, count := range res {
		fmt.Println(city, "-", len(count))
	}

	//задача 13
	c := Company{
		Name: "Vasya",
		HQ: Address{
			City: "Moscow",
		},
	}
	c1 := Company{
		Name: "Masha",
		HQ: Address{
			City: "SPb",
		},
	}
	company := []Company{c, c1}
	for _, v := range company {
		if v.HQ.City == "SPb" {
			fmt.Println(v.Name)
		}
	}

	//задача 14
	fmt.Println(u1.Phone)

	//задача 15
	u1.ShowPhone()

	//задача 16
	con := Contact{Phone: "+7-900-000-00-00"}
	u1.Info()
	con.Info()

	//задача 17
	bAcc := BankAccount{balance: 1000}
	bAcc.GetBalance()
	err := bAcc.SetBalance(-200)
	if err != nil {
		fmt.Println(err)
	}

	//задача 18
	fmt.Println("Было:", bAcc.balance)
	bAcc.Deposit(500)
	bAcc.Withdraw(200)
	fmt.Println("История:", bAcc.transactionHistory)
	fmt.Println("Стало:", bAcc.balance)

	//задача 19
	bAcc1 := BankAccount{balance: 500}
	bAcc2 := BankAccount{balance: 1000}
	remittance := map[string]*BankAccount{
		"Vasya": &bAcc1,
		"Masha": &bAcc2,
	}
	Transfer(remittance, "Vasya", "Masha", 300)

	//задача 20
	emp := Employee{
		Name:   "Vasya",
		Age:    25,
		Salary: 900,
	}
	emp1 := Employee{
		Name:   "Masha",
		Age:    31,
		Salary: 1200,
	}
	emp2 := Employee{
		Name:   "Petr",
		Age:    28,
		Salary: 1000,
	}
	emps := []Employee{emp, emp1, emp2}

	for _, employee := range emps {
		if employee.Age < 30 {
			employee.Salary = int(float64(employee.Salary) * 1.1)
		}
		fmt.Println(employee.Name, "-", employee.Salary)
	}

	//задача 21
	products := []Product{
		{Name: "Телефон",
			Price: 1200},
		{Name: "Чехол",
			Price: 300},
		{Name: "Кабель",
			Price: 200},
	}
	allSum := 0
	for _, product := range products {
		allSum += product.Price
	}
	if allSum > 1000 {
		allSum = int(float64(allSum) * 0.9)
	}
	fmt.Println("Итоговая сумма:", allSum)

	//задача 22
	std := []Student{
		{Name: "Vasya",
			Grades: []int{5, 4, 5}},
		{Name: "Masha",
			Grades: []int{4, 3, 4}},
		{Name: "Olya",
			Grades: []int{5, 5, 5}},
	}
	bestName := ""
	bestAvg := 0.0

	for _, student := range std {
		sum := 0
		count := 0
		for _, grade := range student.Grades {
			sum = sum + grade
			count++
		}
		avg := float64(sum) / float64(count)
		if bestAvg < avg {
			bestAvg = avg
			bestName = student.Name
		}
	}
	fmt.Printf("Лучший студент: %s (%.2f)\n", bestName, bestAvg)

	//задача 23 скип так как была уже супер похожая задача

	//задача 24
	users4 := []User{
		{Name: "Vasya",
			Age: 25,
			Address: Address{
				City: "Moscow"}},
		{Name: "Masha",
			Age: 20,
			Address: Address{
				City: "Kazan"}},
		{Name: "Petr",
			Age: 35,
			Address: Address{
				City: "Moscow"}},
		{Name: "Olya",
			Age: 25,
			Address: Address{
				City: "Kazan"}},
	}
	result := make(map[string][]int)
	avgAge := 0.0
	for _, user := range users4 {
		result[user.Address.City] = append(result[user.Address.City], user.Age)
	}
	for city, age := range result {
		sumAge := 0
		count := 0
		for _, age1 := range age {
			sumAge = sumAge + age1
			count++
		}
		avgAge = float64(sumAge) / float64(count)
		fmt.Println(city, "-", avgAge)
	}

	//задача 25
	accOrders := []User{
		{Name: "Vasya",
			Orders: []Order{
				{"Телефон", 500, 1},
				{"Чехол", 200, 2},
			},
		},
		{Name: "Masha",
			Orders: []Order{
				{"Книга", 300, 3}},
		},
	}
	total := 0
	for _, order := range accOrders {
		sumUser := 0
		for _, order := range order.Orders {
			sumUser += order.Price * order.Count
		}
		fmt.Println(order.Name, ":", sumUser)
		total += sumUser
	}
	fmt.Println("Общая сумма заказов:", total)

	//задача 26
	visits := []Visit{
		{Name: "Vasya",
			Count: 3},
		{Name: "Masha",
			Count: 1},
		{Name: "Vasya",
			Count: 2},
	}
	result1 := make(map[string]int)
	for _, v := range visits {
		result1[v.Name] += v.Count
	}
	for name, count := range result1 {
		fmt.Printf("%s - %d\n", name, count)
	}
}

// задача 3
func FindUser(users []User, name string) *User {
	for i, _ := range users {
		if users[i].Name == name {
			return &users[i]
		}
	}
	return nil
}
func printUser(users []User, name string) {
	u := FindUser(users, name)
	if u != nil {
		fmt.Println("Найден:", u.Name)
	} else {
		fmt.Println("Не найден")
	}
}

// задача 4
func (u User) Hello() {
	fmt.Println("Привет,", u.Name)
}

// задача 5
func (u User) IsAdult() bool {
	if u.Age > 18 {
		fmt.Printf("%s - совершеннолетний\n", u.Name)
		return true
	}
	fmt.Printf("%s - несовершеннолетний\n", u.Name)
	return false
}

// задача 6
func (u *User) Birthday() {
	u.Age++
	fmt.Println("Возраст после дня рождения:", u.Age)
}

// задача 7
func (u User) Card() string {
	return fmt.Sprintf("%s (%d) <%s>", u.Name, u.Age, u.Email)
}

// задача 8
func (u *User) AllBirthday() {
	u.Age++
	fmt.Println(u.Name, "-", u.Age)
}

// задача 9
func (u User) IsAdult1() bool {
	if u.Age > 18 {
		return true
	}
	return false
}

// задача 10
func (u User) WrongBirthday() {
	u.Age++
}
func (u *User) RealBirthday() {
	u.Age++
}

// задача 12
func GroupByCity(users []User) map[string][]User {
	result := make(map[string][]User)
	for _, user := range users {
		result[user.Address.City] = append(result[user.Address.City], user)
	}
	return result
}

// задача 15
func (c Contact) ShowPhone() {
	fmt.Println("Телефон:", c.Phone)
}

// задача 16
func (c User) Info() {
	fmt.Println("User.Info")
}
func (c Contact) Info() {
	fmt.Println("Contact.Info")
}

// задача 17
func (a *BankAccount) SetBalance(amount int) error {
	if amount < 0 {
		return errors.New("ERR: negative amount")
	}
	a.balance += amount
	fmt.Println("New balance:", a.balance)
	return nil
}
func (a BankAccount) GetBalance() int {
	fmt.Println("OK:", a.balance)
	return a.balance
}

// задача 18
func (a *BankAccount) Deposit(amount int) {
	a.balance += amount
	a.transactionHistory = append(a.transactionHistory, amount)
}
func (a *BankAccount) Withdraw(amount int) {
	a.balance -= amount
	a.transactionHistory = append(a.transactionHistory, -amount)
}

// задача 19
func Transfer(reg map[string]*BankAccount, from, to string, amount int) {
	for name, account := range reg {
		if from == name {
			if account.balance > amount {
				account.Withdraw(amount)
				fmt.Println(name, ":", account.balance)
			}
		}
		if to == name {
			account.Deposit(amount)
			fmt.Println(name, ":", account.balance)
		}
	}
}
