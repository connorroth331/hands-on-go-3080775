// types/conversion/begin/main.go
package main
import "fmt"
func main() {
	a := 42.5
	b := uint(a)
	fmt.Printf("%v=%T, %v =%T", a, a, b, b)
	// declare float and convert it to an unsigned int
}
