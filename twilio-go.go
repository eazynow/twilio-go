package twiliogo

import (
	"io/ioutil"
)

// add a constant for version

type TwilioRestClient struct {
	AccountSid          string
	AuthTokenredentials string
	Endpoint            string
	DefaultVersion      string
	NumRetries          int
}

func (twilio *Twilio) post(formValues url.Values, twilioUrl string) (*http.Response, error) {
	return twilio.callTwilio("POST", formValues, twilioUrl)
}

func (twilio *Twilio) get(formValues url.Values, twilioUrl string) (*http.Response, error) {
	return twilio.callTwilio("GET", formValues, twilioUrl)
}

func (twilio *Twilio) callTwilio(method string, formValues url.Values, twilioUrl string) (*http.Response, error) {
	req, err := http.NewRequest(method, twilioUrl, strings.NewReader(formValues.Encode()))
	if err != nil {
		return nil, err
	}

	// use basic auth to connect
	req.SetBasicAuth(twilio.AccountSid, twilio.AuthToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	c := &http.Client{}
	return c.Do(req)
}

func NewTwilioRestClient(accountSid, authToken, endPoint string) {

	trc := &TwilioRestClient{accountSid, authToken, endPoint}
}
