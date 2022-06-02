package main

func main() {
	//f1()
	//f2()
	f3()
}

func f1() {
	ch := make(chan int)
	ch <- 1
}

func f2() {
	ch := make(chan int)
	<-ch
}

func f3() {
	ch := make(chan int)
	ch <- 1
	<-ch
}
