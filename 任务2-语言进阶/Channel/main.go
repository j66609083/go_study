package main

import (
	"fmt"
	"time"
)

// 生成从1到10的整数
func generateNumbers(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

// 从通道中接收这些整数并打印出来
func printNumbers(ch chan int) {
	for i := range ch {
		fmt.Println("Received numbers:", i)
	}
}

// 生产者协程向通道中发送100个整数
func numberProducer(buffer chan<- int) {
	for i := 1; i <= 100; i++ {
		buffer <- i
	}
	close(buffer)
}

// 消费者协程从通道中接收这些整数并打印
func numberConsumer(buffer <-chan int) {
	for num := range buffer {
		fmt.Println("Consumed number:", num)
	}
}

func main() {

	/*
		题目 ：编写一个程序，使用通道实现两个协程之间的通信。
		一个协程生成从1到10的整数，并将这些整数发送到通道中，
		另一个协程从通道中接收这些整数并打印出来。
	*/
	ch := make(chan int)
	go generateNumbers(ch)
	go printNumbers(ch)

	/*
		题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，
		消费者协程从通道中接收这些整数并打印。
	*/
	buffer := make(chan int, 50)
	go numberProducer(buffer)
	go numberConsumer(buffer)

	time.Sleep(5 * time.Second)
}
