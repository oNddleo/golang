// package main

// import "fmt"

// // make chan sum square number
// func main() {
// 	// random number
// 	randomNumbers := []int{}
// 	for i := 0; i < 1000000; i++ {
// 		randomNumbers = append(randomNumbers, i)
// 	}
// 	// generate pipeline
// 	inputChan := generatePipeline(randomNumbers)

// 	// fanout
// 	c1 := fanout(inputChan)
// 	c2 := fanout(inputChan)
// 	c3 := fanout(inputChan)
// 	c4 := fanout(inputChan)

// 	// fanin
// 	c := fanin(c1, c2, c3, c4)

// 	// sum from fanin chan
// 	sum := 0
// 	for i := 0; i < len(randomNumbers); i++ {
// 		sum += <-c
// 	}
// 	fmt.Printf("Total sum = %d", sum)
// }

// func generatePipeline(numbers []int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for _, n := range numbers {
// 			out <- n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func fanout(in <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for n := range in {
// 			out <- n * n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func fanin(inputChannel ...<-chan int) <-chan int {
// 	in := make(chan int)
// 	go func() {
// 		for _, c := range inputChannel {
// 			for n := range c {
// 				in <- n
// 			}
// 		}
// 	}()
// 	return in
// }
