package Lesson2

import "fmt"

func Main() {
	//задача 1
	age := 16
	if age < 18 {
		fmt.Println("Несовершеннолетний")
	} else {
		fmt.Println("Совершеннолетний")
	}

	//задача 2
	fmt.Println()
	temp := 12
	if temp < 0 {
		fmt.Println("Мороз")
	} else if temp > 0 && temp <= 20 {
		fmt.Println("Прохладно")
	} else {
		fmt.Println("Тепло")
	}

	//задача 3
	fmt.Println()
	sum := 1200
	hasCoupon := true
	if sum > 1000 || hasCoupon {
		fmt.Println("Скидка доступна")
	} else {
		fmt.Println("Скидки нет")
	}

	//задача 4
	fmt.Println()
	input := "qwerty"
	real := "qwerty"
	if input == real {
		fmt.Println("Вход выполнен")
	} else {
		fmt.Println("Ошибка")
	}

	//задача 5
	fmt.Println()
	age1 := 15
	withAdult := true
	if age1 > 16 || withAdult {
		fmt.Println("Проход разрешён")
	} else {
		fmt.Println("Проход запрещён")
	}

	//задача 6
	fmt.Println()
	isRaining := false
	temp1 := 18
	if !isRaining && temp1 > 16 {
		fmt.Println("Идем гулять")
	} else {
		fmt.Println("Не идем гулять")
	}

	//задача 7
	fmt.Println()
	login := "user"
	password := "1234"
	if login == "user" && password == "1234" {
		fmt.Println("ОК")
	} else {
		fmt.Println("Ошибка")
	}

	//задача 8
	fmt.Println()
	day := 6
	switch day {
	case 1:
		fmt.Println("Понедельник")
	case 2:
		fmt.Println("Вторник")
	case 3:
		fmt.Println("Среда")
	case 4:
		fmt.Println("Четверг")
	case 5:
		fmt.Println("Пятница")
	case 6:
		fmt.Println("Суббота")
	case 7:
		fmt.Println("Воскресенье")
	default:
		fmt.Println("Неизвестно")
	}

	//задача 9
	fmt.Println()
	mark := 5
	switch mark {
	case 5:
		fmt.Println("Отлично")
	case 4:
		fmt.Println("Хорошо")
	case 3:
		fmt.Println("Удовлетворительно")
	case 2, 1, 0:
		fmt.Println("Неуд")
	default:
		fmt.Println("Ошибочная оценка")
	}

	//задача 10
	fmt.Println()
	month := 12
	switch month {
	case 1, 2, 12:
		fmt.Println("Зима")
	case 3, 4, 5:
		fmt.Println("Весна")
	case 6, 7, 8:
		fmt.Println("Лето")
	case 9, 10, 11:
		fmt.Println("Осень")
	default:
		fmt.Println("Ошибка")
	}

	//задача 11
	fmt.Println()
	sum1 := 0
	for i := 1; i <= 100; i++ {
		sum1 += i
	}
	fmt.Println(sum1)

	//задача 12
	fmt.Println()
	for i := 1; i <= 10; i++ {
		fmt.Printf("7 x %d = %d\n", i, 7*i)
	}

	//задача 13
	fmt.Println()
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	//задача 14
	fmt.Println()
	counter := 0
	for i := 1; i <= 30; i++ {
		if i%3 == 0 {
			counter += 1
		}
	}
	fmt.Println(counter)

	//задача 15
	fmt.Println()
	balance := 1000
	for i := 0; i < 10; i++ {
		balance -= 70
	}
	fmt.Println("Остаток:", balance)

	//задача 16
	fmt.Println()
	x := 3
	for x < 100 {
		fmt.Println(x)
		x *= 2
	}
	fmt.Println("Итог:", x)

	//задача 17
	fmt.Println()
	n := 101
	for n > 100 {
		if n%17 == 0 {
			fmt.Println("Нашли:", n)
			break
		} else {
			n++
		}
	}

	//задача 18
	//fmt.Println()
	//r := 12345
	//sum2 := 0
	//for r > 0 {
	//
	//}

	//задача 19
	fmt.Println()
	for i := 1; i <= 10; i++ {
		if i == 4 || i == 7 {
			continue
		}
		fmt.Println(i)
	}

	//задача 20
	fmt.Println()
	t := 753481
	for t > 0 {
		last := t % 10
		if last%2 == 0 {
			fmt.Println(last)
			break
		} else {
			t /= 10
		}
	}

	//задача 21
	fmt.Println()
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

	//задача 22
	fmt.Println()
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}

	//задача 23
	fmt.Println()
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Print(i+j, " ")
		}
		fmt.Println()
	}

	//задача 24
	fmt.Println()
	for i := 1; i <= 50; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	//задача 25
	fmt.Println()
	for i := 2; i <= 50; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	//задача 26
	fmt.Println()
	todo := 0
	doing := 0
	done := 0
	for i := 1; i <= 15; i++ {
		if i%3 == 0 {
			doing++
		} else if i%2 == 0 {
			done++
		} else {
			todo++
		}
	}
	fmt.Println("К выполнению:", todo)
	fmt.Println("В работе:", doing)
	fmt.Println("Готово:", done)
}
