package main

import (
	"fmt"
	"sync"
	"time"
)

func send(ch chan<- int) {
	i := 0
	for {
		time.Sleep(time.Second * 1)
		select {
		case ch <- i:
			i++
		}
	}
}

func work(ch <-chan int, wg *sync.WaitGroup) {
	ticker := time.NewTicker(2 * time.Second)
	defer wg.Done()
	for {
		time.Sleep(500 * time.Millisecond)
		select {
		case data := <-ch:
			fmt.Println("接收数据 : ", data)
		case <-ticker.C:
			fmt.Println("执行上报, 上报心跳")
		default:
			fmt.Println("队列为空")
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go send(ch)
	go work(ch, &wg)

	wg.Wait()
}
