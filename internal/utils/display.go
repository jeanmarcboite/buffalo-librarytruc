package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"sort"
)

func json2html(xm []byte, checked bool) template.HTML {
	var xmu interface{}
	err := json.Unmarshal(xm, &xmu)
	if err != nil {
		return template.HTML(err.Error())
	}

	return template.HTML(styleSheet + toHTML(xmu, 1, checked))
}

// JSON2HTML -- convert a JSON to HTML
func JSON2HTML(xm string, checked bool) template.HTML {
	return json2html([]byte(xm), checked)
}

// SprintHTML  -- print data in HTML
func SprintHTML(x interface{}, checked bool) template.HTML {
	xm, err := json.Marshal(x)
	if err != nil {
		return template.HTML(err.Error())
	}
	return json2html(xm, checked)
}

func toHTML(x interface{}, id int, checked bool) string {
	switch v := x.(type) {
	case map[string]interface{}:
		return mapToHTML(v, id, checked)
	case []interface{}:
		return arrayToHTML(v, id, checked)
	}

	return fmt.Sprintf("%v", x)
}

func mapToHTML(m map[string]interface{}, id int, checked bool) string {
	checkFlag := ""
	if checked {
		checkFlag = "checked"
	}
	checkbox := `
     <li><input type='checkbox' id='__c%v' %v/>
        <i class='fa fa-angle-double-right'></i>
        <i class='fa fa-angle-double-down'></i>
        <label for='__c%v'>%v %v</label>
        %v
    </li>
    `
	value := `<li>"%v": "%v"</li>`

	bufferString := bytes.NewBufferString("<ul>")
	keys := []string{}
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, k := range keys {
		id++
		//c.Log.Debugf("ID%v: %v\n", id, reflect.TypeOf(v))
		switch m[k].(type) {
		case map[string]interface{}:
			fmt.Fprintf(bufferString, checkbox, id, checkFlag, id, k, "{}", toHTML(m[k], id, checked))
		case []interface{}:
			fmt.Fprintf(bufferString, checkbox, id, checkFlag, id, k, "[]", toHTML(m[k], id, checked))
		default:
			fmt.Fprintf(bufferString, value, k, toHTML(m[k], id, checked))
		}
	}
	bufferString.WriteString("</ul>")
	return bufferString.String()
}

func arrayToHTML(a []interface{}, id int, checked bool) string {
	format := `<li>%v</li>`

	bufferString := bytes.NewBufferString("<ol  type = 'I'>")
	for _, v := range a {
		id++
		fmt.Fprintf(bufferString, format, toHTML(v, id, checked))
	}
	bufferString.WriteString("</ol>")
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
    
    input~.fa-angle-double-down {
        display: none;
    }
    
    input:checked~.fa-angle-double-right {
        display: none;
    }
    
    input:checked~.fa-angle-double-down {
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
        <i class="fa fa-angle-double-right"></i>
        <i class="fa fa-angle-double-down"></i>
        <label for="c1">Dossier A</label>
        <ul>
            <li>Sous dossier A1</li>
            <li>Sous dossier A2</li>
            <li>Sous dossier A3</li>
        </ul>
    </li>
    <li><input type="checkbox" id="c2" />
        <i class="fa fa-angle-double-right"></i>
        <i class="fa fa-angle-double-down"></i>
        <label for="c2">Dossier B</label>
        <ul>
            <li>Sous dossier B1</li>
            <li><input type="checkbox" id="c3" />
                <i class="fa fa-angle-double-right"></i>
                <i class="fa fa-angle-double-down"></i>
                <label for="c3">Sous dossier B2</label>
                <ul>
                    <li>Sous-sous dossier B21</li>
                    <li><input type="checkbox" id="c4" />
                        <i class="fa fa-angle-double-right"></i>
                        <i class="fa fa-angle-double-down"></i>
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
