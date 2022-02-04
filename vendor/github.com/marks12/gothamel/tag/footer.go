package tag

import "encoding/xml"

type Footer struct {
	XMLName   xml.Name      `xml:"Footer"`
	Attributes
	Children
}

