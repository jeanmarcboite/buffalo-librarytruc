package book

// Metadata -- book metadata
type Metadata struct {
	ID           string
	Title        string
	SubTitle     string
	Authors      []Author
	Series string
	Tags string
	Ratings       string
	RatingsPercent string
	ReviewsCount int
	RatingsSum   int
	RatingsCount int
	URL          string
	Cover string
	Covers       []string
	Identifiers  Identifiers
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

// Author -- Author name and id
type Author struct {
	Name string
	Key  string
	ID   string
}

// Identifiers -- book identifiers
type Identifiers struct {
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
