package twiliogo

import (
	"fmt"
	"github.com/eazynow/twilio-go/resources"
	"log"
	"net/url"
)

// add a constant for version

type TwilioRestClient struct {
	Connection    *resources.TwilioConnection
	NumRetries    int
	Notifications resources.Notifications
	Calls         resources.Calls
	Recordings    resources.Recordings
}

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
		Connection:    &tcon,
		Notifications: resources.Notifications{Connection: &tcon},
		Calls:         resources.Calls{Connection: &tcon},
		Recordings:    resources.Recordings{Connection: &tcon},
		NumRetries:    retries}
}

func (trc *TwilioRestClient) SetRetries(retries int) {
	trc.NumRetries = retries
	trc.Connection.NumRetries = retries
	// Connection is a pointer reference so all other structs should get the change
}
