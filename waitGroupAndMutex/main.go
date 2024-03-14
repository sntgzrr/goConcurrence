package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Eliminando Race condition con Mutex.
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	fmt.Println("CPU's:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	counter := 0

	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			runtime.Gosched()
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines number:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
