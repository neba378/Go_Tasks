package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
)

// DisplayInterface shows the main menu and handles user input
func DisplayInterface() {
	library := services.NewLibrary() // Use constructor to initialize Library

	// Predefined members
	members := []models.Member{
		{ID: 1, Name: "Abebe", BorrowedBooks: []models.Book{}},
		{ID: 2, Name: "Nati", BorrowedBooks: []models.Book{}},
	}

	for _, member := range members {
		library.AddMember(member)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Exit")
		fmt.Print("Select an option: ")

		scanner.Scan()
		choiceStr := scanner.Text()

		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter book title: ")
			scanner.Scan()
			title := scanner.Text()

			err := library.AddBook(models.Book{Title: title, Status: "Available"})
			if err != nil {
				fmt.Println("Error adding book:", err)
			} else {
				fmt.Println("Book added successfully.")
			}
		case 2:
			fmt.Print("Enter book ID to remove: ")
			scanner.Scan()
			bookIDStr := scanner.Text()

			bookID, err := strconv.Atoi(bookIDStr)
			if err != nil {
				fmt.Println("Invalid ID format.")
				continue
			}

			err = library.RemoveBook(bookID)
			if err != nil {
				fmt.Println("Error removing book:", err)
			} else {
				fmt.Println("Book removed successfully.")
			}
		case 3:
			fmt.Print("Enter member ID: ")
			scanner.Scan()
			memberIDStr := scanner.Text()

			memberID, err := strconv.Atoi(memberIDStr)
			if err != nil {
				fmt.Println("Invalid ID format.")
				continue
			}

			fmt.Print("Enter book ID to borrow: ")
			scanner.Scan()
			bookIDStr := scanner.Text()

			bookID, err := strconv.Atoi(bookIDStr)
			if err != nil {
				fmt.Println("Invalid ID format.")
				continue
			}

			err = library.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error borrowing book:", err)
			} else {
				fmt.Println("Book borrowed successfully.")
			}
		case 4:
			fmt.Print("Enter member ID: ")
			scanner.Scan()
			memberIDStr := scanner.Text()

			memberID, err := strconv.Atoi(memberIDStr)
			if err != nil {
				fmt.Println("Invalid ID format.")
				continue
			}

			fmt.Print("Enter book ID to return: ")
			scanner.Scan()
			bookIDStr := scanner.Text()

			bookID, err := strconv.Atoi(bookIDStr)
			if err != nil {
				fmt.Println("Invalid ID format.")
				continue
			}

			err = library.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error returning book:", err)
			} else {
				fmt.Println("Book returned successfully.")
			}
		case 5:
			books := library.ListAvailableBooks()
			if len(books) == 0 {
				fmt.Println("No available books.")
			} else {
				fmt.Println("Available books:")
				for _, book := range books {
					fmt.Printf("ID: %d, Title: %s\n", book.ID, book.Title)
				}
			}
		case 6:
			fmt.Print("Enter member ID: ")
			scanner.Scan()
			memberIDStr := scanner.Text()

			memberID, err := strconv.Atoi(memberIDStr)
			if err != nil {
				fmt.Println("Invalid ID format.")
				continue
			}

			books := library.ListBorrowedBooks(memberID)
			if len(books) == 0 {
				fmt.Println("No borrowed books.")
			} else {
				fmt.Println("Borrowed books:")
				for _, book := range books {
					fmt.Printf("ID: %d, Title: %s\n", book.ID, book.Title)
				}
			}
		case 7:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
