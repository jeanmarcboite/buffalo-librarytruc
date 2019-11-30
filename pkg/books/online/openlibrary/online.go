package openlibrary

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jeanmarcboite/librarytruc/pkg/books/book"
	"github.com/jeanmarcboite/librarytruc/pkg/books/online/net"
)

// LookUpISBN -- lookup a work on openlibrary, with isbn
func LookUpISBN(isbn string) (book.Metadata, error) {
	return get(isbn, net.Koanf.String("openlibrary.url.isbn"))
}

func get(isbn string, where string) (book.Metadata, error) {
	url := fmt.Sprintf(where, isbn)
	olresp, err := net.HTTPGet(url)
	resp := strings.Replace(string(olresp), fmt.Sprintf("ISBN:%v", isbn), "data", 1)
	if err != nil {
		net.Logger.DPanic(url, err)
		return book.Metadata{}, err
	}
	//fmt.Printf("%v/n", string(resp))

	var response BookResponse
	json.Unmarshal([]byte(resp), &response)

	return getMeta(response.Data)
}

func getMeta(response Book) (book.Metadata, error) {
	meta := book.Metadata{
		ID:          response.Details.Key,
		Title:       response.Details.Title,
		Authors:     []book.Author{},
		Description: response.Details.Description,
	}
	for _, a := range response.Details.Authors {
		author := book.Author{
			Name: a.Name,
			Key:  a.Key,
		}
		meta.Authors = append(meta.Authors, author)
	}

	return meta, nil
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
