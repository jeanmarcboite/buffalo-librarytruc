package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"sort"
	// "github.com/rs/xid"
)

func json2html(xm []byte) template.HTML {
	var xmu interface{}
	err := json.Unmarshal(xm, &xmu)
	if err != nil {
		return template.HTML(err.Error())
	}

	l := `<div id="jsontree">
`
	r := `
    </div>
`
	return template.HTML(l + toHTML(xmu) + r)
	//return template.HTML(styleSheet + list)
}

// JSON2HTML -- convert a JSON to HTML
func JSON2HTML(xm string) template.HTML {
	return json2html([]byte(xm))
}

// SprintHTML  -- print data in HTML
func SprintHTML(x interface{}) template.HTML {
	xm, err := json.Marshal(x)
	if err != nil {
		return template.HTML(err.Error())
	}
	return json2html(xm)
}

func toHTML(x interface{}) string {
	switch v := x.(type) {
	case map[string]interface{}:
		return mapToHTML(v)
	case []interface{}:
		return arrayToHTML(v)
	}

	return fmt.Sprintf("%v", x)
}

func mapToHTML(m map[string]interface{}) string {
	checkbox := `
     <li>%v %v 
     %v
     </li>
    `
	value := `<li>%v: "%v"</li>
    `

	//guid := xid.New().String()

	bufferString := bytes.NewBufferString("<ul>")
	keys := []string{}
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, k := range keys {
		switch m[k].(type) {
		case map[string]interface{}:
			fmt.Fprintf(bufferString, checkbox, k, "{}", toHTML(m[k]))
		case []interface{}:
			fmt.Fprintf(bufferString, checkbox, k, "[]", toHTML(m[k]))
		default:
			fmt.Fprintf(bufferString, value, k, toHTML(m[k]))
		}
	}
	bufferString.WriteString("</ul>")
	return bufferString.String()
}

func arrayToHTML(a []interface{}) string {
	format := `<li>%v</li>
    `

	bufferString := bytes.NewBufferString("<ul>")
	for _, v := range a {
		fmt.Fprintf(bufferString, format, toHTML(v))
	}
	bufferString.WriteString(`</ul>
`)
	return bufferString.String()
}
