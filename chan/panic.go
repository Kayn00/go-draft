package main

import (
	"time"
)

func main() {
	p1()
}

func p1() {
	ch := make(chan int, 1)
	close(ch)
	ch <- 1
}

func p2() {
	ch := make(chan int, 1)
	done := make(chan struct{}, 1)
	go func() {
		<-time.After(2 * time.Second)
		println("close2")
		close(ch)
		close(done)
	}()

	go func() {
		<-time.After(1 * time.Second)
		println("close1")
		ch <- 1
		close(ch)
	}()
}
