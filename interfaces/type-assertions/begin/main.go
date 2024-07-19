// interfaces/type-assertions/begin/main.go
package main
import "fmt"
func main() {
	// perform a type assertion

	var i any = "hello"

	//assumes it is a string
	fmt.Println(i.(string))
	if _, ok := i.(int); !ok{
		fmt.Println("not int")

	}

	

	// perform a type assertion and handle the error
}
