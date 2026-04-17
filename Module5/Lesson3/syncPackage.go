package Lesson3

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Task struct {
	ID int
}

type UserCache struct {
	mu   sync.RWMutex
	data map[string]string
}

func (uc *UserCache) Set(name, value string) {
	defer uc.mu.Unlock()
	uc.mu.Lock()
	uc.data[name] = value
	fmt.Println("Writer обновил", name, "->", value)
}
func (uc *UserCache) Get(name string) {
	defer uc.mu.RUnlock()
	uc.mu.RLock()
	value := uc.data[name]
	fmt.Println("Reader:", name, "=", value)
}

type Stats struct {
	mu      sync.Mutex
	Total   int
	Success int
	Failed  int
}

func (s *Stats) Add(success bool) {
	defer s.mu.Unlock()
	s.mu.Lock()
	s.Total++
	if success == true {
		s.Success++
	} else {
		s.Failed++
	}
}

type User struct {
	ID   int
	Name string
}

type Cache struct {
	pool sync.Pool
}

func (c *Cache) Put(value string) {
	fmt.Println("Added:", value)
	c.pool.Put(value)

}
func (c *Cache) Get() string {
	v := c.pool.Get().(string)
	return v
}

type SafeCounter struct {
	mu    sync.RWMutex
	count int
}

func (sc *SafeCounter) Inc() {
	defer sc.mu.Unlock()
	sc.mu.Lock()
	sc.count++
}
func (sc *SafeCounter) Value() int {
	defer sc.mu.RUnlock()
	sc.mu.RLock()
	return sc.count
}

type Database struct {
	mu   sync.RWMutex
	data map[string]string
}

func (db *Database) Write(key, value string) {
	defer db.mu.Unlock()
	db.mu.Lock()
	db.data[key] = value
	fmt.Println("Writer:", key, "updated")
}
func (db *Database) Read(key string) string {
	defer db.mu.RUnlock()
	db.mu.RLock()
	value := db.data[key]
	fmt.Println("Reader:", key, "->", value)
	return value
}

type AtomicCounter struct {
	value int64
}

func (c *AtomicCounter) Inc() {
	atomic.AddInt64(&c.value, 1)
}
func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func Main() {
	// WaitGroup — зачем нужна и как использовать
	// задача 1
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Привет из горутины")
	}()
	wg.Wait()

	// задача 2
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go func(n int) {
			defer wg.Done()
			fmt.Println("Горутина", i)
		}(i)
	}
	wg.Wait()

	// задача 3
	doneCh := make(chan struct{})
	for i := 1; i <= 3; i++ {
		go func(n int) {
			fmt.Printf("[done] Горутина %d завершена\n", i)
			doneCh <- struct{}{}
		}(i)
	}
	<-doneCh
	<-doneCh
	<-doneCh

	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go func(n int) {
			defer wg.Done()
			fmt.Printf("[wg]   Горутина %d завершена\n", i)
		}(i)
	}
	wg.Wait()

	// задача 4
	tasks := []string{"one", "two", "three", "four", "five"}
	ch := make(chan string)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for task := range ch {
				fmt.Printf("Worker %d: %s\n", workerID, task)
			}
		}(i)
	}

	for _, task := range tasks {
		ch <- task
	}
	close(ch)

	wg.Wait()

	// Гонки данных — зачем нужен sync вообще
	// задача 1
	var counter int

	for i := 0; i < 2; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				counter++
			}
		}()
	}
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Результат:", counter)

	// Mutex и RWMutex — управление доступом к данным
	// задача 1
	var counter1 int
	var mu sync.Mutex

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			counter1++
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			counter1++
			mu.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println("Final counter value:", counter1)

	// задача 2
	turn := "A"
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 3; {
			mu.Lock()
			if turn != "A" {
				mu.Unlock()
				continue
			}
			fmt.Print(turn + " ")
			turn = "B"
			i++
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 3; {
			mu.Lock()
			if turn != "B" {
				mu.Unlock()
				continue
			}
			fmt.Print(turn + " ")
			turn = "A"
			i++
			mu.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println()

	// задача 3
	sc := &SafeCounter{}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 2500; i++ {
				sc.Inc()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final count:", sc.Value())

	// задача 4
	sc1 := &SafeCounter{}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 2500; j++ {
				sc1.Inc()
			}
		}()
	}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				_ = sc1.Value()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final count:", sc1.Value())

	// задача 5
	db := Database{data: make(map[string]string)}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			db.Write("user1", "Alice")
		}()
	}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			db.Read("user1")
		}()
	}

	wg.Wait()

	// задача 6
	var printer sync.Mutex

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			printer.Lock()
			defer printer.Unlock()
			fmt.Printf("User %d started printing...\n", n)
			time.Sleep(time.Second)
			fmt.Printf("User %d finished.\n", n)
		}(i)
	}

	wg.Wait()

	// Atomic — атомарные операции
	// задача 1
	var counter2 int64
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			atomic.AddInt64(&counter2, 1)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			atomic.AddInt64(&counter2, 1)
		}
	}()

	wg.Wait()
	fmt.Println("Final counter:", counter2)

	// задача 2
	fmt.Println("Mutex:", runWithMutex())
	fmt.Println("Atomic:", runWithAtomic())

	// задача 3
	ac := &AtomicCounter{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				ac.Inc()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final value:", ac.Value())

	// sync.Pool — оптимизация и переиспользование объектов
	// задача 1
	pool := sync.Pool{New: func() any {
		fmt.Println("Создан новый буфер")
		return make([]byte, 4)
	}}

	buf := pool.Get().([]byte)
	buf[0] = 1
	buf[1] = 2
	buf[2] = 3
	buf[3] = 4

	fmt.Println("Используем буфер:", buf)
	pool.Put(buf)

	buf1 := pool.Get().([]byte)
	fmt.Println("Повторно используем:", buf1)

	// задача 2
	pool1 := sync.Pool{New: func() any {
		fmt.Println("Создан новый пользователь")
		return &User{
			ID:   1,
			Name: "Alice"}
	}}

	user := pool1.Get().(*User)
	fmt.Printf("User: ID=%d, name=%s\n", user.ID, user.Name)

	user.ID = 2
	user.Name = "Bob"
	fmt.Printf("User: ID=%d, name=%s\n", user.ID, user.Name)

	pool1.Put(user)

	user1 := pool1.Get().(*User)
	fmt.Printf("User: ID=%d, name=%s\n", user1.ID, user1.Name)

	// задача 3
	pool2 := sync.Pool{New: func() any {
		fmt.Println("Создана новая задача")
		return func(id int) {
			fmt.Printf("Worker %d выполняет задачу\n", id)
			time.Sleep(100 * time.Millisecond)
		}
	},
	}

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			task := pool2.Get().(func(int))
			task(id)
			pool2.Put(task)
		}(i)
	}

	wg.Wait()

	// задача 4
	cache := Cache{}
	cache.pool.New = func() any { return "empty" }

	cache.Put("hello")
	cache.Put("world")

	fmt.Println("Get:", cache.Get())
	fmt.Println("Get:", cache.Get())
	fmt.Println("Get:", cache.Get())

	// Context — управление временем и отменой задач
	// задача 1
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, 1)
	go worker(ctx, 2)

	time.Sleep(time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond)

	// задача 2
	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	defer cancel1()
	wg.Add(1)
	go longTask(ctx1, &wg)
	wg.Wait()

	// задача 3
	parent, cancel2 := context.WithCancel(context.Background())
	child, _ := context.WithTimeout(parent, 2*time.Second)

	go func() {
		<-child.Done()
		fmt.Println("Child контекст завершён:", child.Err())
	}()

	time.Sleep(time.Second)
	cancel2()
	time.Sleep(100 * time.Millisecond)

	// задача 4
	ctx2, cancel3 := context.WithCancel(context.Background())
	defer cancel3()
	errCh := make(chan error, 3)
	wg.Add(3)

	go func() {
		defer wg.Done()
		err := fetchData(ctx2)
		if err != nil {
			errCh <- err
			cancel3()
		}
	}()
	go func() {
		defer wg.Done()
		err := processData(ctx2)
		if err != nil {
			errCh <- err
			cancel3()
		}
	}()
	go func() {
		defer wg.Done()
		err := saveData(ctx2)
		if err != nil {
			errCh <- err
			cancel3()
		}
	}()

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil && !errors.Is(err, context.Canceled) {
			fmt.Println("Ошибка", err)
			fmt.Println("Контекст отменён — прекращаем выполнение")
		}
	}

	// Pro-задачи по sync
	// задача 1
	stats := Stats{}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 50; i++ {
				if rand.Intn(2) == 0 {
					stats.Add(true)
				} else {
					stats.Add(false)
				}
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Всего: %d, Успешно: %d, Ошибок: %d\n", stats.Total, stats.Success, stats.Failed)

	// задача 2
	var data []int

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 1; i <= 5; i++ {
				mu.Lock()
				data = append(data, i)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Результат:", data)

	// задача 3
	userCache := UserCache{data: make(map[string]string)}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			userCache.Get("user1")
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		userCache.Set("user1", "Alice")
	}()

	wg.Wait()

	// задача 4
	var counter3 int64
	taskPool := sync.Pool{New: func() any { return &Task{} }}

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			tp := taskPool.Get().(*Task)
			tp.ID = id
			fmt.Printf("Task %d выполнена\n", tp.ID)
			taskPool.Put(tp)
			atomic.AddInt64(&counter3, 1)
		}(i)
	}

	wg.Wait()
	fmt.Println("Всего выполнено:", counter3)

	// задача 5
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var results []int

	for _, number := range numbers {
		n := number
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			mu.Lock()
			results = append(results, n*n)
			mu.Unlock()
		}(n)
	}

	wg.Wait()
	fmt.Println("Результаты:", results)
}

// Atomic — атомарные операции
// задача 2
func runWithMutex() time.Duration {
	start := time.Now()
	var mu sync.Mutex
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	return time.Since(start)
}

func runWithAtomic() time.Duration {
	start := time.Now()
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	wg.Wait()
	return time.Since(start)
}

// Context — управление временем и отменой задач
// задача 1
func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина", id, "завершена")
			return
		default:
			fmt.Println("Горутина", id, "работает")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

// задача 2
func longTask(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Контекст отменён:", ctx.Err())
			return
		default:
			fmt.Println("выполняется...")
			time.Sleep(300 * time.Millisecond)
		}
	}
}

// задача 4
func fetchData(ctx context.Context) error {
	select {
	case <-ctx.Done():
		fmt.Println("fetchData остановлен")
		return ctx.Err()
	case <-time.After(300 * time.Millisecond):
		fmt.Println("fetchData: ok")
		return nil
	}
}

func processData(ctx context.Context) error {
	select {
	case <-ctx.Done():
		fmt.Println("processData остановлен")
		return ctx.Err()
	case <-time.After(300 * time.Millisecond):
		fmt.Println("processData: ok")
		return nil
	}
}

func saveData(ctx context.Context) error {
	select {
	case <-ctx.Done():
		fmt.Println("saveData остановлен")
		return ctx.Err()
	case <-time.After(300 * time.Millisecond):
		fmt.Println("saveData: ok")
		return nil
	}
}
