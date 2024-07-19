// interfaces/basics/begin/main.go
package main
import "fmt"

type humanoid interface{
	speak()
	walk()
}

type person struct {name string}

func (p person) speak() {fmt.Println("speaking " + p.name)}
func (p person) walk() {fmt.Println("walking " + p.name)}
// define a humanoid interface with speak and walk methods returning string

// define a person type that implements humanoid interface

// implement the Stringer interface for the person type
type dog struct{}
func (d dog) walk() {fmt.Println("DOG WALK")}
// define a dog type that can walk but not speak

func main() {
	// invoke with a person
	// doHumanThings(person{})
	p := person{name: "cj"}
	doHumanThings(p)
	fmt.Println(p)
	// d := dog{}

	// // can we invoke with a dog?
	//  doHumanThings(d)

	// fmt.Println(p)
}

func doHumanThings(h humanoid) {
	h.speak()
	h.walk()
}
