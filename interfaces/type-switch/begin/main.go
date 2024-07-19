// interfaces/type-switch/begin/main.go
package main
import "fmt"
// define whatAmI which takes in an argument of any type and returns inforamtion about the underlying value's type
func whatAmI(val interface{}) string{
	switch val.(type){
	case int:
		return "integer"
	case string:
		return "string"
	case bool: 
		return "boolean"
	default:
		return "not supported"
	}
}
func main() {
	// invoke whatAmI function
	fmt.Println(whatAmI(1))
	fmt.Println(whatAmI("hello"))
	fmt.Println(whatAmI(true))
}
