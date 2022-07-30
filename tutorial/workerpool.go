package main

import "fmt"

// use worker to calculate fibonanci
func main() {
	number := 1000
	numberOfWorker := 5
	jobs := make(chan int, number)
	results := make(chan int, number)
	// start worker
	for i := 0; i < numberOfWorker; i++ {
		go worker(jobs, results)
	}
	// push value to jobs
	for i := 1; i <= number; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < number; i++ {
		fmt.Printf("%d ", <-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
