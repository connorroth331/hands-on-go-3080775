// concurrency/sync/begin/main.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	// given a list of names
	names := []string{"james", "c", "j"}

	nameMap := make(map[string]int)

	// initialize a map to store the length of each name

	// initialize a wait group for the goroutines that will be launched

	var wg sync.WaitGroup
	wg.Add(len(names))
	// launch a goroutine for each name we want to process

	var m sync.Mutex
	for _, name := range names {
		go func(name string){
			defer wg.Done()
			m.Lock()
			defer m.Unlock()
			nameMap[name] = len(name)

		}(name)
	}
	// wait for all goroutines to finish

	wg.Wait()
	// print the map
	fmt.Println(nameMap)
}
