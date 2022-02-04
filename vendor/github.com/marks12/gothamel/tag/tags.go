package tag

import (
	"encoding/xml"
	"fmt"
	"github.com/marks12/gothamel/constants"
)

type P struct {
	XMLName   	xml.Name    `xml:"P"`
	Children	*[]Nodes	`xml:",omitempty"`
	Text  		string      `xml:",chardata"`
	Attributes
}

type Div struct {
	XMLName   xml.Name      `xml:"Div"`
	Attributes
	Children	*[]Nodes	`xml:",omitempty"`
	Text  string            `xml:",chardata"`
}

type Article struct {
	XMLName   xml.Name      `xml:"Article"`
	Attributes
	Children	*[]Nodes	`xml:",omitempty"`
}


type Body struct {
	XMLName   xml.Name      `xml:"Body"`
	Nodes     *[]Nodes      `xml:"content,omitempty"`
}

type Head struct {
	XMLName   xml.Name      `xml:"Head"`
	Title string
}

type Html struct {
	XMLName     xml.Name     `xml:"Html"`
	Head        Head         `xml:"Head"`
	Body        Body  `xml:"Body"`
}

func (html *Html) GetHtmlPage(isDebug bool) string {

	indent := ""

	if isDebug {
		indent = "    "
	}

	output, err := xml.MarshalIndent(html, "", indent)

	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return constants.HEADER + string(output)
}

func (html *Html) AddHeader(header Header) *Html {
	html.Body.Nodes = &[]Nodes{
		header,
		html.Body.Nodes,
	}
	return html
}

func (html *Html) SetContent(article Article) *Html {
	html.Body.Nodes = &[]Nodes{article}
	return html
}

func (html *Html) AddFooter(footer Footer) *Html {
	html.Body.Nodes = &[]Nodes{
		html.Body.Nodes,
		footer,
	}
	return html
}
