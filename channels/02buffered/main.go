package main

import "fmt"

func main() {
	//Buffered channel --Se menciona la cantidad de datos que se pueden quedar en el canal.
	ca := make(chan int, 1)
	ca <- 7
	fmt.Println(<-ca)
}
