package attribs

type UrlCallAttributes struct {
	Url    string `xml:"url,attr,omitempty"`
	Method string `xml:"method,attr,omitempty"`
}

type ActionAttributes struct {
	Action string `xml:"action,attr,omitempty"`
	Method string `xml:"method,attr,omitempty"`
}
