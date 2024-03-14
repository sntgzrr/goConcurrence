package main

import "fmt"

func main() {
	ca := make(<-chan int, 1)
	fmt.Println(<-ca)
}
