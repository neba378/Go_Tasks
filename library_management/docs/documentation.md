Library Management System Documentation

Introduction
The Library Management System is a command-line application designed to manage books and members in a library. It supports operations such as adding and removing books, borrowing and returning books, and listing available and borrowed books.

System Requirements
- Go 1.22 or higher
- A terminal or command-line interface

Running the Application
1. Navigate to the `main.go` file in your terminal or command prompt.
2. Run the application using the following command:
```go run main.go```

This will start the application and display the main menu in your terminal.

Functionalities
1. **Add Book**
   - Description: Adds a new book to the library.
   - Input: Book title (string)
   - Example: `Enter book title: Oromay`

2. **Remove Book**
   - Description: Removes a book from the library by its ID.
   - Input: Book ID (integer)
   - Example: `Enter book ID to remove: 1`

3. **Borrow Book**
   - Description: Allows a member to borrow a book if it is available.
   - Inputs: Member ID (integer) and Book ID (integer)
   - Example:
     ```
     Enter member ID: 1
     Enter book ID to borrow: 2
     ```

4. **Return Book**
   - Description: Allows a member to return a borrowed book.
   - Inputs: Member ID (integer) and Book ID (integer)
   - Example:
     ```
     Enter member ID: 1
     Enter book ID to return: 2
     ```

5. **List Available Books**
   - Description: Lists all available books in the library.
   - Output: List of books with their IDs, titles, and statuses.

6. **List Borrowed Books**
   - Description: Lists all books borrowed by a specific member.
   - Input: Member ID (integer)
   - Output: List of borrowed books with their IDs and titles.

7. **Exit**
   - Description: Exits the application.

Error Handling
The system handles errors such as:
- Invalid book or member IDs
- Attempting to borrow a book that is not available
- Attempting to return a book that is not borrowed by the member

Errors are reported with descriptive messages to guide the user in correcting the issue.

Code Structure
- `main.go`: The entry point of the application. Initializes the Library and starts the user interface.
- `services/library_services.go`: Contains the Library struct and methods for managing books and members.
- `models/book.go`: Contains the Book struct.
- `models/member.go`: Contains the Member struct.
- `controllers/library_controller.go`: Contains the `DisplayInterface` function that handles user interaction.
