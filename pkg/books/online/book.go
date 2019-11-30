package online

import (
	"github.com/jeanmarcboite/librarytruc/pkg/books/book"
	"github.com/jeanmarcboite/librarytruc/pkg/books/online/goodreads"
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

	return book.New(isbn, metadata)
}
