package main

import "fmt"

func main() {
	c := make(chan int)
	go send(c)
	receive(c)
}

func send(c chan<- int) {
	c <- 7
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}
