package Lesson1

import (
	"errors"
	"fmt"
	"strconv"
)

type FieldError struct {
	Field   string
	Message string
}

type AppError struct {
	Code    int
	Message string
}

type ValidationError struct {
	Field   string
	Message string
}

type SystemError struct {
	Message string
	Code    int
}

func (e FieldError) Error() string {
	return fmt.Sprintf("Ошибка: поле %q %s", e.Field, e.Message)
}

func (e AppError) Error() string {
	return fmt.Sprintf("код %d: %s", e.Code, e.Message)
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("поле %q %s", v.Field, v.Message)
}

func (err *SystemError) Error() string {
	return err.Message
}

var ErrInvalidLogin = errors.New("неверные данные входа")
var ErrToShort = errors.New("пароль слишком короткий")
var ErrNoDigit = errors.New("в пароле должна быть хотя бы одна цифра")
var ErrInvalidAge = errors.New("возраст вне диапозона")
var ErrConnection = errors.New("ошибка соединения")
var ErrAccessDenied = errors.New("доступ запрещён")
var ErrNotFound = errors.New("запись не найдена")

func Main() {
	// задача 1
	err := checkNumber(1)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Число корректное")
	}
	err1 := checkNumber(-1)
	if err1 != nil {
		fmt.Println("Ошибка:", err1)
	} else {
		fmt.Println("Число корректное")
	}

	// задача 2
	result, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println("Ошибка:", err2)
	} else {
		fmt.Println(result)
	}

	// задача 3
	result1, err3 := parseAge("25")
	if err3 != nil {
		fmt.Println("Ошибка:", err3)
	} else {
		fmt.Println("Возраст:", result1)
	}
	result2, err4 := parseAge("-5")
	if err4 != nil {
		fmt.Println("Ошибка:", err4)
	} else {
		fmt.Println("Возраст:", result2)
	}
	result3, err5 := parseAge("abc")
	if err5 != nil {
		fmt.Println("Ошибка:", err5)
	} else {
		fmt.Println("Возраст:", result3)
	}

	// задача 4
	err6 := login("abc", "4321")
	if errors.Is(err6, ErrInvalidLogin) {
		fmt.Println("Ошибка:", err6)
	} else {
		fmt.Println("Успешный вход")
	}

	// задача 5
	passwords := []string{"short", "password", "goodpass1"}

	for _, p := range passwords {
		err := validatePassword(p)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("\"%s\" — ok\n", p)
		}
	}

	// задача 6
	err7 := checkUser()
	if err7 != nil {
		fmt.Println(err7)
	}

	// задача 7
	err8 := validateField("")
	if err8 != nil {
		fmt.Println(err8)
	}

	// задача 8
	err9 := authorize("ytrewq", "1234")
	if err9 != nil {
		fmt.Println(err9)
	} else {
		fmt.Println("ок")
	}

	err10 := authorize("qwerty", "4321")
	if err10 != nil {
		fmt.Println(err10)
	} else {
		fmt.Println("ок")
	}

	// задача 9
	err11 := registerUser("")
	if err11 != nil {
		var ve *ValidationError
		if errors.As(err11, &ve) {
			fmt.Println("Ошибка в поле:", ve.Field)
		} else {
			fmt.Println("Другая ошибка:", err11)
		}
	}

	// задача 10
	err12 := processData()
	if err12 != nil {
		var ve *ValidationError
		if errors.As(err12, &ve) {
			fmt.Println("Ошибка:", err12)
		}
	}

	// задача 11
	modes := []string{"validation", "system"}

	for _, m := range modes {
		err := runOperation(m)
		if err != nil {
			var ve *ValidationError
			var se *SystemError

			if errors.As(err, &ve) {
				fmt.Println("Ошибка валидации:", ve)
			} else if errors.As(err, &se) {
				fmt.Println("Системная ошибка:", se)
			}
		}
	}

	// задача 12
	err13 := validateAge(-5)
	if err13 != nil {
		if errors.Is(err13, ErrInvalidAge) {
			fmt.Println("возраст некорректный")
		}
	}

	// задача 13
	err14 := createUser()
	if err14 != nil {
		var ve *ValidationError
		if errors.As(err14, &ve) {
			fmt.Println("Ошибка в поле:", ve.Field)
		}
	}

	// задача 14
	err15 := connect()
	fmt.Println(errors.Unwrap(err15))

	// задача 15
	err16 := runOperation1("access")
	if errors.Is(err16, ErrAccessDenied) {
		fmt.Println("Ошибка:", ErrAccessDenied)
	}

	err17 := runOperation1("system")
	if err17 != nil {
		var se *SystemError
		if errors.As(err17, &se) {
			fmt.Println("Ошибка: системная, код", se.Code)
		}
	}

	// задача 16
	err18 := initApp()
	if errors.Is(err18, ErrNotFound) {
		fmt.Println("данные не найдены")
	}
}

// задача 1
func checkNumber(n int) error {
	if n < 0 {
		return errors.New("число не может быть отрицательным")
	}
	return nil
}

// задача 2
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("деление на ноль: a=%d, b=%d", a, b)
	}
	return a / b, nil
}

// задача 3
func parseAge(input string) (int, error) {
	age, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("возраст должен быть числом: %w", err)
	}
	if age < 0 || age > 120 {
		return 0, errors.New("возраст вне допустимого диапозона [0...120]")
	}
	return age, nil
}

// задача 4
func login(user, pass string) error {
	if user != "qwer" || pass != "1234" {
		return ErrInvalidLogin
	}
	return nil
}

// задача 5
func validatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("\"%s\" - ошибка: %w", password, ErrToShort)
	}
	hasDigits := false
	for _, char := range password {
		if char >= '0' && char <= '9' {
			hasDigits = true
			break
		}
	}
	if !hasDigits {
		return fmt.Errorf("\"%s\" - ошибка: %w", password, ErrNoDigit)
	}
	return nil
}

// задача 6
func checkName(name string) error {
	if name == "" {
		return errors.New("имя не может быть пустым")
	}
	return nil
}
func checkUser() error {
	err := checkName("")
	if err != nil {
		return fmt.Errorf("не удалось создать пользователя: %w", err)
	}
	return nil
}

// задача 7
func validateField(value string) error {
	if value == "" {
		return FieldError{Field: "username", Message: "не может быть пустым"}
	}
	return nil
}

// задача 8
func authorize(user, pass string) error {
	if user != "qwerty" {
		return AppError{Code: 404, Message: "пользователь не найден"}
	}
	if pass != "1234" {
		return AppError{Code: 401, Message: "неверный пароль"}
	}
	return nil
}

// задача 9
func registerUser(name string) error {
	if name == "" {
		return &ValidationError{Field: "name", Message: "имя обязательно"}
	}
	return nil
}

// задача 10
func checkData() error {
	return &ValidationError{
		Field:   "id",
		Message: "некорректное значение",
	}
}
func processData() error {
	err := checkData()
	if err != nil {
		return fmt.Errorf("обработка данных: %w", err)
	}
	return nil
}

// задача 11
func runOperation(mode string) error {
	switch mode {
	case "validation":
		return &ValidationError{
			Field:   "email",
			Message: "некорректно",
		}
	case "system":
		return &SystemError{
			Message: "Диск недоступен",
		}
	}
	return nil
}

// задача 12
func validateAge(age int) error {
	if age < 0 {
		return fmt.Errorf("проверка возраста: %w", ErrInvalidAge)
	}
	return nil
}

// задача 13
func checkUser1(name string) error {
	if name == "" {
		return &ValidationError{Field: "name", Message: "имя обязательно"}
	}
	return nil
}
func createUser() error {
	err := checkUser1("")
	if err != nil {
		return fmt.Errorf("ошибка создания пользователя: %w", err)
	}
	return nil
}

// задача 14
func connect() error {
	return fmt.Errorf("не удалось подключиться: %w", ErrConnection)
}

// задача 15
func runOperation1(mode string) error {
	switch mode {
	case "access":
		return fmt.Errorf("операция: %w", ErrAccessDenied)
	case "system":
		return &SystemError{Code: 500}
	}
	return nil
}

// задача 16
func readData() error {
	return ErrNotFound
}
func loadResource() error {
	err := readData()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
func initApp() error {
	err := loadResource()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
