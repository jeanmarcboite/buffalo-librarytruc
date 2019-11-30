package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"sort"
	// "github.com/rs/xid"
)

// https://codepen.io/bisserof/pen/fdtBm
type Debug struct {
	data map[string]interface{}
}

func NewDebug(d map[string]interface{}) Debug {
	return Debug{data: d}
}

func (d Debug) HTML(key string) template.HTML {
	if val, ok := d.data[key]; ok {
		return SprintHTML(val)
	}

	return template.HTML("cannot parse object")
}

func json2html(xm []byte) template.HTML {
	var xmu interface{}
	err := json.Unmarshal(xm, &xmu)
	if err != nil {
		return template.HTML(err.Error())
	}

	l := `<div id="jstree">
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
	value := `<li>"%v": "%v"</li>
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

const styleSheet = `
<link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.11.2/css/all.min.css" rel="stylesheet">
<style>
    /* https://makina-corpus.com/blog/metier/2014/construire-un-tree-view-en-css-pur  */
    /* fonctionnel */
    
    input {
        display: none;
    }
    
    input~ul {
        display: none;
    }
    
    input:checked~ul {
        display: block;
    }
    
    input~ol {
        display: none;
    }
    
    input:checked~ol {
        display: block;
    }
    
    input~.tree-node-down {
        display: none;
    }
    
    input:checked~.tree-node-right {
        display: none;
    }
    
    input:checked~.tree-node-down {
        display: inline;
    }
    /* habillage */
    
    li {
        display: block;
        font-family: 'Arial';
        font-size: 15px;
        padding: 0.2em;
        border: 1px solid grey;
    }
    
    li:hover {
        border: 3px solid black;
        border-radius: 3px;
        background-color: lightgrey;
    }
</style>
`

const list = `

<ul>
    <li><input type="checkbox" id="c1" />
        <i class="fa fa-angle-double-right tree-node-right"></i>
        <i class="fa fa-angle-double-down tree-node-down"></i>
        <label for="c1">Dossier A</label>
        <ul>
            <li>Sous dossier A1</li>
            <li>Sous dossier A2</li>
            <li>Sous dossier A3</li>
        </ul>
    </li>
    <li><input type="checkbox" id="c2" />
        <i class="fa fa-angle-double-right tree-node-right"></i>
        <i class="fa fa-angle-double-down tree-node-down"></i>
        <label for="c2">Dossier B</label>
        <ul>
            <li>Sous dossier B1</li>
            <li><input type="checkbox" id="c3" />
                <i class="fa fa-angle-double-right tree-node-right"></i>
                <i class="fa fa-angle-double-down tree-node-down"></i>
                <label for="c3">Sous dossier B2</label>
                <ul>
                    <li>Sous-sous dossier B21</li>
                    <li><input type="checkbox" id="c4" />
                        <i class="fa fa-angle-double-right tree-node-right"></i>
                        <i class="fa fa-angle-double-down tree-node-down"></i>
                        <label for="c4">Sous-sous dossier B22</label>
                        <ul>
                            <li>Sous-sous-sous dossier B221</li>
                            <li>Sous-sous-sous dossier B222</li>
                        </ul>
                    </li>
                    <li>Sous-sous dossier B23</li>
                </ul>
            </li>
        </ul>
    </li>
</ul>
`
