//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Title string
type Name string

type LendAudit struct {
	checkIn  time.Time
	checkOut time.Time
}
type Member struct {
	name  Name
	books map[Title]LendAudit
}
type Books struct {
	checkedOut int
	returned   int
}

type BookEntry struct {
	total  int
	lended int
}

type Library struct {
	member map[Name]Member
	books  map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "not returned yet"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

func printMemberAudits(library *Library) {
	for _, member := range library.member {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	for title, book := range library.books {
		fmt.Println(title, "/ total", book.total, "/ lended:", book.lended)
	}
	fmt.Println()
}

func checkedOutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}
	if book.lended == book.total {
		fmt.Println("No book available to lend")
		return false
	}

	book.lended += 1
	library.books[title] = book

	member.books[title] = LendAudit{checkOut: time.Now()}
	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}

	audit, found := member.books[title]
	if !found {
		fmt.Println("Member didn't check out this book")
		return false
	}

	book.lended -= 1
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

func main() {
	//--Requirements:
	//* The library must have books and members, and must include:
	//  - Which books have been checked out
	//  - What time the books were checked out
	//  - What time the books were returned
	//* Perform the following:
	//  - Add at least 4 books and at least 3 members to the library
	//  - Check out a book
	//  - Check in a book
	//  - Print out initial library information, and after each change
	//* There must only ever be one copy of the library in memory at any time

	library := Library{
		books:  make(map[Title]BookEntry),
		member: make(map[Name]Member),
	}

	library.books["Webapps in Go"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["The little Go"] = BookEntry{
		total:  3,
		lended: 0,
	}
	library.books["Go Bootcamp"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["Learn Go"] = BookEntry{
		total:  2,
		lended: 0,
	}

	library.member["Dora"] = Member{"Dora", make(map[Title]LendAudit)}
	library.member["Nobi"] = Member{"Nobi", make(map[Title]LendAudit)}
	library.member["Gian"] = Member{"Gian", make(map[Title]LendAudit)}

	fmt.Println("\n Initial: ")
	printLibraryBooks(&library)
	printMemberAudits(&library)

	member := library.member["Dora"]
	checkOut := checkedOutBook(&library, "Go Bootcamp", &member)
	fmt.Println("\nCheck out a book")
	if checkOut {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}

	returned := returnBook(&library, "Go Bootcamp", &member)
	fmt.Println("\nCheck in a Book")
	if returned {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}
}
