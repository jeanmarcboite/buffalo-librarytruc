package book

import (
	"fmt"
	"github.com/jeanmarcboite/bookins/pkg/books/online/net"
)

// Info -- Book info and metadata
type Info struct {
	Title          string
	SubTitle       string
	Authors        []string
	Identifiers    Identifiers
	URL            map[string]string
	Cover          string
	Covers         map[string]string
	Subjects       []interface{}
	Ratings        string
	RatingsPercent string

	Inforigin   []interface{}
	Description string
	Metadata    map[string]Metadata
}

// Identifiers -- book identifiers
type Identifiers struct {
	ISBN10       string
	ISBN13       string
	LCCN         string
	Openlibrary  string
	Goodreads    string
	Librarything string
}

// New -- pack Info
func New(ISBN string, metadata map[string]Metadata) (Info, error) {
	this := Info{Metadata: metadata}

	this.Title = metadata["librarything"].Title

	this.Cover = fmt.Sprintf(net.Koanf.String("librarything.url.cover"),
		net.Koanf.String("librarything.key"), ISBN)

	return this, nil
}
