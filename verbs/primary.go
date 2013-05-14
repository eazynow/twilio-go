// twilio primary verbs
package verbs

import (
	"encoding/xml"
)

/*
Say
Play
Gather
Record
Sms
Dial
    Number
    Sip
    Client
    Conference
    Queue
*/

type Say struct {
	Voice    string `xml:"voice,attr,omitempty"`
	Language string `xml:"language,attr,omitempty"`
	Loop     int    `xml:"loop,attr,omitempty"`
	Text     string `xml:",chardata"`
}

type Play struct {
}

type Gather struct {
}

type Record struct {
}

type Sms struct {
}

type Dial struct {
}

type Number struct {
}

type Sip struct {
	XMLName  xml.Name `xml:"Sip"`
	Username string   `xml:"username,attr,omitempty"`
	Password string   `xml:"password,attr,omitempty"`
	Uri      string   `xml:",chardata"`

	// specifies a twiml url to run on the called parties end
	// after they answer but before the 2 parties are connected
	UrlCallAttributes
}

type Client struct {
}

type Conference struct {
}

type Queue struct {
}
