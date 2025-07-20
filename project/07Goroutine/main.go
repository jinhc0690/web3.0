package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

////=====================================================================================================

// Goroutine 1.题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。 考察点 ： go 关键字的使用、协程的并发执行。

func jishu() {
	for i := 1; i <= 10; i += 2 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}
}

func oushu() {
	for i := 2; i <= 10; i += 2 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}
}

// Goroutine 2.题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。考察点 ：协程原理、并发任务调度。
type Task interface {
	oushu()
	jishu()
}

type Sche struct {
	time int
}

func (s Sche) jishu(si string) {
	for i := 1; i <= 10; i += 2 {
		time.Sleep(time.Duration(s.time) * time.Millisecond)
		fmt.Println(i, si)
	}
}

func (s Sche) oushu(si string) {
	for i := 2; i <= 10; i += 2 {
		time.Sleep(time.Duration(s.time) * time.Millisecond)
		fmt.Println(i, si)
	}
}

func main() {

	// Goroutine 1.题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。 考察点 ： go 关键字的使用、协程的并发执行。
	go jishu()
	go oushu()

	// Goroutine 2.题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。考察点 ：协程原理、并发任务调度。
	s1 := Sche{
		time: 100,
	}

	s2 := Sche{
		time: 1000,
	}

	go s1.jishu("s1奇数")
	go s1.oushu("s1偶数")

	go s2.jishu("s2奇数")
	go s2.oushu("s2偶数")

	say("hello")
}
