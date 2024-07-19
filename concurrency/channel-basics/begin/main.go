// concurrency/channel-basics/begin/main.go
package main

import (
	"fmt"
	//"time"
)

// sum calculates and prints the sum of numbers
func sum(nums []int, ch chan<-int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	//fmt.Println("Result:", sum)
	ch <- sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	ch1 := make(chan int) //int resuklt
	// invoke the sum function as a goroutine
	go sum(nums, ch1)

	result := <-ch1
	fmt.Println(result)
	ch2 := make(chan string, 2)
	ch2 <- "hello"
	ch2 <- "world"

	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	// force main thread to sleep
	//time.Sleep(100 * time.Millisecond)
}
