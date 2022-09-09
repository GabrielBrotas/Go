package main

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

func (h *HtmlElement) String() string {
	return h.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize * indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n",
	  i, e.name))
	if len(e.text) > 0 {
	  sb.WriteString(strings.Repeat(" ",
		indentSize * (indent + 1)))
	  sb.WriteString(e.text)
	  sb.WriteString("\n")
	}
  
	for _, el := range e.elements {
	  sb.WriteString(el.string(indent+1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n",
	  i, e.name))
	return sb.String()
  }
  

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootName,
		HtmlElement{rootName, "", []HtmlElement{}},
	}
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
}

func (b *HtmlBuilder) AddChildFluent(
	childName, childText string) *HtmlBuilder {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
	return b
}

func main() {
	// 1 ------------------
	// using go built in strings Builder
	sb := strings.Builder{}

	// composition of string that will concatenate together by the end
	sb.WriteString("hello ")
	sb.WriteString("world ")
	sb.WriteString("from strings build")

	fmt.Println(sb.String())
	sb.Reset()

	// 2 ------------------

	words := []string{"hello", "world"}
	  // <ul><li>...</li><li>...</li><li>...</li></ul>'
	sb.WriteString("<ul>")
  	
	for _, v := range words {
		// this code is repetitive so we can create a builder for html elements that we know it has open and closed tags
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())

	// 3 ------------------
	b := NewHtmlBuilder("ul")
	
	b.AddChildFluent("li", "hello").
		AddChildFluent("li", "world")
	
	fmt.Println(b.String())
}
