// twilio secondary verbs
package verbs

/*
Enqueue
Leave
Hangup
Redirect
Reject
Pause
*/

import (
	"encoding/xml"
)

type Enqueue struct {
	XMLName xml.Name `xml:"Enqueue"`

	// A request is made to the action URL when the call
	// leaves the queue, describing the dequeue reason
	// and details about the time spent in the queue
	Action string `xml:"action,attr,omitempty"`
	Method string `xml:"method,attr,omitempty"`

	// The 'waitUrl' attribute specifies a URL pointing to a
	// TwiML document containing TwiML verbs that will be
	// executed while the caller is waiting in the queue.
	WaitUrl       string `xml:"waitUrl,attr,omitempty"`
	WaitUrlMethod string `xml:"waitUrlMethod,attr,omitempty"`

	// the name of the queue
	QueueName string `xml:",chardata"`
}

type Leave struct {
	XMLName xml.Name `xml:"Leave"`
}

type Hangup struct {
	XMLName xml.Name `xml:"Hangup"`
}

type Redirect struct {
	XMLName xml.Name `xml:"Redirect"`
	Url     string   `xml:",chardata"`
	Method  string   `xml:"method,attr,omitempty"`
}

type Reject struct {
	XMLName xml.Name `xml:"Reject"`
	Reason  string   `xml:"reason,attr,omitempty"`
}

type Pause struct {
	XMLName xml.Name `xml:"Pause"`
	Length  int      `xml:"length,attr"`
}
