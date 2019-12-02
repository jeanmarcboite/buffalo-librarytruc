package online

import (
	"github.com/jeanmarcboite/librarytruc/pkg/books/book"
	"github.com/jeanmarcboite/librarytruc/pkg/books/online/goodreads"
	"github.com/jeanmarcboite/librarytruc/pkg/books/online/google"
	"github.com/jeanmarcboite/librarytruc/pkg/books/online/librarything"
	"github.com/jeanmarcboite/librarytruc/pkg/books/online/openlibrary"
)

// LookUpISBN -- lookup a work on goodreads and openlibrary, with isbn
func LookUpISBN(isbn string) (book.Info, error) {
	metadata := make(map[string]book.Metadata)
	l, err := librarything.LookUpISBN(isbn)

	if err != nil {
		return book.Info{}, err
	}
	metadata["librarything"] = l
	o, err := openlibrary.LookUpISBN(isbn)
	if err == nil {
		metadata["openlibrary"] = o
	}

	g, err := goodreads.LookUpISBN(isbn)
	if err == nil {
		metadata["goodreads"] = g
	}

	goog, err := google.LookUpISBN(isbn)
	if err == nil {
		metadata["google"] = goog
	}

	return book.New(metadata)
}

// SearchTitle --
func SearchTitle(title string) (book.Info, error) {
	metadata := make(map[string]book.Metadata)
	o, err := openlibrary.SearchTitle(title)
	if err == nil {
		metadata["openlibrary"] = o
	}
	return book.New(metadata)

}
