// concurrency/channel-select/begin/main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	// declare an empty struct channel for signaling
	//EMPTY STRUCTS ARE GOOD BECAUSE THEY DO NOT STORE MEMORY
	sigChan := make(chan struct{})

	// declare a timer channel
	timeChanel := time.After(1 * time.Second)

	// launch a goroutine to signal after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		sigChan <- struct{}{}
	}()

	// wait for a signal on either channel
	//
	select{
		case <- sigChan:
			fmt.Println("signal channel here")
		case <- timeChanel:
			fmt.Println("time channel here")
	}
}
