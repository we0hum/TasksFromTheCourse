package Lesson2

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func Main() {
	//Создание и базовая работа с каналами
	// задача 1
	fmt.Println(getMessage())

	// задача 2
	fmt.Println(echo("ping"))

	// задача 3
	fmt.Println(runSteps())

	// задача 4
	fmt.Println(waitForSignal())

	// задача 5
	fmt.Println(pipeline())

	// задача 6
	fmt.Println(waitAll())

	// Перебор канала с помощью range
	// задача 1
	fmt.Println(readNumbers())

	// задача 2
	fmt.Println(messageStream())

	// задача 3
	fmt.Println(sumChannel())

	// задача 4
	fmt.Println(multiSender())

	// задача 5
	fmt.Println(processTasks())

	// задача 6
	fmt.Println(stopOnFive())

	// Буферизированные каналы и отличие от небуферизированных
	// задача 7
	fmt.Println(mailBox())

	// задача 8
	fmt.Println(overflowTest())

	// задача 9
	fmt.Println(messageQueue())

	// задача 10
	fmt.Println(sumBuffered([]int{2, 4, 6}))

	// задача 11
	fmt.Println(producerConsumer())

	// Закрытие каналов и сигналы завершения
	// задача 12
	fmt.Println(readAll())

	// задача 13
	fmt.Println(checkClosed())

	// задача 14
	fmt.Println(signalDone())

	// задача 15
	fmt.Println(multiSignal())

	// задача 16
	fmt.Println(safeClose())

	// задача 17
	fmt.Println(multiReceiver())

	// задача от ChatGPT Fan-in(слияние каналов)
	fmt.Println(fanIn())

	// Однонаправленные каналы (chan<- и <-chan)
	// задача 1
	ch := make(chan string, 3)

	go sendData(ch)
	result := receiveData(ch)

	fmt.Println(result)

	// задача 2
	ch1 := make(chan int, 3)

	go producer(ch1)

	fmt.Println(consumer(ch1))

	// Работа с select
	// задача 1
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func() {
		time.Sleep(10 * time.Millisecond)
		ch2 <- "Быстро"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch3 <- "Медленно!"
	}()

	select {
	case v := <-ch2:
		fmt.Println("Первым ответил ch2:", v)
	case v := <-ch3:
		fmt.Println("Первым ответил ch3:", v)
	}

	// задача 2
	ch4 := make(chan string)

	select {
	case <-ch4:
	default:
		fmt.Println("канал пуст")
	}

	// задача 3
	ch5 := make(chan string)
	ch6 := make(chan string)

	go func() {
		for i := 1; i <= 3; i++ {
			ch5 <- "A" + strconv.Itoa(i)
			time.Sleep(30 * time.Millisecond)
		}
	}()

	go func() {
		for i := 1; i <= 3; i++ {
			ch6 <- "B" + strconv.Itoa(i)
			time.Sleep(30 * time.Millisecond)
		}
	}()

	for i := 0; i < 6; i++ {
		select {
		case v := <-ch5:
			fmt.Println("Получено из ch5:", v)
		case v := <-ch6:
			fmt.Println("Получено из ch6:", v)
		}
	}

	// задача 4
	ch7 := make(chan string)
	ch8 := make(chan string)
	ch9 := make(chan string)

	go func() {
		time.Sleep(20 * time.Millisecond)
		ch7 <- "Готово"
		close(ch7)
	}()

	go func() {
		time.Sleep(30 * time.Millisecond)
		ch8 <- "Готово"
		close(ch8)
	}()

	go func() {
		time.Sleep(10 * time.Millisecond)
		ch9 <- "Готово"
		close(ch9)
	}()

	select {
	case <-ch7:
		fmt.Println("Первым завершился ch7")
	case <-ch8:
		fmt.Println("Первым завершился ch8")
	case <-ch9:
		fmt.Println("Первым завершился ch9")
	}

	// Pro-задачи
	// задача 1
	squares := make(chan int)
	cubes := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			squares <- i * i
		}
		close(squares)
	}()

	go func() {
		for i := 1; i <= 5; i++ {
			cubes <- i * i * i
		}
		close(cubes)
	}()

	for squares != nil || cubes != nil {
		select {
		case v, ok := <-squares:
			if ok {
				fmt.Println("square:", v)
			} else {
				squares = nil
			}

		case v, ok := <-cubes:
			if ok {
				fmt.Println("cube:", v)
			} else {
				cubes = nil
			}
		}
	}

	// задача 2
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, jobs, results)
		}(i)
	}

	for i := 1; i <= 5; i++ {
		jobs <- i
	}

	close(jobs)
	wg.Wait()
	close(results)

	var res []int
	for v := range results {
		res = append(res, v)
	}

	fmt.Println("Результаты:", res)

	// задача 3
	ch10 := make(chan int)
	ch11 := make(chan int)
	done := make(chan struct{})

	go generator(ch10)
	go doubler(ch10, ch11)
	go printer(ch11, done)
	<-done

	// задача 4
	doneCh := make(chan struct{})

	go longTask(doneCh)
	<-doneCh
	fmt.Println("Работа завершена")

	// задача 5
	ch12 := make(chan string)
	ch13 := make(chan string)
	merged := make(chan string)

	go func() {
		defer close(ch12)
		for i := 1; i <= 3; i++ {
			ch12 <- "A" + strconv.Itoa(i)
		}
	}()

	go func() {
		defer close(ch13)
		for i := 1; i <= 3; i++ {
			ch13 <- "B" + strconv.Itoa(i)
		}
	}()

	go func() {
		for ch12 != nil || ch13 != nil {
			select {
			case v, ok := <-ch12:
				if ok {
					merged <- v
				} else {
					ch12 = nil
				}

			case v, ok := <-ch13:
				if ok {
					merged <- v
				} else {
					ch13 = nil
				}
			}
		}
		close(merged)
	}()

	for v := range merged {
		fmt.Println(v)
	}
}

// Создание и базовая работа с каналами
// задача 1
func getMessage() string {
	ch := make(chan string)
	go func() {
		ch <- "Go сила!"
	}()
	return <-ch
}

// задача 2
func echo(msg string) string {
	ch := make(chan string)
	go func() {
		received := <-ch
		fmt.Println("Получено:", received)
		ch <- received
	}()
	ch <- msg
	return <-ch
}

// задача 3
func runSteps() []string {
	out := make(chan string)
	done := make(chan struct{})
	done1 := make(chan struct{})

	go func() {
		out <- "scan"
		done <- struct{}{}
	}()
	go func() {
		<-done
		out <- "process"
		done1 <- struct{}{}
	}()
	go func() {
		<-done1
		out <- "save"
	}()

	return []string{
		<-out,
		<-out,
		<-out,
	}
}

// задача 4
func waitForSignal() string {
	start := make(chan struct{})
	res := make(chan string)

	go func() {
		<-start
		res <- "Работа началась!"
	}()
	start <- struct{}{}

	return <-res
}

// задача 5
func pipeline() string {
	dataCh := make(chan string)
	resultCh := make(chan string)

	go func() {
		dataCh <- "data"
	}()

	go func() {
		v := <-dataCh
		resultCh <- "processed " + v
	}()

	return <-resultCh
}

// задача 6
func waitAll() string {
	doneCh := make(chan struct{})

	go func() {
		doneCh <- struct{}{}
	}()

	go func() {
		doneCh <- struct{}{}
	}()

	go func() {
		doneCh <- struct{}{}
	}()

	for i := 0; i < 3; i++ {
		<-doneCh
	}

	return "Все завершены"
}

// Перебор канала с помощью range
// задача 1
func readNumbers() []int {
	ch := make(chan int, 3)

	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	var res []int
	for val := range ch {
		res = append(res, val)
	}

	return res
}

// задача 2
func messageStream() []string {
	ch := make(chan string, 5)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- "msg" + strconv.Itoa(i)
		}
		close(ch)
	}()

	var res []string
	for v := range ch {
		res = append(res, v)
	}

	return res
}

// задача 3
func sumChannel() int {
	ch := make(chan int, 5)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	allSum := 0
	for v := range ch {
		allSum += v
	}

	return allSum
}

// задача 4
func multiSender() []int {
	ch := make(chan int, 6)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 3; i++ {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 4; i <= 6; i++ {
			ch <- i
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	var res []int
	for v := range ch {
		res = append(res, v)
	}

	return res
}

// задача 5
func processTasks() string {
	ch := make(chan string)

	go func() {
		for i := 1; i < 5; i++ {
			ch <- "task " + strconv.Itoa(i)
		}
		close(ch)
	}()

	for range ch {
	}

	return "all tasks processed"
}

// задача 6
func stopOnFive() []int {
	ch := make(chan int)
	doneCh := make(chan struct{})

	go func() {
		defer close(ch)
		for i := 1; i <= 10; i++ {
			select {
			case ch <- i:
			case <-doneCh:
				return
			}
		}
	}()

	var res []int
	for v := range ch {
		if v == 5 {
			close(doneCh)
			break
		}
		res = append(res, v)
	}

	return res
}

// Буферизированные каналы и отличие от небуферизированных
// задача 7
func mailBox() []string {
	ch := make(chan string, 2)

	ch <- "first"
	ch <- "second"

	return []string{
		<-ch,
		<-ch,
	}
}

// задача 8
func overflowTest() string {
	ch := make(chan string, 1)

	ch <- "first"
	fmt.Println(<-ch)
	ch <- "second"

	return <-ch
}

// задача 9
func messageQueue() []string {
	ch := make(chan string, 3)

	ch <- "one"
	ch <- "two"
	ch <- "three"

	return []string{
		<-ch,
		<-ch,
		<-ch,
	}
}

// задача 10
func sumBuffered(nums []int) int {
	inputCh := make(chan int, len(nums))
	outputCh := make(chan int, 1)

	for _, num := range nums {
		inputCh <- num
	}
	close(inputCh)

	go func() {
		allSum := 0
		for v := range inputCh {
			allSum += v
		}
		outputCh <- allSum
	}()

	return <-outputCh
}

// задача 11
func producerConsumer() []string {
	ch := make(chan string, 2)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- "msg" + strconv.Itoa(i)
		}
		close(ch)
	}()

	var res []string
	for v := range ch {
		res = append(res, v)
	}

	return res
}

// Закрытие каналов и сигналы завершения
// задача 12
func readAll() []string {
	ch := make(chan string)

	go func() {
		ch <- "one"
		ch <- "two"
		ch <- "three"
		close(ch)
	}()

	var res []string
	for v := range ch {
		res = append(res, v)
	}

	return res
}

// задача 13
func checkClosed() string {
	ch := make(chan string)
	close(ch)

	if _, ok := <-ch; !ok {
		return "closed"
	}

	return "open"
}

// задача 14
func signalDone() string {
	doneCh := make(chan struct{})

	go func() {
		time.Sleep(1 * time.Millisecond)
		doneCh <- struct{}{}
	}()
	<-doneCh

	return "done"
}

// задача 15
func multiSignal() string {
	dataCh := make(chan int)

	go func() {
		defer close(dataCh)
		for i := 1; i <= 3; i++ {
			dataCh <- i
		}
	}()

	for range dataCh {
	}

	return "finished"
}

// задача 16
func safeClose() (result string) {
	ch := make(chan string)

	defer func() {
		if r := recover(); r != nil {
			result = "panic prevented"
		}
	}()

	close(ch)
	close(ch)

	return "closed safely"
}

// задача 17
func multiReceiver() string {
	ch := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- "msg" + strconv.Itoa(i)
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for range ch {
		}
	}()

	go func() {
		defer wg.Done()
		for range ch {
		}
	}()
	wg.Wait()

	return "done"
}

// задача от ChatGPT Fan-in(слияние каналов)
func fanIn() []int {
	ch1 := make(chan int)
	ch2 := make(chan int)
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer close(ch1)
		for i := 1; i <= 3; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 4; i <= 6; i++ {
			ch2 <- i
		}
	}()

	go func() {
		defer wg.Done()
		for v := range ch1 {
			out <- v
		}

	}()

	go func() {
		defer wg.Done()
		for v := range ch2 {
			out <- v
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	var res []int
	for v := range out {
		res = append(res, v)
	}

	return res
}

// Однонаправленные каналы (chan<- и <-chan)
// задача 1
func sendData(ch chan<- string) {
	ch <- "Go"
	ch <- "Rust"
	ch <- "Python"
	close(ch)
}
func receiveData(ch <-chan string) []string {
	var res []string
	for v := range ch {
		res = append(res, v)
	}
	return res
}

// задача 2
func producer(ch chan<- int) {
	defer close(ch)
	for i := 1; i <= 3; i++ {
		ch <- i
	}
}
func consumer(ch <-chan int) string {
	for range ch {
	}

	return "done"
}

// Pro-задачи
// задача 2
func worker(id int, jobs <-chan int, results chan<- int) {
	for v := range jobs {
		results <- v * v
		fmt.Printf("worker %d обрабатывает задачу %d\n", id, v)
	}
}

// задача 3
func generator(out chan<- int) {
	for i := 1; i <= 5; i++ {
		out <- i
	}
	close(out)
}

func doubler(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * 2
	}
	close(out)
}

func printer(in <-chan int, done chan<- struct{}) {
	for v := range in {
		fmt.Println(v)
	}
	close(done)
}

// задача 4
func longTask(done chan struct{}) {
	fmt.Println("Запуск задачи...")
	time.Sleep(3 * time.Second)
	done <- struct{}{}
}
