package nouns

import (
	"encoding/xml"
	"github.com/eazynow/twilio-go/attribs"
)

type Number struct {
	XMLName xml.Name `xml:"Number"`
}

type Sip struct {
	XMLName  xml.Name `xml:"Sip"`
	Username string   `xml:"username,attr,omitempty"`
	Password string   `xml:"password,attr,omitempty"`
	Uri      string   `xml:",chardata"`

	// specifies a twiml url to run on the called parties end
	// after they answer but before the 2 parties are connected
	attribs.UrlCallAttributes
}

type Client struct {
	XMLName xml.Name `xml:"Client"`
}

type Conference struct {
	XMLName xml.Name `xml:"Conference"`
}

type Queue struct {
	XMLName xml.Name `xml:"Queue"`
}
