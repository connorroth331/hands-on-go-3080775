// challenges/types-primitive/begin/main.go
package main
import ("fmt")

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"
func printGlobal() string{
	return val
}
func updateGlobal(a string) {
	val = a
}
func main() {
	// create a local variable "val" and assign it an int value
	val := 34

	// print the value of the local variable "val"

	fmt.Println(val)

	// print the value of the package-level variable "val"
	fmt.Println(printGlobal())

	// update the package-level variable "val" and print it again
	updateGlobal("Okay")

fmt.Println(printGlobal())
	fmt.Printf("%T, local &val = %v\n", &val, &val)
	// print the pointer address of the local variable "val"
	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*(&val)  = 24
	fmt.Println(val)
}
