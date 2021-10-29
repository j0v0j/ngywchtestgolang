package main

import (
	"fmt"
	"sync"
	"test1/libdll"
	"time"
)

var wg sync.WaitGroup

func hello() {
	for i := 0; i < 100; i++ {
		time.Sleep(500 * time.Microsecond)
		fmt.Printf("%v\r\n", i)
	}
	defer wg.Done()
}
func main() {
	// modbustest.Modbustest()

	// for i := 0; i < 10; i++ {
	// 	wg.Add(1) // 启动一个goroutine就登记+1
	// 	go hello(i)
	// }
	wg.Add(1)
	go hello()
	wg.Add(1)
	go libdll.Usedll(&wg)
	wg.Wait() // 等待所有登记的goroutine都结束
	// time.Sleep(20 * time.Second)
}
