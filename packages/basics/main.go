// packages/basics/main.go
// go run runs and compiles packages
package main
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello go")
	fmt.Printf("today is %s\n", time.Now().Weekday())
}
