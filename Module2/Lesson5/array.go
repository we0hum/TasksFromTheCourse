package Lesson5

import "fmt"

func Main() {
	//задача 1
	arr := [3]int{1, 2, 3}
	fmt.Println("Длина:", len(arr))
	fmt.Println(arr)

	//задача 2
	arr1 := [4]string{"a", "b", "c", "d"}
	arr1[1] = "x"
	fmt.Println(arr1)

	//задача 3
	arr2 := [5]int{}
	for i, _ := range arr2 {
		arr2[i] = 5 * (i + 1)
	}
	fmt.Println(arr2)
	result := arr2[0] + arr2[len(arr2)-1]
	fmt.Println(result)

	//задача 4
	arr3 := [6]int{8, 2, 7, 4, 9, 1}
	result1 := arr3[0]
	for _, n := range arr3 {
		if result1 >= n {
			result1 = n
		}
	}
	fmt.Println("Минимум:", result1)

	//задача 5
	arr4 := [8]int{1, 0, 5, 0, 7, 0, 9, 0}
	count := 0
	for i, _ := range arr4 {
		if arr4[i] == 0 {
			count++
		}
	}
	fmt.Println("Нулей:", count)

	//задача 6
	arr5 := [6]int{1, 2, 3, 4, 5, 6}
	s := arr5[1:4]
	fmt.Println(s)

	//задача 7
	arr6 := [8]int{10, 20, 30, 40, 50, 60, 70, 80}
	s1 := arr6[2:6]
	fmt.Println(s1, len(s1), cap(s1))

	//задача 8.1
	nums := make([]int, 5)
	for i, _ := range nums {
		nums[i] = i + 1
	}
	fmt.Println(nums)
	fmt.Println("Длина:", len(nums))
	fmt.Println("Емкость:", cap(nums))

	//задача 8.2
	arr7 := [5]int{100, 200, 300, 400, 500}
	s2 := arr7[1:3]
	s2[0] = 777
	fmt.Println("Массив:", arr7)
	fmt.Println("Срез:", s2)

	//задача 9
	arr8 := [4]int{10, 20, 30, 40}
	s3 := make([]int, len(arr8))
	copy(s3, arr8[:])
	s3[0] = 76
	fmt.Println(arr8)
	fmt.Println(s3)

	//задача 10
	x := make([]int, 3, 6)
	x[0], x[1], x[2] = 1, 2, 3
	y := append(x, 4)
	z := append(x, 5)
	fmt.Println(x, y, z)

	//задача 11
	a := []int{1, 2, 3}
	b := append(a, 4, 5, 6, 7)
	a[0] = 99
	fmt.Println(a)
	fmt.Println(b)

	//Практические задачи
	//задача 12
	array := [5]int{1, 2, 3, 4, 5}
	result2 := 0
	for _, i2 := range array {
		result2 += i2
	}
	fmt.Println(result2)

	//задача 13
	array2 := [6]int{15, 2, 30, 7, 9, 100}
	max, min := array2[0], array2[0]
	for _, n := range array2 {
		if max >= n {
		} else {
			max = n
		}
		if min <= n {
		} else {
			min = n
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)

	//задача 14
	array3 := [7]int{1, 2, 3, 4, 5, 6, 7}
	count2 := 0
	for i, _ := range array3 {
		if array3[i]%2 == 0 {
			count2++
		}
	}
	fmt.Println("Чётных:", count2)

	//задача 15
	array4 := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(array4)/2; i++ {
		j := len(array4) - 1 - i
		array4[i], array4[j] = array4[j], array4[i]
	}
	fmt.Println(array4)

	//задача 15.1 разворот без изменения оригинала
	array5 := []int{1, 2, 3, 4, 5}
	newSlise := make([]int, len(array5))

	for i := 0; i <= len(array5)/2; i++ {
		j := len(array5) - 1 - i
		newSlise[i], newSlise[j] = array5[j], array5[i]
	}

	fmt.Println("Исходный:", array5)
	fmt.Println("Новый:", newSlise)

	//задача 16
	array6 := []int{-5, 10, -3, 7, 0, 8}
	newSlise2 := []int{}
	for _, n := range array6 {
		if n > 0 {
			newSlise2 = append(newSlise2, n)
		}
	}
	fmt.Println(newSlise2)

	//задача 17
	array7 := []int{10, 20, 30, 40, 50}
	array7 = append(array7[:2], array7[3:]...)
	fmt.Println(array7)

	//задача 18
	array8 := []int{1, 2, 3}
	array9 := []int{4, 5, 6}
	array8 = append(array8, array9...)
	fmt.Println(array8)

	//задача 19
	array10 := []int{1, 2, 3}
	array11 := []int{}
	for _, n := range array10 {
		array11 = append(array11, n, n)
	}
	fmt.Println(array11)

	//гибридные задачи
	//задача 20
	ns := []int{1, 2, 2, 3, 3, 3, 4}
	ns2 := []int{}
	for _, n := range ns {
		if !Contains(ns2, n) {
			ns2 = append(ns2, n)
		}
	}
	fmt.Println(ns2)

	//задача 21
	word := []rune("level")
	isPalindrome := true

	for i := 0; i < len(word)/2; i++ {
		j := len(word) - 1 - i
		if word[i] == word[j] {
			continue
		} else {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		fmt.Println("Да")
	} else {
		fmt.Println("Нет")
	}

	//задача 22
	ns3 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	maxSum := 0
	var segment []int
	for i := 0; i < len(ns3)-2; i++ {
		tSum := ns3[i] + ns3[i+1] + ns3[i+2]
		if maxSum < tSum {
			maxSum = tSum
			segment = ns3[i : i+3]
		}
	}
	fmt.Println(segment, "сумма =", maxSum)

	////задача 23
	ns4 := []int{1, 2, 2, 3, 3, 3}
	unique := []int{}
	for _, n := range ns4 {
		if !Contains(unique, n) {
			unique = append(unique, n)
		}
	}
	for _, u := range unique {
		count3 := 0
		for _, n := range ns4 {
			if n == u {
				count3++
			}
		}
		fmt.Println(u, count3)
	}

	//задача 24
	nums1 := []int{5, 3, 8, 1, 2}
	for j := 0; j < len(nums1); j++ {
		for i := 0; i < len(nums1)-1-j; i++ {
			if nums1[i] > nums1[i+1] {
				nums1[i], nums1[i+1] = nums1[i+1], nums1[i]
			}
		}
	}
	fmt.Println(nums1)

	//задача 25
	var nums2 []int
	prevCap := cap(nums2)
	fmt.Println("Начальная емкость:", prevCap)
	for i := 1; i <= 100; i++ {
		nums2 = append(nums2, i)
		if prevCap < cap(nums2) {
			prevCap = cap(nums2)
			fmt.Println("Новая емкость:", prevCap)
		}
	}
}

func Contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
