package openlibrary

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jeanmarcboite/librarytruc/pkg/books/online/net"
)

// LookUpISBN -- lookup a work on openlibrary, with isbn
func LookUpISBN(isbn string) (Book, error) {
	url := fmt.Sprintf(net.Koanf.String("openlibrary.url.isbn"), isbn)
	olResponse, err := net.HTTPGet(url)
	response := strings.Replace(string(olResponse), fmt.Sprintf("ISBN:%v", isbn), "data", 1)

	if err != nil {
		return Book{}, err
	}

	var book BookResponse
	json.Unmarshal([]byte(response), &book)

	book.Data.Cover = fmt.Sprintf(net.Koanf.String("openlibrary.url.cover"), "isbn", isbn)
	return book.Data, nil
}

// Search -- search for a work with a title
func Search(title string, author string) (Response, error) {
	w := title
	if idx := strings.IndexAny(w, "(-"); idx >= 0 {
		w = w[:idx]
	}
	plusWords := strings.Join(strings.Fields(w), "+")

	var url string
	if len(author) <= 0 {
		url = fmt.Sprintf(net.Koanf.String("openlibrary.url.title"), plusWords)
	} else {
		plusAuthor := strings.Join(strings.Fields(author), "+")
		url = fmt.Sprintf(net.Koanf.String("openlibrary.url.titleauthor"), plusWords, plusAuthor)
	}

	data, err := net.HTTPGet(url)
	if err != nil {
		return Response{}, err
	}

	var response Response
	json.Unmarshal(data, &response)

	for _, doc := range response.Docs {
		s, _ := json.MarshalIndent(doc, "", "\t")
		net.Logger.Debugf("openlibrary: %s\n", string(s))
	}

	return response, nil

}
