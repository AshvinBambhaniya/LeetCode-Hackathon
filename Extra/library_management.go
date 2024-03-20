package main

import "fmt"

type Book struct {
	Title         string
	Author        string
	PublicationYear int
}


type Library struct {
	Books []Book
}

func (lib *Library) AddBook(book Book) {
	lib.Books = append(lib.Books, book)
}

func (lib *Library) RemoveBook(title string) {
	for i, book := range lib.Books {
		if book.Title == title {
			lib.Books = append(lib.Books[:i], lib.Books[i+1:]...)
			fmt.Printf("Book '%s' removed from the library.\n", title)
			return
		}
	}
	fmt.Printf("Book '%s' not found in the library.\n", title)
}

func (lib *Library) DisplayBooks() {
	fmt.Println("Books in the library:")
	for _, book := range lib.Books {
		fmt.Printf("Title: %s, Author: %s, Publication Year: %d\n", book.Title, book.Author, book.PublicationYear)
	}
}

func main() {

	myLibrary := Library{}

	myLibrary.AddBook(Book{"1984", "George Orwell", 1949})
	myLibrary.AddBook(Book{"To Kill a Mockingbird", "Harper Lee", 1960})
	myLibrary.AddBook(Book{"The Great Gatsby", "F. Scott Fitzgerald", 1925})

	myLibrary.DisplayBooks()

	myLibrary.RemoveBook("To Kill a Mockingbird")

	myLibrary.DisplayBooks()
}