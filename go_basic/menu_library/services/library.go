package library

import (
	"fmt"

	"trung.dev/go_lang/db"
	"trung.dev/go_lang/models"
	"trung.dev/go_lang/utils"
)

func AddBook() {
	fmt.Println("=======>Add book for library <=======")
	bookName := utils.ReadInput("Enter book name: ")
	bookAuthor := utils.ReadInput("Enter book author: ")
	bookID := utils.GenerateUUID()

	newBook := models.Books{
		ID:     bookID,
		Name:   bookName,
		Author: bookAuthor,
	}
	err := db.AddBook(newBook)
	if err != nil {
		fmt.Println("❌AddBook to library failed with error:", err)
		return
	}
	fmt.Println("✅AddBook to library success with book ID:", bookID)
}

func ViewAllBooks() {
	fmt.Println("=======>View all books in library <=======")
	db.PrintAllBooks()
}

func AddUserBorrowBook() {
	fmt.Println("=======>Add user want to borrow book <=======")
}

func ViewAllUsersBorrowBook() {
	fmt.Println("=======>View all users want to borrow book <=======")
}

func BorrowBook() {
	fmt.Println("=======>Borrow book <=======")
}

func ReturnBook() {
	fmt.Println("=======>Return book <=======")
}

func SearchInfoBook() {
	fmt.Println("=======>Search info of book <=======")
}

func ViewHistoryBorrowBook() {
	fmt.Println("=======>View history of borrow book <=======")
}
