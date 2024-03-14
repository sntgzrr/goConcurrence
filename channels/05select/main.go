package main

import "fmt"

func main() {
	pair := make(chan int)
	odd := make(chan int)
	exit := make(chan int)

	go send(pair, odd, exit)
	receive(pair, odd, exit)
	fmt.Println("Ending....")
}

func send(p, o, e chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			p <- i
		} else {
			o <- i
		}
	}
	close(e)
}

func receive(p, o, e <-chan int) {
	for {
		select {
		case v := <-p:
			fmt.Println("Canal par:", v)
		case v := <-o:
			fmt.Println("Canal impar:", v)
		case v, ok := <-e:
			if !ok {
				fmt.Println("Canal salir:", v)
				return
			}
		}
	}
}
