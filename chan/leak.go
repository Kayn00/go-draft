package main

import (
	"fmt"
	"time"
)

func main() {
	//leak1()
	leak2()
}

func leak1() {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 100
	}()

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("timeout! exit...")
	case result := <-ch:
		fmt.Printf("result:%d\n", result)
	}
}

func leak2() {
	ch := make(chan int)

	go func() {
		for result := range ch {
			fmt.Printf("result:%d\n", result)
		}
	}()

	ch <- 1
	ch <- 2
	time.Sleep(time.Second)
	fmt.Println("main goroutine g2 done...")
}
