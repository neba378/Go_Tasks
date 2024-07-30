package services

import (
	"fmt"
	"library_management/models"
)

// Library struct to manage books and members
type Library struct {
	books   map[int]models.Book
	members map[int]models.Member
}

// NewLibrary creates a new Library instance with initialized maps
func NewLibrary() *Library {
	return &Library{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}

type LibraryManager interface {
	AddBook(book models.Book) error
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}
//AddMember adds new member and I used it when i create user for the first time
func (l *Library) AddMember(member models.Member) {
	l.members[member.ID] = member
}

// This function adds new book to the library
func (l *Library) AddBook(book models.Book) error {
	// check if there is already created book
	if l.books == nil {
		l.books = make(map[int]models.Book)
	}

	// Check if the book already exists in the library 
	// I used the author and title to check if a book is existed
	for _, existingBook := range l.books {
		if existingBook.Title == book.Title && existingBook.Author == book.Author {
			return fmt.Errorf("book '%s' by '%s' already exists", book.Title, book.Author)
		}
	}

	// Find the next available book ID (iterates until it finds free number)
	bookNum := 1
	for {
		if _, ok := l.books[bookNum]; !ok {
			break
		}
		bookNum++
	}

	// Assign the new ID to the book and add it to the library
	book.ID = bookNum
	l.books[bookNum] = book
	return nil
}

// Used to remove books from library
func (l *Library) RemoveBook(bookID int) error {
    // Check if the book exists in the library
    if _, exists := l.books[bookID]; !exists {
        return fmt.Errorf("book with ID %d does not exist", bookID)
    }

    // Remove the book from the library
    delete(l.books, bookID)
    return nil
}

// this function enables members to borrow book
func (l *Library) BorrowBook(bookID int, memberID int) error {
	member, memberExists := l.members[memberID]
	if !memberExists {
		return fmt.Errorf("member with ID %d does not exist", memberID)
	}

	book, bookExists := l.books[bookID]
	if !bookExists {
		return fmt.Errorf("book with ID %d does not exist", bookID)
	}
	if book.Status == "Borrowed" {
		return fmt.Errorf("book with ID %d is not available right now", bookID)
	}

	// Append the book to the member's BorrowedBooks slice
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	book.Status = "Borrowed"
	l.books[bookID] = book
	// Update the member in the map
	l.members[memberID] = member

	return nil
}

// enables member to return book to the library
func (l *Library) ReturnBook(bookID int, memberID int) error {
	member, memberExists := l.members[memberID]
	if !memberExists {
		return fmt.Errorf("member with ID %d does not exist", memberID)
	}

	book, bookExists := l.books[bookID]
	if !bookExists {
		return fmt.Errorf("book with ID %d does not exist", bookID)
	}
	bookIndex := -1
	for i, borrowedBook := range member.BorrowedBooks {
		if borrowedBook.ID == bookID {
			bookIndex = i
			break
		}
	}

	if bookIndex == -1 {
		return fmt.Errorf("you do not have a book with book ID %d", bookID)
	}

	// Remove the book from the slice
	member.BorrowedBooks = append(member.BorrowedBooks[:bookIndex], member.BorrowedBooks[bookIndex+1:]...)
	book.Status = "Available"
	l.books[bookID] = book
	l.members[memberID] = member

	return nil
}

// displays the id and title to users
func (l *Library) ListAvailableBooks() []models.Book {
	arr := []models.Book{}

	for _, book := range l.books {
		if book.Status == "Available" {
			arr = append(arr, book)
		}
	}
	return arr
}

// shows books that the user has borrowed
func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	arr := []models.Book{}

	member, exists := l.members[memberID]
	if !exists {
		return arr
	}

	arr = append(arr, member.BorrowedBooks...)
	return arr
}
