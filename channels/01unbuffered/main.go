package main

import "fmt"

func main() {
	/*Unbuffered channel -- Se bloquea cuando se envía una Goroutine pero no hay una que lo reciba, es decir, debe de haber una
	goroutine que envie y otra que esté recibiendo. (Transmisor, receptor)*/
	ca := make(chan int)
	go func() {
		ca <- 42
	}()
	fmt.Println(<-ca)
}
