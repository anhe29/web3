package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// func main() {
// 	task3()
// 	ScheduleRun()
// }

// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
func task3() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Println(i)
			}
		}
		defer wg.Done()
	}()
	go func() {
		for i := 2; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
		defer wg.Done()
	}()
	wg.Wait()
}

// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。

type result struct {
	taskName string
	time     time.Duration
}

func ScheduleRun() {
	taskList := []func(){
		func() {
			time.Sleep(time.Second * 1)
		},
		func() {
			time.Sleep(time.Second * 2)
		},
		func() {
			time.Sleep(time.Second * 3)
		},
	}
	result := task4(&taskList)
	for _, res := range result {
		fmt.Println("任务名称：", res.taskName, " 执行时间：", res.time)
	}
}

func task4(task *[]func()) []result {
	wg := sync.WaitGroup{}
	results := make([]result, len(*task))

	for index, ts := range *task {
		wg.Add(1)
		go func(index int, ts func()) {
			defer wg.Done()
			start := time.Now()
			ts()
			end := time.Since(start)
			results[index] = result{
				taskName: "task" + strconv.Itoa(index),
				time:     end,
			}
		}(index, ts)
	}
	wg.Wait()
	return results
}
