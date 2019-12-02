package google

import (
	"encoding/json"
	"fmt"

	"github.com/jeanmarcboite/librarytruc/pkg/books/book"
	"github.com/jeanmarcboite/librarytruc/pkg/books/online/net"
)

// LookUpISBN -- lookup a work on google, with isbn
func LookUpISBN(isbn string) (book.Metadata, error) {
	return get(isbn, net.Koanf.String("google.url.isbn"))
}

func get(isbn string, where string) (book.Metadata, error) {
	url := fmt.Sprintf(where, isbn)
	resp, err := net.HTTPGet(url)
	if err != nil {
		net.Logger.DPanic(url, err)
		return book.Metadata{}, err
	}

	var response Response
	json.Unmarshal([]byte(resp), &response)

	meta, err := getMeta(response)
	meta.ISBN = isbn
	return meta, err
}

func getMeta(response Response) (book.Metadata, error) {
	if response.TotalItems < 1 {
		return book.Metadata{}, nil
	}

	item := response.Items[0]

	authors := make([]book.Author, 0)

	for _, author := range item.VolumeInfo.Authors {
		authors = append(authors, book.Author{
			Name: author,
		})
	}

	return book.Metadata{
		ID:            item.ID,
		Title:         item.VolumeInfo.Title,
		Authors:       authors,
		Publishers:    []string{item.VolumeInfo.Publisher},
		Description:   item.VolumeInfo.Description,
		NumberOfPages: item.VolumeInfo.PageCount,
	}, nil
}
