package goodreads

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/jeanmarcboite/librarytruc/pkg/books/book"
	"github.com/jeanmarcboite/librarytruc/pkg/books/online/net"
)

// LookUpISBN -- lookup a work on goodreads, with isbn
func LookUpISBN(isbn string) (book.Metadata, error) {
	return get(isbn, net.Koanf.String("goodreads.url.isbn"))
}

// SearchTitle -- search for a work with a title
func SearchTitle(title string) (book.Metadata, error) {
	return get(strings.Join(strings.Fields(title), "+"),
		net.Koanf.String("goodreads.url.title"))
}

func get(what string, where string) (book.Metadata, error) {
	url := fmt.Sprintf(where, what)

	response, err := net.HTTPGetWithKey(url,
		net.Koanf.String("goodreads.keyname"),
		net.Koanf.String("goodreads.key"))
	if err != nil {
		net.Logger.DPanic(url, err)
		return book.Metadata{}, err
	}

	var goodreads Response

	/* response could be: <error>Page not found</error> */
	xml.Unmarshal(response, &goodreads)

	if goodreads.XMLName.Local == "GoodreadsResponse" {
		return getMeta(goodreads.Books[0])
	}

	net.Logger.DPanic(url, err)
	return book.Metadata{}, fmt.Errorf("Nothing found on goodreads for '%v'", what)
}

func getMeta(goodreads Book) (book.Metadata, error) {
	meta := book.Metadata{
		ID:      goodreads.ID,
		Title:   goodreads.Title,
		Authors: []book.Author{},
		Identifiers: book.Identifiers{
			ISBN10:     []string{goodreads.ISBN},
			ISBN13:     []string{goodreads.ISBN13},
			ASIN:       goodreads.ASIN,
			KindleASIN: goodreads.KindleASIN,
		},
		PublishCountry: goodreads.CountryCode,
		Publishers:     []string{goodreads.Publisher},
		Description:    goodreads.Description,
		Cover:          goodreads.ImageURL,
		IsEbook:        goodreads.IsEbook,
		ReviewsCount:   goodreads.Work.ReviewsCount,
		RatingsSum:     goodreads.Work.RatingsSum,
		RatingsCount:   goodreads.Work.RatingsCount,
		Ratings:        goodreads.Work.RatingDist,
	}

	return meta, nil
}
