// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	wg := sync.WaitGroup{}
// 	ch := make(chan int, 50)
// 	wg.Add(2)
// 	// anonymous function goroutine
// 	// only receive from chan
// 	go func(ch <-chan int) {
// 		i := <-ch
// 		fmt.Println(i)
// 		wg.Done()
// 	}(ch)

// 	// only send to chan
// 	go func(ch chan<- int) {
// 		ch <- 42
// 		wg.Done()
// 	}(ch)
// 	wg.Wait()
// }
