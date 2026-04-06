package Lesson34

import (
	"fmt"
)

func Main() {
	welcome()

	bye()
	bye()
	bye()

	hello("Катя")
	hello("Максим")

	add(3, 7)

	fmt.Println(square(6))
	fmt.Println(square(9))

	fmt.Println(average(3, 5))

	q, r := divide(10, 4)
	fmt.Println("Частное:", q, "Остаток:", r)

	min, max := minmax(3, 7)
	fmt.Println("Минимум:", min, "Максимум:", max)

	n := 10
	inc(n)
	fmt.Println("После вызова:", n)

	k := 10
	incPtr(&k)
	fmt.Println("После функции:", k)

	x, y := 5, 10
	swap(&x, &y)
	fmt.Println("x:", x, "y:", y)

	j := 99
	reset(&j)
	fmt.Println("j:", j)

	fmt.Println(sumAll(1, 2, 3))
	fmt.Println(sumAll(10, 25, 30, 45))

	fmt.Println(minAll())
	fmt.Println(minAll(3, 7, 1, 8, 2))

	//задача 15
	calc := func(x int) int {
		var result int
		if x < 21 {
			result = x * x
		} else {
			result = x % 3
		}
		return result
	}
	fmt.Println(calc(8))
	fmt.Println(calc(26))

	//задача 16
	isEven := func(n int) (b bool) {
		if n%2 == 0 {
			b = true
		}
		return b
	}
	fmt.Println(isEven(4))
	fmt.Println(isEven(7))

	fmt.Println(sumTo(5))

	fmt.Println(power(2, 5))

	fmt.Println(fib(6))

	res, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Результат:", res)

	//Pro-задачи
	b := 1000
	updateBalance(&b, -200)
	fmt.Println("Баланс:", b)

	acc1, acc2 := 500, 300
	transfer(&acc1, &acc2, 200)
	fmt.Println("acc1:", acc1, "acc2:", acc2)

	d, dn, t := report(15)
	fmt.Println("Doing:", d, "Done:", dn, "Todo:", t)

	fmt.Println(calculator(5, 3, "+"))
	fmt.Println(calculator(5, 3, "-"))
	fmt.Println(calculator(5, 3, "*"))

	fmt.Println(validatePassword("GoLang123"))
	fmt.Println(validatePassword("golang"))
	fmt.Println(validatePassword("GOLANG2025"))

	sum, count := digitStats(12345)
	fmt.Println("Сумма:", sum, "Кол-во:", count)

	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(123))
}

// задача 1
func welcome() {
	fmt.Println("Добро пожаловать в Go!")
}

// задача 2
func bye() {
	fmt.Println("До встречи!")
}

// задача 3
func hello(name string) {
	fmt.Println("Привет,", name)
}

// задача 4
func add(a, b int) {
	fmt.Println(a + b)
}

// задача 5
func square(a int) int {
	return a * a
}

// задача 6
func average(a, b int) float64 {
	return float64(a+b) / 2
}

// задача 7
func divide(a, b int) (int, int) {
	return a / b, a % b
}

// задача 8
func minmax(a, b int) (int, int) {
	min := b
	max := a
	if a < b {
		min = a
		max = b
	}
	return min, max
}

// задача 9
func inc(x int) {
	x += 1
	fmt.Println("Внутри:", x)
}

// задача 10
func incPtr(x *int) {
	*x += 1
}

// задача 11
func swap(a, b *int) {
	*a, *b = *b, *a
}

// задача 12
func reset(x *int) {
	if *x > 50 {
		*x = 0
	}
}

// задача 13
func sumAll(nums ...int) int {
	result := 0
	for _, n := range nums {
		result += n
	}
	return result
}

// задача 14
func minAll(nums ...int) int {
	if len(nums) == 0 {
		fmt.Println("Список пуст")
		return 0
	}

	min := nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}

////задача 15
//calc := func(x int) int {
//	var result int
//	if x < 21 {
//		result = x * x
//	} else {
//		result = x % 3
//	}
//	return result
//}

////задача 16
//isEven := func(n int) (b bool) {
//	if n%2 == 0 {
//		b = true
//	}
//	return b
//}

// задача 17
func sumTo(n int) int {
	if n == 1 {
		return 1
	}
	return n + sumTo(n-1)
}

// задача 18
func power(base, exp int) int {
	if exp == 0 {
		return 1
	}
	return base * power(base, exp-1)
}

// задача 19
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// задача 20
func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("деление на ноль")
	}
	return a / b, nil
}

// Pro-задачи
// задача 21 нет ответа хотя код правильный
func updateBalance(balance *int, amount int) {
	*balance += amount
}

// задача 22 тоже что и в 21 задачи не отображается ответ
func transfer(from, to *int, amount int) {
	if *from >= amount {
		*from -= amount
		*to += amount
	} else {
		fmt.Println("Не хватает денег")
	}
}

// задача 23
func report(tasks int) (int, int, int) {
	doing := 0
	done := 0
	todo := 0
	for i := 1; i <= tasks; i++ {
		if i%3 == 0 {
			doing++
		} else if i%2 == 0 {
			done++
		} else {
			todo++
		}
	}
	return doing, done, todo
}

// задача 24
func calculator(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	default:
		return 0
	}
}

// задача 25
func validatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	var hasDigit, hasLower, hasUpper bool

	for _, ch := range password {
		if ch >= '0' && ch <= '9' {
			hasDigit = true
		} else if ch >= 'a' && ch <= 'z' {
			hasLower = true
		} else if ch >= 'A' && ch <= 'Z' {
			hasUpper = true
		}
		if hasDigit && hasLower && hasUpper {
			break
		}
	}
	return hasDigit && hasLower && hasUpper
}

// задача 26
func digitStats(n int) (int, int) {
	if n == 0 {
		return 0, 1
	}
	sum := 0
	count := 0
	for n > 0 {
		last := n % 10
		sum += last
		count++
		n /= 10
	}
	return sum, count
}

// задача 27
func isPalindrome(n int) bool {
	original := n
	rev := 0
	for n > 0 {
		last := n % 10
		rev = rev*10 + last
		n /= 10
	}
	return original == rev
}
