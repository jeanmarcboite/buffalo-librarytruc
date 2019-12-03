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

	return getMeta(isbn, response.Data)
}

// SearchTitle -- search for a work with a title
func SearchTitle(title string) ([]book.Metadata, error) {
	w := normalizeString(title)
	url := fmt.Sprintf(net.Koanf.String("openlibrary.url.title"),
		w)

	data, err := net.HTTPGet(url)
	if err != nil {
		return nil, err
	}
	var response Response
	json.Unmarshal(data, &response)

	if response.NumFound < 1 {
		return nil, fmt.Errorf("No book found for '%s'", title)
	}
	metadata := make([]book.Metadata, 0)

	for _, doc := range response.Docs {
		fmt.Printf("isbns:: %v\n", doc.ISBN)
		if len(doc.ISBN) > 0 {
			mi, err := LookUpISBN(doc.ISBN[0])
			if err != nil {
				return nil, err
			}
			metadata = append(metadata, mi)
		}
	}

	return metadata, err
}

func normalizeString(s string) string {
	w := s
	if idx := strings.IndexAny(s, "(-"); idx >= 0 {
		w = s[:idx]
	}
	return strings.Join(strings.Fields(w), "+")
}

func getMeta(isbn string, response Book) (book.Metadata, error) {
	meta := book.Metadata{
		ISBN:    isbn,
		ID:      response.Details.Key,
		Title:   response.Details.Title,
		Authors: []book.Author{},
		Identifiers: book.Identifiers{
			ISBN10:       response.Details.ISBN10,
			ISBN13:       response.Details.ISBN13,
			Goodreads:    response.Details.Identifiers.Goodreads,
			Librarything: response.Details.Identifiers.Librarything,
		},
		PublishCountry: response.Details.PublishCountry,
		Publishers:     response.Details.Publishers,
		Description:    response.Details.Description,
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
