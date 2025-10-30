package db

import (
	"fmt"

	"trung.dev/go_lang/models"
)

type Library struct {
	books map[string]models.Books
}

var LibraryDB = Library{}

func AddBook(book models.Books) error {
	if _, exists := LibraryDB.books[book.ID]; exists {
		return fmt.Errorf("book ID %s already exists", book.ID)

	}
	LibraryDB.books[book.ID] = book
	return nil
}
func PrintAllBooks() {
	listBook := LibraryDB.books
	for _, book := range listBook {
		fmt.Printf("Book name: %s, Book author: %s, Book ID: %s\n", book.Name, book.Author, book.ID)
	}
}
