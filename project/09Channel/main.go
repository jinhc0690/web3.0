package main

import (
	"fmt"
	"time"
)

////=====================================================================================================

// Channel 1.题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。 考察点 ：通道的基本使用、协程间通信。
// 只接收channel的函数
func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接收到: %d\n", v)
	}
}

// 只发送channel的函数
func sendOnly(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("发送: %d\n", i)
	}
	close(ch)
}

// Channel 2.题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。 考察点 ：通道的缓冲机制。
// 只接收channel的函数
func receiveOnly1(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接收到: %d\n", v)
	}
}

// 只发送channel的函数
func sendOnly1(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Printf("发送: %d\n", i)
	}
	close(ch)
}

func main() {

	// Channel 1.题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。 考察点 ：通道的基本使用、协程间通信。
	ch := make(chan int)

	// 启动发送goroutine
	go sendOnly(ch)

	// 启动接收goroutine
	go receiveOnly(ch)

	// 使用select进行多路复用
	timeout := time.After(2 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Channel已关闭")
				return
			}
			fmt.Printf("主goroutine接收到: %d\n", v)
		case <-timeout:
			fmt.Println("操作超时")
			return
		default:
			fmt.Println("没有数据，等待中...")
			time.Sleep(500 * time.Millisecond)
		}
	}

	// Channel 2.题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。 考察点 ：通道的缓冲机制。
	ch1 := make(chan int, 3)

	// 启动发送goroutine
	go sendOnly1(ch1)

	// 启动接收goroutine
	go receiveOnly1(ch1)

	// // 使用select进行多路复用
	timeout1 := time.After(2 * time.Second)
	for {
		select {
		case v, ok := <-ch1:
			if !ok {
				fmt.Println("Channel已关闭")
				return
			}
			fmt.Printf("主goroutine接收到: %d\n", v)
		case <-timeout1:
			fmt.Println("操作超时")
			return
		default:
			fmt.Println("没有数据，等待中...")
			time.Sleep(500 * time.Millisecond)
		}
	}

}
