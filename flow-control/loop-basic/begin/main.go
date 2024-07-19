// flow-control/loop-basic/begin/main.go
package main
import "fmt"
func main() {
	// declare a string to iterate over
	//
	s := "moby dick"

	for i:=0; i < len(s); i++{
		fmt.Println(string(s[i]))
	}

	// iterate over the string with basic for loop
	//
	n := []int{200, 300, 400}
	for i, v := range n{
		fmt.Println(i, v)
	}

	numsmap := map[string]int{"a":1, "b":2, "c":3}

	for key, value := range numsmap {
		fmt.Println(key, value)
	}

}
