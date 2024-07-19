// challenges/types-composite/begin/main.go
package main
import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	
)
// define an author type with a name
//
type author struct{
	name string
}

// define a book type with a title and author
//
type book struct{
	title string
	author author
}

// define a library type to track a list of books
//
type library map[string][]book

// define addBook to add a book to the library
//
func (l library) addBook(b book){
	l[b.author.name] = append(l[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
//
func (l library) lookupByAuthor( name string) []book{
	books, ok := l[name]
	if ok == true{
		return books

	}else {
		return nil
	} 
}

func main() {
	// create a new library
	//
	newlib := library{}



	// add 2 books to the library by the same author
	//
	b := book{title:"mo dick", author: author{name:"cj"}}
	e := book{title:"mo dick 2", author: author{name:"cj"}}
	g := book{title:"mo dick 3", author: author{name:"cjr"}}

	newlib.addBook(b)
	newlib.addBook(e)
	newlib.addBook(g)

	// add 1 book to the library by a different author
	//

	// dump the library with spew
	//
	spew.Dump(newlib)

	// lookup books by known author in the library
	//
	fmt.Println(newlib.lookupByAuthor("cj"))

	// print out the first book's title and its author's name
	//
	fmt.Println(newlib.lookupByAuthor("cj")[0])
	

}
