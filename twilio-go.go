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
}

func NewTwilioRestClient(sid, token string) *TwilioRestClient {
	apibase := "https://api.twilio.com"
	apiversion := "2010-04-01"

	retries := 1

	twilioUrl, err := url.Parse(fmt.Sprintf("%s/%s", apibase, apiversion))

	if err != nil {
		log.Fatalf("twilio-go: Bad twilio api version (%s) or endpoint (%s) : %s", apibase, apiversion, err)
	}

	tcred := resources.TwilioAuth{
		AccountSid: sid,
		AuthToken:  token,
	}

	tcon := resources.TwilioConnection{
		Credentials: tcred,
		Endpoint:    twilioUrl.String(),
		NumRetries:  retries,
	}

	tnot := resources.Notifications{
		Connection: &tcon,
	}

	return &TwilioRestClient{
		Connection:    &tcon,
		Notifications: tnot,
		NumRetries:    retries}
}

func (trc *TwilioRestClient) SetRetries(retries int) {
	trc.NumRetries = retries
	trc.Connection.NumRetries = retries
	fmt.Println("trcConnRetries :", retries)

	fmt.Println("NotificationRetries :", trc.Notifications.Connection.NumRetries)
}
