package book

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
func New(metadata map[string]Metadata) (Info, error) {
	this := Info{Metadata: metadata}

	return this, nil
}
