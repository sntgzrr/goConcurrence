package main

import (
	"fmt"
	"sync"
)

func main() {
	pair := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)
	go send(pair, odd)
	go receive(pair, odd, fanin)
	for v := range fanin {
		fmt.Println(v)
	}
	fmt.Println("Ending...")
}

func send(pair, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			pair <- i
		} else {
			odd <- i
		}
	}
	close(pair)
	close(odd)
}

func receive(pair, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range pair {
			fanin <- v
		}
		wg.Done()
	}()
	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)
}
