package tag

import "encoding/xml"

type Header struct {
	XMLName   xml.Name      `xml:"Header"`
	Attributes
	Children
}

