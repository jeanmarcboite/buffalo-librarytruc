package book

// Metadata -- book metadata
type Metadata struct {
	ID           string
	Title        string
	SubTitle     string
	Authors      []Author
	Rating       string
	ReviewsCount int
	RatingsSum   int
	RatingsCount int
	URL          string
	Covers       []string
	Identifiers  struct {
		ISBN13       []string
		ISBN10       []string
		Amazon       []string
		ASIN         string
		KindleASIN   string
		Google       []string
		Gutenberg    []string
		Goodreads    []string
		Librarything []string
	}
	PublishDate    string
	Publishers     []string
	PublishCountry string
	Description    string
	Subjects       string
	NumberOfPages  int
	Preview        string
	PhysicalFormat string
	IsEbook        string
	LanguageCode   string
	Legal          string
}

// Author
type Author struct {
	Name string
	Key  string
	ID   string
}
