// concurrency/channel-non-blocking/begin/main.go
package main

import "fmt"

func main() {
	// declare a signal channel to read from

	signalChannel := make(chan struct{})
	// use a default case in select to prevent blocking
	select{
	case <-signalChannel:
		fmt.Println("recieved from signal chan")
	default:
		fmt.Println("no signal recieved")
	}
}
