package book

import (
	"fmt"
	"reflect"

	"github.com/jeanmarcboite/bookins/pkg/books/online/net"
)

// Info -- Book info and metadata
type Info struct {
	Metadata
	Online map[string]Metadata
}

// New -- pack Info
func New(metadata map[string]Metadata) (Info, error) {
	this := Info{Online: metadata}

	for online := range metadata {
		for _, what := range []string{"ISBN", "Title", "Authors", "Description", "Identifiers", "NumberOfPages"} {
			assign(&this, online, what)
		}
	}

	this.Cover = fmt.Sprintf(net.Koanf.String("librarything.url.cover"),
		net.Koanf.String("librarything.key"), this.ISBN)

	return this, nil
}

func display(this *Info, key string) {
	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

func assign(this *Info, key string, fieldName string) {
	value := reflect.ValueOf(this.Online[key]).FieldByName(fieldName)
	field := reflect.ValueOf(this).Elem().FieldByName(fieldName)

	if field.IsZero() {
		// A Value can be changed only if it is
		// addressable and was not obtained by
		// the use of unexported struct fields.
		if field.IsValid() && field.CanSet() {
			field.Set(value)
			/**
			if field.Kind() == reflect.String {
				field.SetString(value.String())
			}
			**/
		}
	}
}
