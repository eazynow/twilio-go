// twilio primary verbs
package verbs

import (
	"encoding/xml"
	"twilio-go/attribs"
	"twilio-go/nouns"
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
	XMLName xml.Name `xml:"Play"`
	Loop    int      `xml:"loop,attr,omitempty"`
	Url     string   `xml:",chardata"`
}

type Gather struct {
	XMLName xml.Name `xml:"Gather"`
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
	XMLName xml.Name `xml:"Record"`

	// specifies url to call after record is finished
	attribs.ActionAttributes

	Timeout            int    `xml:"timeout,attr,omitempty"`
	FinishOnKey        string `xml:"finishOnKey,attr,omitempty"`
	MaxLength          int    `xml:"maxLength,attr,omitempty"`
	Transcribe         bool   `xml:"transcribe,attr,omitempty"`
	TranscribeCallback string `xml:"transcribeCallback,attr,omitempty"`
	PlayBeep           bool   `xml:"playBeep,attr,omitempty"`
}

type Sms struct {
	XMLName xml.Name `xml:"Sms"`
	To      string   `xml:"to,attr,omitempty"`
	From    string   `xml:"from,attr,omitempty"`
	Text    string   `xml:",chardata"`

	attribs.ActionAttributes

	StatusCallback string `xml:statusCallback,attr,omitempty"`
}

type Dial struct {
	XMLName xml.Name `xml:"Dial"`

	attribs.ActionAttributes
	Timeout      int    `xml:"timeout,attr,omitempty"`
	HangupOnStar bool   `xml:"hangupOnStar,attr,omitempty"`
	TimeLimit    int    `xml:"timeLimit,attr,omitempty"`
	CallerId     string `xml:"callerId,attr,omitempty"`
	Record       bool   `xml:"record,attr,omitempty"`

	NumberToDial string `xml:",chardata"`

	// WhoToDial defines the noun that can be embeded.
	// Valid nouns are Number, Client, Sip, Conference, Queue
	NounToDial interface{}
}

func (d *Dial) SetDialToNumberString(s string) {
	d.NumberToDial = s
	d.NounToDial = nil
}

func (d *Dial) SetDialToNumber(n nouns.Number) {
	d.NumberToDial = ""
	d.NounToDial = n
}

func (d *Dial) SetDialToClient(c nouns.Client) {
	d.NumberToDial = ""
	d.NounToDial = c
}

func (d *Dial) SetDialToSip(s nouns.Sip) {
	d.NumberToDial = ""
	d.NounToDial = s
}

func (d *Dial) SetDialToConference(c nouns.Conference) {
	d.NumberToDial = ""
	d.NounToDial = c
}

func (d *Dial) SetDialToQueue(q nouns.Queue) {
	d.NumberToDial = ""
	d.NounToDial = q
}
