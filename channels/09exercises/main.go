package main

func main() {
	/*Exercise 01
	c := gen()
	receive(c)
	fmt.Println("Ending...")*/

	/*Exercise 02
	exit := make(chan int)
	c := gen(exit)
	receive(c, exit)
	fmt.Println("Ending...")*/

	/*Exercise 03
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}*/

	/*Exercise 04
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}
	for v := 0; v < 100; v++ {
		fmt.Println(<-c)
	}*/
}

/*Exercise 01
func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}*/

/*Exercise 02
func gen(exit chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		exit <- 0
		close(c)
	}()
	return c
}

func receive(c, exit <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-exit:
			return
		}
	}
}*/
