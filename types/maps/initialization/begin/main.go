// types/maps/initialization/begin/main.go
package main
import  "fmt"
type author struct {
	name string
}

func main() {
	// declare a map of string keys and author values
	//
	var authors map[string]author


	// initialize the map with make
	//
	authors = make(map[string]author)



	// add authors to the map
	//
	authors["hello"] = author{name:"world"}
	authors["world"] = author{name:"hello"}

	// print the map with fmt.Printf
	//
	authors["world"] = author{name: "big"}
	authors["new key"] = author{name: "deleteMe"}

	fmt.Println(authors)

	// read a value from the map with a known key
	//
	fmt.Println(authors["hello"])

	delete(authors, "new key")
	a, ok := authors["new key"]
	fmt.Printf("a = %v, ok = %v\n", a, ok)
}
