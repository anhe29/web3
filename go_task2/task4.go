package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	ch := make(chan int, 10)
// 	go sent(ch, &wg)
// 	go recive(ch, &wg)
// 	wg.Wait()
// 	task8()
// }

//题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，
// 另一个协程从通道中接收这些整数并打印出来。
//考察点 ：通道的基本使用、协程间通信。

func sent(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
	}
}

func recive(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var num int
	for i := 1; i <= 10; i++ {
		num = <-ch
		fmt.Println(num)
	}
}

//实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，
// 消费者协程从通道中接收这些整数并打印。
//考察点 ：通道的缓冲机制。

func task8() {
	mg := sync.WaitGroup{}
	mg.Add(2)
	ch := make(chan int, 100)

	go func() {
		defer mg.Done()
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		defer mg.Done()
		for num := range ch {
			fmt.Println(num)
		}
	}()
	mg.Wait()
}
