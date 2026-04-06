package Lesson2

import (
	"fmt"
)

type Worker interface {
	Do()
}

type Printer interface {
	Print()
}

type Pricer interface {
	Price() int
}

type Updater interface {
	UpdateAge()
}

type Logger interface {
	Log(msg string)
}

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type LogWriter interface {
	Logger
	Writer
}

type ReadWriter interface {
	Reader
	Writer
}

type Sender interface {
	Send(msg string)
}

type Formater interface {
	Format(msg string) string
}

type Notifier interface {
	Sender
	Formater
}

type Payment interface {
	Pay(amount int)
}

type Storage interface {
	Save(key string, value interface{})
	Get(key string) interface{}
}

type Server struct {
	Name   string
	Logger Logger
}
type Dev struct{}
type Tester struct{}
type TaskManager struct {
	Worker Worker
}
type ConsoleWriter struct{}
type EmailNotifier struct{}
type ConsoleLogger struct{}
type ActivityLogger struct{}
type FileLogger struct {
	History []string
}
type SilentLogger struct{}
type App struct {
	Name   string
	Logger Logger
}
type Note struct {
	text string
}
type User struct {
	Name   string
	Age    int
	Logger Logger
}
type Product struct {
	Name      string
	UnitPrice int
}
type Service struct {
	Name           string
	Minutes        int
	RatePerMinutes int
}
type CardPayment struct{}
type CashPayment struct{}
type MemoryStorage struct {
	data map[string]interface{}
}

func (ms *MemoryStorage) Save(key string, value interface{}) {
	ms.data[key] = value
}
func (ms *MemoryStorage) Get(key string) interface{} {
	return ms.data[key]
}

func (tm TaskManager) RunAll(workers []Worker) {
	for _, w := range workers {
		w.Do()
	}
}

func (dev Dev) Do() {
	fmt.Println("Пишу код")
}
func (tester Tester) Do() {
	fmt.Println("Тестирую")
}

func (al ActivityLogger) Log(msg string) {
	fmt.Println("[LOG]:", msg)
}

func (s Server) Log(msg string) {
	fmt.Println(msg)
}
func (s Server) Start() {
	s.Logger.Log(s.Name + " started")
}
func (s Server) Stop() {
	s.Logger.Log(s.Name + " stopped")
}

func (cw ConsoleWriter) Write(msg string) {
	fmt.Println("Write:", msg)
}
func (cw ConsoleWriter) Log(msg string) {
	fmt.Println("Log:", msg)
}

func (e EmailNotifier) Send(msg string) {
	fmt.Println(msg)
}
func (e EmailNotifier) Format(msg string) string {
	return "EMAIL: " + msg
}

func (ConsoleLogger) Log(msg string) {
	fmt.Println("[LOG]:", msg)
}

func (f *FileLogger) Log(msg string) {
	f.History = append(f.History, msg)
}
func Run(log Logger) {
	log.Log("Started")
}

func (SilentLogger) Log(msg string) {
}

func (a App) Start() {
	a.Logger.Log("App " + a.Name + " started")
}

func (n *Note) Read() string {
	return n.text
}
func (n *Note) Write(data string) {
	n.text = data
}

func (u User) Do(action string) {
	u.Logger.Log(u.Name + " сделал " + action)
}
func (u User) Print() {
	fmt.Println(u.Name, u.Age)
}

//	func (u User) UpdateAge() {
//		u.Age++
//	}
func (u *User) UpdateAge() {
	u.Age++
}

func (p Product) Price() int {
	return p.UnitPrice
}

func (s Service) Price() int {
	return s.RatePerMinutes * s.Minutes
}

func (cp CashPayment) Pay(amount int) {
	fmt.Println("Оплата наличными:", amount)
}

func (cp CardPayment) Pay(amount int) {
	fmt.Println("Оплата картой:", amount)
}

func Process(p Payment) {
	p.Pay(100)
}

func Main() {
	//задача 1
	users := []Printer{
		User{Name: "Vasya",
			Age: 19},
		User{Name: "Masha",
			Age: 21},
		User{Name: "Petr",
			Age: 30},
	}
	for _, user := range users {
		user.Print()
	}

	//задача 2
	products := []Pricer{
		Product{"Phone", 1200},
		Product{"Case", 200},
		Service{
			Name:           "Repair",
			Minutes:        30,
			RatePerMinutes: 10,
		},
	}
	allSum := 0
	for _, product := range products {
		allSum += product.Price()
	}
	fmt.Println("Итого:", allSum)

	//задача 3
	var up Updater
	u := User{
		Age: 25,
	}
	fmt.Println(u)
	up = &u
	up.UpdateAge()
	fmt.Println(u)

	//задача 4
	items := []interface{}{"Go", 10, true, 3.14}
	for _, item := range items {
		fmt.Printf("%v - %T\n", item, item)
	}

	//задача 5
	items2 := []interface{}{"hi", 2, "Go", 4, true, 7}
	intCount1 := 0
	stringCount1 := 0
	boolCount1 := 0
	for _, i2 := range items2 {
		switch i2.(type) {
		case int:
			intCount1++
		case string:
			stringCount1++
		case bool:
			boolCount1++
		}
	}
	fmt.Printf("ints: %d\nstrings: %d\nbools: %d\n", intCount1, stringCount1, boolCount1)

	//задача 6
	items3 := []interface{}{"go", 42, false}
	PrintAll(items3)

	//задача 7
	m := map[string]interface{}{
		"Name":   "Vasya",
		"Age":    25,
		"active": true,
	}
	for k, value := range m {
		fmt.Printf("%s: %v (%T)\n", k, value, value)
	}

	//задача 8
	var a interface{}
	var b interface{} = (*int)(nil)
	fmt.Println(a == nil)
	fmt.Println(b == nil)

	//задача 9
	var x interface{} = 42
	if value, ok := x.(int); ok {
		fmt.Println("ok int:", value)
	} else {
		fmt.Println("not int")
	}
	if value, ok := x.(string); ok {
		fmt.Println("ok string:", value)
	} else {
		fmt.Println("not string")
	}

	//задача 10
	DetectType(10)
	DetectType("hi")
	DetectType(true)
	DetectType(3.14)

	//задача 11
	Add(3, 4)
	Add(2.5, 3.5)
	Add(3, "hi")

	//задача 12
	items4 := []interface{}{1, "go", true, 2, "hi", 3}
	intCount := 0
	stringCount := 0
	boolCount := 0
	for _, i2 := range items4 {
		switch i2.(type) {
		case int:
			intCount++
		case string:
			stringCount++
		case bool:
			boolCount++
		}
	}
	fmt.Printf("int: %d\nstring: %d\nbool: %d\n", intCount, stringCount, boolCount)

	//задача 13
	PrintAnySlice(items)

	//задача 14
	items5 := []interface{}{"hi", 5, true, "go", 8, false}
	str := []string{}
	ints := []int{}
	bl := []bool{}
	for _, item := range items5 {
		switch v := item.(type) {
		case int:
			ints = append(ints, v)
		case string:
			str = append(str, v)
		case bool:
			bl = append(bl, v)
		default:
			fmt.Println("неизвестный тип")
		}
	}
	fmt.Println("ints:", ints)
	fmt.Println("strings:", str)
	fmt.Println("bools:", bl)

	//задача 15
	s := "Hello"
	rw := &Note{}
	rw.Write(s)
	fmt.Println(rw.Read())

	//задача 16
	app := App{
		Name:   "GoApp",
		Logger: ConsoleLogger{},
	}
	app.Start()
	fmt.Println()

	//задача 17
	app.Start()
	app.Logger = SilentLogger{}
	app.Start()

	//задача 18
	word := "Hello"
	emailN := EmailNotifier{}
	emailN.Send(emailN.Format(word))

	//задача 19
	lw := ConsoleWriter{}
	lw.Write("Data")
	lw.Log("Event")

	//задача 20
	server := Server{
		Name:   "Main",
		Logger: Server{},
	}
	server.Start()
	server.Stop()

	//задача 21
	Run(ConsoleLogger{})
	fl := &FileLogger{}
	Run(fl)
	fmt.Println(fl.History)

	//задача 22
	Process(CardPayment{})
	Process(CashPayment{})

	//задача 23
	ms := MemoryStorage{
		data: make(map[string]interface{}),
	}
	ms.Save("user", "Vasya")
	ms.Save("age", 25)

	for _, key := range []string{"user", "age"} {
		value := ms.Get(key)

		switch v := value.(type) {
		case string:
			fmt.Printf("%s (string): %v\n", key, v)
		case int:
			fmt.Printf("%s (int): %v\n", key, v)
		}
	}

	//задача 24
	u1 := User{
		Name: "Task",
	}
	PrintInfo("Go")
	PrintInfo(42)
	PrintInfo(u1)
	PrintInfo(true)

	//задача 25
	u2 := User{
		Name:   "Vasya",
		Logger: ActivityLogger{},
	}
	u3 := User{
		Name:   "Masha",
		Logger: ActivityLogger{},
	}
	u2.Do("login")
	u3.Do("logout")

	//задача 26
	fmt.Println(ConvertToString("Go"))
	fmt.Println(ConvertToString(42))
	fmt.Println(ConvertToString(true))
	fmt.Println(ConvertToString(nil))
	fmt.Println(ConvertToString(3.1))

	//задача 27
	tm := TaskManager{}
	d := Dev{}
	t := Tester{}
	tm.RunAll([]Worker{d, t})

	//задача 28
	fmt.Println("Количество чисел:", CountInts(items4))

	//задача 29
	items6 := []interface{}{"hi", 5, true, "go", 8, []int{1, 2, 3}}
	Process1(items6)
}

// задача 6
func PrintAll(values []interface{}) {
	for i, value := range values {
		fmt.Printf("[%d] %v (тип %T)\n", i, value, value)
	}
}

// задача 10
func DetectType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Число")
	case string:
		fmt.Println("Строка")
	case bool:
		fmt.Println("Булево")
	default:
		fmt.Println("неизвестный тип")
	}
}

// задача 11
func Add(a, b interface{}) {
	switch a := a.(type) {
	case int:
		switch b := b.(type) {
		case int:
			fmt.Println(a + b)
		default:
			fmt.Println("Невозможно сложить")
		}
	case float64:
		switch b := b.(type) {
		case float64:
			fmt.Println(a + b)
		default:
			fmt.Println("Невозможно сложить")
		}
	default:
		fmt.Println("Невозможно сложить")
	}
}

// задача 13
func PrintAnySlice(items []interface{}) {
	for _, item := range items {
		switch v := item.(type) {
		case int:
			fmt.Println("Число:", v)
		case string:
			fmt.Println("Строка:", v)
		case bool:
			fmt.Println("Булево:", v)
		default:
			fmt.Println("неизвестный тип")
		}
	}
}

// задача 24
func PrintInfo(v interface{}) {
	switch value := v.(type) {
	case int:
		fmt.Println("Число:", value)
	case string:
		fmt.Println("Строка:", value)
	case User:
		fmt.Println("Объект:", value.Name)
	default:
		fmt.Println("Неизветсный тип")
	}
}

// задача 26
func ConvertToString(x interface{}) string {
	switch value := x.(type) {
	case int:
		return fmt.Sprintf("%d", value)
	case string:
		return value
	case bool:
		return fmt.Sprintf("%t", value)
	case nil:
		return "nil"
	}
	return fmt.Sprintf("unknown")
}

// задача 28
func CountInts(values []interface{}) int {
	countInt := 0
	for _, item := range values {
		switch item.(type) {
		case int:
			countInt++
		}
	}
	return countInt
}

// задача 29
func Process1(items []interface{}) {
	for _, item := range items {
		switch v := item.(type) {
		case int:
			fmt.Println("Число:", v)
		case string:
			fmt.Println("Строка:", v)
		case bool:
			fmt.Println("Булево:", v)
		case []int:
			fmt.Println("Срез чисел:", v)
		default:
			fmt.Println("неизвестный тип")
		}
	}
}
