package main

import "fmt"

func fib(pos int) int {
	if pos <= 1 {
		return pos
	}

	return fib(pos-2) + fib(pos-1)
}

// sends (<-chan); receives (chan<-)
func worker(tasks <-chan int, results chan<- int) {
	for task := range tasks {
		results <- fib(task)
	}
}

// Worker pools
func main() {
	loops := 45

	tasks := make(chan int, loops)
	results := make(chan int, loops)

	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < loops; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < loops; i++ {
		result := <-results
		fmt.Println(result)
	}
}
