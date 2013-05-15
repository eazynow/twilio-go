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
	XMLName  xml.Name `xml:"Say"`
	Voice    string   `xml:"voice,attr,omitempty"`
	Language string   `xml:"language,attr,omitempty"`
	Loop     int      `xml:"loop,attr,omitempty"`

	Text string `xml:",chardata"`
}

type Play struct {
	Loop int    `xml:"loop,attr,omitempty"`
	Url  string `xml:",chardata"`
}

type Gather struct {
	// by declaring these as interfaces, they don't get
	// included if empty
	Say   interface{}
	Play  interface{}
	Pause interface{}
}

func (g *Gather) AddSay(s Say) {
	g.Say = s
}

func (g *Gather) RemoveSay() {
	g.Say = nil
}

func (g *Gather) AddPlay(p Play) {
	g.Play = p
}

func (g *Gather) RemovePlay() {
	g.Play = nil
}

func (g *Gather) AddPause(length int) {
	g.Pause = Pause{Length: length}
}

func (g *Gather) RemovePause() {
	g.Play = nil
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
