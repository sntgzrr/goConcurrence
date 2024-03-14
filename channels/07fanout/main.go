package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// FanOut es un patron de diseño para trabajar con código concurrente.
func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go send(c1)
	go fanOutIn(c1, c2)
	//go fanOutIn2(c1, c2)
	for v := range c2 {
		fmt.Println(v)
	}
	fmt.Println("Ending...")
}

func send(c1 chan int) {
	for i := 0; i < 100; i++ {
		c1 <- i
	}
	close(c1)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v int) {
			c2 <- timeProcess(v)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

// Regulado
func fanOutIn2(c1, c2 chan int) {
	var wg sync.WaitGroup
	goroutines := 10
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			for v := range c1 {
				func(v int) {
					c2 <- timeProcess(v)
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}

func timeProcess(v int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	return v + rand.Intn(1000)
}
