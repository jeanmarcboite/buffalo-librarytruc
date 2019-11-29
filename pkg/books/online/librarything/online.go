package librarything

import (
	"encoding/xml"
	"fmt"
	"strings"

	xml2json "github.com/basgys/goxml2json"
	"github.com/jeanmarcboite/librarytruc/pkg/books/book"
	"github.com/jeanmarcboite/librarytruc/pkg/books/online/net"
)

// LookUpISBN -- lookup a Book on goodreads, with isbn
func LookUpISBN(isbn string) (book.Metadata, error) {
	return get(isbn, net.Koanf.String("librarything.url.isbn"))
}

func getMeta(response Response) (book.Metadata, error) {
	author := book.Author{
		Name: response.Ltml.Item.Author.Text,
		ID:   response.Ltml.Item.Author.ID,
		Key:  response.Ltml.Item.Author.Authorcode,
	}

	return book.Metadata{
		ID:      response.Ltml.Item.ID,
		Title:   response.Ltml.Item.Title,
		Authors: []book.Author{author},
	}, nil
}

func get(what string, where string) (book.Metadata, error) {
	url := fmt.Sprintf(where, what)

	resp, err := net.HTTPGetWithKey(url,
		net.Koanf.String("librarything.keyname"),
		net.Koanf.String("librarything.key"))
	if err != nil {
		return book.Metadata{}, err
	}

	var response Response

	/* Book could be: <error>Page not found</error> */
	xml.Unmarshal(resp, &response)

	if response.XMLName.Local == "response" {
		if response.Stat == "fail" {
			xml := strings.NewReader(string(resp))
			json, _ := xml2json.Convert(xml)

			return book.Metadata{}, fmt.Errorf("Librarything response error: %v", json)
		}

		return getMeta(response)
	}

	return book.Metadata{}, fmt.Errorf("LibraryThing for '%v': %v", what, response)
}
