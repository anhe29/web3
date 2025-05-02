package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	task9()
	task10()
}

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。

func task9() {
	var count int
	var mx = sync.Mutex{}
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				mx.Lock()
				count++
				mx.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

//使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

func task10() {
	var count int32
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				atomic.AddInt32(&count, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
