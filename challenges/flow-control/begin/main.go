// challenges/flow-control/begin/main.go
package main
import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panics that might occur anywhere in this function
	//
	defer func() {
		if err := recover(); err != nil{
			fmt.Println("REVOCER RECOVER")
		}
	}()

	// use os package to read the file specified as a command line argument
	//
	bs, err := os.ReadFile(os.Args[1])
	if(err != nil){
		panic("ERROR IN FILE")
	}
	stringBytes := string(bs[:])

	// convert the bytes to a string
	//


	// initialize a map to store the counts
	//
	statsMap := map[string]int{}

	// use the standard library utility package that can help us split the string into words
	//
	words := strings.Split(stringBytes, " ")

	var lenw int = len(words)
	statsMap["words"] = lenw


	// capture the length of the words slice
	//

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	//
	for _, word := range words{
		for _, char := range word{
			if unicode.IsLetter(char){
				statsMap["letters"]++
			}else if unicode.IsNumber(char){
				statsMap["numbers"]++
			}else{
				statsMap["symbols"]++
			}
		}

		

	}

	// dump the map to the console using the spew package
	//
	spew.Dump(statsMap)
}
