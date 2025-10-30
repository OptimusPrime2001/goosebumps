package main

import (
	"fmt"

	library "trung.dev/go_lang/services"
	"trung.dev/go_lang/utils"
)

func main() {
	for {
		utils.ClearScreen()
		fmt.Println("📗📗📗📗 CHUONG TRINH QUAN LY THU VIEN 📗📗📗📗")
		fmt.Println("1. Add book for library")
		fmt.Println("2. View all books in library")
		fmt.Println("3. Add user want to borrow book")
		fmt.Println("4. View all users want to borrow book")
		fmt.Println("5. Borrow book")
		fmt.Println("6. View history of borrow book")
		fmt.Println("7. Return book")
		fmt.Println("8. Search info of book")
		fmt.Println("9. Exit")

		choice := utils.GetPostiveInt("Enter your choice: ")
		utils.ClearScreen()
		switch choice {
		case 1:
			library.AddBook()
		case 2:
			library.ViewAllBooks()
		case 3:
			library.AddUserBorrowBook()
		case 4:
			library.ViewAllUsersBorrowBook()
		case 5:
			library.BorrowBook()
		case 6:
			library.ViewHistoryBorrowBook()
		case 7:
			library.ReturnBook()
		case 8:
			library.SearchInfoBook()
		case 9:
			fmt.Println("Exit")
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 9.")
		}

		utils.ReadInput("Press Enter to continue...")
	}
}
