package twiliogo

import (
	"fmt"
	"github.com/eazynow/twilio-go/resources"
	"log"
	"net/url"
)

// add a constant for version

// TwilioRestClient is the main type for the twilio-go adapter. It has references to all
// the other types that the API defines. It is created using the NewTwilioRestClient func
type TwilioRestClient struct {
	Connection     *resources.TwilioConnection
	NumRetries     int
	Notifications  resources.Notifications
	Calls          resources.Calls
	Recordings     resources.Recordings
	Transcriptions resources.Transcriptions
	Queues         resources.Queues
	Usage          resources.Usage
}

// NewTwilioRestClient is used to create an instance of TwilioRestClient using the twilio
// master account sid and token provided. It sets up the TwilioConnection type and sets the pointer
// in all the sub resrources so they share it.
func NewTwilioRestClient(sid, token string) *TwilioRestClient {
	apibase := "https://api.twilio.com"
	apiversion := "2010-04-01"

	retries := 1

	twilioUrl, err := url.Parse(fmt.Sprintf("%s/%s", apibase, apiversion))

	if err != nil {
		log.Fatalf("twilio-go: Bad twilio api version (%s) or endpoint (%s) : %s", apibase, apiversion, err)
	}

	tcon := resources.TwilioConnection{
		Credentials: resources.TwilioAuth{AccountSid: sid, AuthToken: token},
		Endpoint:    twilioUrl.String(),
		NumRetries:  retries,
	}

	return &TwilioRestClient{
		Connection:     &tcon,
		Notifications:  resources.Notifications{Connection: &tcon},
		Calls:          resources.Calls{Connection: &tcon},
		Recordings:     resources.Recordings{Connection: &tcon},
		Transcriptions: resources.Transcriptions{Connection: &tcon},
		Queues:         resources.Queues{Connection: &tcon},
		Usage:          resources.Usage{Records: resources.UsageRecords{Connection: &tcon}, Triggers: resources.UsageTriggers{Connection: &tcon}},
		NumRetries:     retries}
}

// SetRetries is used to set the number of times that the adapter should attempt
// to call twilio
func (trc *TwilioRestClient) SetRetries(retries int) {
	trc.NumRetries = retries
	trc.Connection.NumRetries = retries
	// Connection is a pointer reference so all other structs should get the change
}
