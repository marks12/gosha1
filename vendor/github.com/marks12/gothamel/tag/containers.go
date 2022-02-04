package tag

type Attributes struct {
	Id      string          `xml:"id,attr,omitempty"`
	Name    string          `xml:"name,attr,omitempty"`
}

type Content struct {
	Text  string            `xml:",chardata"`
}

func (cont *Content) SetText(text string) *Content {
	cont.Text = text
	return cont
}

type Children struct {
	Child  *[]Nodes         `xml:"content,omitempty"`
}

type Nodes interface {
}
