package Lesson6

import (
	"TasksFromTheCourse/Module2/Lesson5"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Main() {
	//задача 1
	phoneBook := make(map[string]int)
	phoneBook["Masha"] = 7435354674
	phoneBook["Petya"] = 4323423561
	phoneBook["Vasya"] = 1234123567
	fmt.Println("Количество контактов:", len(phoneBook))
	fmt.Println(phoneBook)

	//задача 2
	exRange := map[string]float64{
		"USD": 97.5,
		"EUR": 104.3,
		"CNY": 13.1,
	}
	exRange["BYN"] = 29.8
	fmt.Println(exRange)

	//задача 3
	city := map[string]int{
		"Moscow": 13000000,
		"SPb":    5600000,
		"Kazan":  1400000,
		"Sochi":  480000,
	}
	fmt.Println("Общее количество городов:", len(city))
	fmt.Println("Население Moscow:", city["Moscow"], "\nНаселение Kazan:", city["Kazan"])

	//задача 4
	product := map[string]int{
		"apple":      120,
		"orange":     150,
		"banana":     90,
		"watermelon": 380,
		"melon":      250,
	}
	checkProduct("apple", product)
	checkProduct("lemon", product)

	//задача 5
	employees := map[string]int{
		"Masha": 800,
		"Petya": 1200,
		"Vasya": 900,
	}
	raiseSalaries(employees)
	fmt.Println(employees)

	//задача 6
	product1 := map[string]int{
		"apple":  120,
		"orange": 150,
		"banana": 90,
		"pear":   140,
	}
	deleteItem(product1, "banana")
	fmt.Println(product1)

	//задача 7
	employees1 := map[string]int{
		"Masha": 1500,
		"Petya": 1000,
		"Vasya": 1200,
	}
	updateSalary(employees1, "Vasya")
	fmt.Println("Vasya-", employees1["Vasya"])

	//задача 8
	product2 := map[string]int{
		"apple":  120,
		"orange": 230,
		"banana": 150,
		"pear":   140,
	}
	sumPrices(product2)

	//задача 9
	product3 := map[string]int{
		"apple":  120,
		"orange": 230,
		"banana": 90,
		"pear":   140,
	}
	name, price := maxItem(product3)
	fmt.Printf("Самый дорогой товар — %s (%d)\n", name, price)

	//задача 10
	product4 := map[string]int{
		"apple":  120,
		"orange": 150,
		"banana": 90,
		"pear":   140,
	}
	result := filterAbove(product4, 100)
	for _, names := range result {
		fmt.Println(names)
	}

	//задача 11
	students := map[string]map[string]int{
		"Vasya": {"math": 5, "phys": 4, "rus": 5},
		"Petya": {"math": 3, "phys": 4, "rus": 4},
	}
	avgGrades := avgGrades(students)
	for name, avg := range avgGrades {
		fmt.Printf("%s — %.2f\n", name, avg)
	}

	//задача 12
	words := []string{"cat", "book", "sun", "milk", "go", "tea"}
	result2 := groupByLength(words)
	for len, words := range result2 {
		fmt.Printf("%d:%v\n", len, words)
	}

	//задача 13
	foodStuff := map[string]map[string]int{
		"Fruits": {"apple": 120, "banana": 90, "orange": 150},
		"Veg":    {"tomato": 110, "cucumber": 90},
	}
	result3 := sumByCategory(foodStuff)
	for title, sum := range result3 {
		fmt.Printf("%s — %d\n", title, sum)
	}

	//задача 14
	text := "go go is fun go is great"
	freq := wordFreq(text)
	for word, count := range freq {
		fmt.Printf("%s:%d\n", word, count)
	}

	//задача 15
	words1 := topWords(freq, 3)
	for _, word := range words1 {
		fmt.Printf("%s - %d\n", word, freq[word])
	}

	//задача 16
	emp := map[string]string{
		"Anya":  "Marketing",
		"Petya": "Dev",
		"Vasya": "Dev",
	}
	result4 := invertMap(emp)
	for dep, emp := range result4 {
		fmt.Printf("%s:%s\n", dep, emp)
	}

	//задача 17
	m := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	fmt.Println(m)
	fmt.Println(copyWithout(m, "B"))

	//задача 18
	m1 := map[string]int{
		"A": 1,
		"B": 2,
	}
	m2 := map[string]int{
		"B": 5,
		"C": 7,
	}
	fmt.Println(mergeMaps(m1, m2))

	//задача 19
	stdScores := map[string][]int{
		"Vasya": {5, 5, 4},
		"Petya": {3, 4, 3},
	}
	avgScores := avgScores(stdScores)
	for name, avgS := range avgScores {
		fmt.Printf("%s — %.2f\n", name, avgS)
	}

	//задача 20
	age := 25
	age1 := 20
	usersAge := map[string]*int{
		"Vasya": &age,
		"Masha": &age1,
	}
	increaseAge(usersAge, "Vasya")

	//задача 21
	prc := 120
	prc1 := 150
	goods := map[string]*int{
		"apple":  &prc,
		"orange": &prc1,
	}
	if productName, ok := goods["orange"]; ok {
		applyDiscount(productName, 20)
	}
	for productName, price := range goods {
		fmt.Printf("%s - %d\n", productName, *price)
	}

	//задача 22
	nums := []int{1, 2, 2, 3, 3, 4, 1}
	newNums := unique(nums)
	fmt.Println(newNums)

	//задача 23
	nums1 := []int{1, 2, 3, 5}
	nums2 := []int{2, 3, 4, 5}
	fmt.Println(intersect(nums1, nums2))

	//задача 24
	std1 := map[string]float64{
		"Vasya": 4.8,
		"Petya": 3.9,
		"Masha": 4.5,
	}
	stdF := filterAbove1(std1, 4.5)
	for _, name := range stdF {
		fmt.Println(name)
	}

	//задача 25
	std2 := map[string]float64{
		"Vasya": 4.0,
		"Masha": 5.0,
	}
	std3 := map[string]float64{
		"Vasya": 5.0,
		"Petya": 3.5,
	}
	avgS2 := mergeAverage(std2, std3)
	for name, score := range avgS2 {
		fmt.Printf("%s - %.1f\n", name, score)
	}

	//задача 26
	listSales := [][]string{
		{"Vasya", "apple", "100"},
		{"Vasya", "orange", "200"},
		{"Masha", "apple", "300"},
		{"Vasya", "pear", "300"},
		{"Masha", "pear", "100"},
	}
	salesReport(listSales)
}

// задача 4
func checkProduct(name string, product map[string]int) (string, bool) {
	if price, ok := product[name]; ok {
		fmt.Printf("Цена товара %s: %d\n", name, price)
	} else {
		fmt.Println("Нет в наличии")
	}
	return name, false
}

// задача 5
func raiseSalaries(employees map[string]int) {
	for employee, price := range employees {
		if price < 1000 {
			employees[employee] = int(float64(price) * 1.1)
		}
	}
}

// задача 6
func deleteItem(product1 map[string]int, name string) {
	if _, ok := product1[name]; ok {
		delete(product1, name)
	}
}

// задача 7
func updateSalary(employees1 map[string]int, name string) {
	if price, ok := employees1[name]; ok {
		employees1[name] = int(float64(price) * 1.15)
	} else {
		fmt.Println("Нет сотрудника:", name)
	}
}

// задача 8
func sumPrices(product2 map[string]int) int {
	result := 0
	for _, price := range product2 {
		result += price
	}
	fmt.Printf("Общая сумма: %d\n", result)
	return result
}

// задача 9
func maxItem(product3 map[string]int) (string, int) {
	maxName := ""
	maxPrice := 0
	for name, price := range product3 {
		if maxPrice < price {
			maxPrice = price
			maxName = name
		}
	}
	return maxName, maxPrice
}

// задача 10
func filterAbove(product4 map[string]int, limit int) []string {
	result := make([]string, 0)
	for name, price := range product4 {
		if price > limit {
			result = append(result, name)
		}
	}
	return result
}

// задача 11
func avgGrades(students map[string]map[string]int) map[string]float64 {
	result := make(map[string]float64)
	for name, subjects := range students {
		sum := 0
		count := 0
		for _, score := range subjects {
			sum += score
			count++
		}
		result[name] = float64(sum) / float64(count)
	}
	return result
}

// задача 12
func groupByLength(words []string) map[int][]string {
	result := make(map[int][]string)
	for _, word := range words {
		result[len(word)] = append(result[len(word)], word)
	}
	return result
}

// задача 13
func sumByCategory(foodStaff map[string]map[string]int) map[string]int {
	result := make(map[string]int)
	for title, s := range foodStaff {
		sum := 0
		for _, i := range s {
			sum += i
		}
		result[title] = sum
	}
	return result
}

// задача 14
func wordFreq(text string) map[string]int {
	result := make(map[string]int)
	words := strings.Split(text, " ")
	for _, word := range words {
		result[word]++
	}
	return result
}

// задача 15
func topWords(freq map[string]int, n int) []string {
	type Item struct {
		Word  string
		Count int
	}
	var result []Item

	for word, count := range freq {
		result = append(result, Item{
			Word:  word,
			Count: count})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Count > result[j].Count
	})

	if n > len(result) {
		n = len(result)
	}
	result = result[:n]

	words := make([]string, 0, n)
	for _, i := range result {
		words = append(words, i.Word)
	}
	return words
}

// задача 16
func invertMap(emp map[string]string) map[string][]string {
	result := make(map[string][]string)
	for emp, dep := range emp {
		result[dep] = append(result[dep], emp)
	}
	return result
}

// задача 17
func copyWithout(m map[string]int, drop string) map[string]int {
	result := make(map[string]int)
	for s, _ := range m {
		if s == drop {
			delete(m, s)
		}
	}
	result = m
	return result
}

// задача 18
func mergeMaps(m1, m2 map[string]int) map[string]int {
	for s, i := range m2 {
		m1[s] = i
	}
	return m1
}

// задача  19
func avgScores(stdScores map[string][]int) map[string]float64 {
	result := make(map[string]float64)
	for name, score := range stdScores {
		sum := 0
		count := 0
		for _, score := range score {
			sum += score
			count++
		}
		result[name] = float64(sum) / float64(count)
	}
	return result
}

// задача 20
func increaseAge(users map[string]*int, name string) {
	if _, ok := users[name]; ok {
	} else {
		fmt.Println("Нет такого пользователя:", name)
	}
	for names, age := range users {
		if names == name {
			*age++
		}
		fmt.Printf("%s - %d\n", names, *users[names])
	}
}

// задача 21
func applyDiscount(price *int, percent int) {
	*price = int(float64(*price) * (1 - float64(percent)/100))
}

// задача 22
func unique(nums []int) []int {
	result := []int{}
	for _, n := range nums {
		if !Lesson5.Contains(result, n) {
			result = append(result, n)
		}
	}
	return result
}

// задача 23
func intersect(a, b []int) []int {
	result := []int{}
	for _, nums1 := range a {
		for _, nums2 := range b {
			if nums1 == nums2 {
				result = append(result, nums1)
			}
		}
	}
	return result
}

// задача 24
func filterAbove1(m map[string]float64, threshold float64) []string {
	result := make([]string, 0)
	for name, score := range m {
		if score >= threshold {
			result = append(result, name)
		}
	}
	return result
}

// задача 25
func mergeAverage(std1, std2 map[string]float64) map[string]float64 {
	result := make(map[string]float64)
	for name1, score1 := range std1 {
		result[name1] = score1
	}
	for name2, score2 := range std2 {
		if _, ok := result[name2]; ok {
			result[name2] = (score2 + result[name2]) / 2
		} else {
			result[name2] = score2
		}
	}
	return result
}

// задача 26
func salesReport(sales [][]string) {
	stats := map[string][]int{}
	for _, sale := range sales {
		name := sale[0]
		priceStr := sale[2]
		price, _ := strconv.Atoi(priceStr)
		if _, ok := stats[name]; !ok {
			stats[name] = []int{0, 0}
		}
		stats[name][0]++
		stats[name][1] += price
	}
	for name, data := range stats {
		count := data[0]
		sum := data[1]
		avg := sum / count

		fmt.Printf("%s — %d продажи, сумма %d, средний чек %d\n",
			name, count, sum, avg)
	}
}
