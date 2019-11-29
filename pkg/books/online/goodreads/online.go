package goodreads

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/jeanmarcboite/librarytruc/pkg/books/online/net"
)

// LookUpISBN -- lookup a work on goodreads, with isbn
func LookUpISBN(isbn string) (Response, error) {
	return get(isbn, net.Koanf.String("goodreads.url.isbn"))
}

// SearchTitle -- search for a work with a title
func SearchTitle(title string) (Response, error) {
	return get(strings.Join(strings.Fields(title), "+"),
		net.Koanf.String("goodreads.url.title"))
}

func get(what string, where string) (Response, error) {
	url := fmt.Sprintf(where, what)

	response, err := net.HTTPGetWithKey(url,
		net.Koanf.String("goodreads.keyname"),
		net.Koanf.String("goodreads.key"))
	if err != nil {
		return Response{}, err
	}

	var goodreads Response

	/* response could be: <error>Page not found</error> */
	xml.Unmarshal(response, &goodreads)

	if goodreads.XMLName.Local == "GoodreadsResponse" {
		return goodreads, nil
	}

	return goodreads, fmt.Errorf("Nothing found on goodreads for '%v'", what)
}
