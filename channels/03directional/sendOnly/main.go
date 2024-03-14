package main

func main() {
	ca := make(chan<- int, 1)
	go func() {
		ca <- 7
	}()
}
