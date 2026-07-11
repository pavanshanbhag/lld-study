package main

import (
	"fmt"

	"iterator"
)

func main() {
	collection := iterator.NewBookCollection()

	collection.AddBook(iterator.NewBook("The Great Gatsby", "F. Scott Fitzgerald"))
	collection.AddBook(iterator.NewBook("To Kill a Mockingbird", "Harper Lee"))
	collection.AddBook(iterator.NewBook("1984", "George Orwell"))

	it := collection.CreateIterator()

	fmt.Println("Iterating through the book collection:")
	for it.HasNext() {
		book := it.Next().(*iterator.Book)
		fmt.Printf("Book: %s by %s\n", book.Title, book.Author)
	}
}
