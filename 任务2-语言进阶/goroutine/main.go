package main

import (
	"fmt"
	"time"
)

// 打印从1到10的奇数
func printOddNumbers() {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

// 打印从2到10的偶数
func printEvenNumbers() {
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

type Task struct {
	name string
	t    func()
}

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
func taskScheduler(tasks []Task) {
	for _, task := range tasks {
		go func() {
			start := time.Now()
			task.t()
			elapsed := time.Since(start)
			fmt.Printf("Task(%s) completed in %s\n", task.name, elapsed)
		}()
	}
}

func main() {
	go printEvenNumbers()
	go printOddNumbers()

	taskScheduler([]Task{
		{name: "Print Odd Numbers", t: printOddNumbers},
		{name: "Print Even Numbers", t: printEvenNumbers},
	})
	time.Sleep(5 * time.Second)
}
