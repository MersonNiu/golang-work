package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

// 指针
// 题目1
func plusten(num *int) int {
	return *num + 10
}

// 题目2
func slicepointer(slice *[]int) []int {
	for i, value := range *slice {
		(*slice)[i] = value * 2
	}
	return *slice
}

// Goroutine
// 题目1
func goprint() {
	go func() {
		for i := range 10 {
			if i%2 == 1 {
				fmt.Println(i)
			}
		}
	}()
	go func() {
		for i := range 11 {
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	}()
}

// 题目2
type Task func()
type TaskResult struct {
	Name     string
	Duration time.Duration
}
type Scheduler struct {
	tasks []struct {
		name string
		task Task
	}
}

func (s *Scheduler) AddTask(name string, task Task) {
	s.tasks = append(s.tasks, struct {
		name string
		task Task
	}{name: name, task: task})
}
func (s *Scheduler) Run() map[string]time.Duration {
	var wg sync.WaitGroup
	results := make(chan TaskResult, len(s.tasks))
	for _, item := range s.tasks {
		wg.Add(1)
		go func(name string, task Task) {
			defer wg.Done()
			start := time.Now()
			defer func() {
				results <- TaskResult{
					Name:     name,
					Duration: time.Since(start)}
			}()
			task()
		}(item.name, item.task)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	resultsMap := make(map[string]time.Duration)
	for res := range results {
		resultsMap[res.Name] = res.Duration
	}
	return resultsMap
}

// 面向对象
// 题目1
type Shape interface {
	Area()
	Perimeter()
}
type Rectangle struct {
	lenth  int
	weigth int
}
type Circle struct {
	radius float64
}

func (rec *Rectangle) Area(lenth int, weigth int) int {
	return lenth * weigth

}
func (rec *Rectangle) Perimeterrea(lenth int, weigth int) int {
	return (lenth + weigth) * 2
}
func (rec *Circle) Area(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}
func (rec *Circle) Perimeterrea(radius float64) float64 {
	return 2 * math.Pi * radius
}

// 题目2
type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
}

func (emp *Employee) Printinfo(name string, age int) {
	fmt.Printf("员工%s的年龄是%d", name, age)
}

// Channel

func chanprgraten() {
	sumchan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(sumchan)
		for i := 1; i < 11; i++ {
			sumchan <- i
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range sumchan {
			fmt.Println(i)
		}
	}()
	wg.Wait()
}

// 题目2实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
func chanprgraonehun() {
	sumchan := make(chan int, 100)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(sumchan)
		for i := 1; i < 101; i++ {
			sumchan <- i
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range sumchan {
			fmt.Println(i)
		}
	}()
	wg.Wait()
}

// 锁机制
// 题目1编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 创建一个计数器，附加一个递增方法，启动协程
type countrt struct {
	countnum int
}

var mu sync.Mutex
var wg sync.WaitGroup

func (c *countrt) increment() {
	mu.Lock()
	for i := 0; i < 1000; i++ {
		c.countnum++
	}
	mu.Unlock()
}
func lockptcgo() int {
	counter := countrt{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.increment()
		}()
	}
	wg.Wait()
	return counter.countnum

}

// 题目2
type counteran struct {
	countertwo int64
}

func (c *counteran) atomicincre() {

	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&c.countertwo, 1)
	}
}
func atomicgo() int {
	counter := counteran{}
	for i := 0; i < 10; i++ {
		counter.atomicincre()
	}
	return int(counter.countertwo)
}

func main() {
	// fmt.Println(atomicgo())
	// fmt.Println(lockptcgo())
	// chanprgraonehun()
	// chanprgraten()
	// empone := Employee{}
	// empone.Printinfo("Judy", 23)
	// recone := Rectangle{}
	// perone := Circle{}
	// fmt.Println("这个矩形的面积是：", recone.Area(3, 4))
	// fmt.Println("这个矩形的周长是：", recone.Perimeterrea(3, 4))
	// fmt.Println("这个圆的面积是：", perone.Area(2.5))
	// fmt.Println("这个圆形的周长是：", perone.Perimeterrea(2.5))
	// scheduler := Scheduler{}
	// scheduler.AddTask("任务1", func() {
	// 	time.Sleep(300 * time.Millisecond)
	// 	fmt.Println("任务1完成")
	// })
	// scheduler.AddTask("任务2", func() {
	// 	time.Sleep(500 * time.Millisecond)
	// 	fmt.Println("任务2完成")
	// })
	// scheduler.AddTask("任务3", func() {
	// 	time.Sleep(800 * time.Millisecond)
	// 	fmt.Println("任务3完成")
	// })
	// results := scheduler.Run()
	// for name, duration := range results {
	// 	fmt.Printf("%s: %v\n", name, duration)
	// }
}
