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

func (twilio *TwilioRestClient) post(formValues url.Values, twilioUrl string) (body []byte, status int, error) {
	return twilio.callTwilio("POST", formValues, twilioUrl)
}

func (twilio *TwilioRestClient) get(formValues url.Values, twilioUrl string) (body []byte, status int, error) {
	return twilio.callTwilio("GET", formValues, twilioUrl)
}

func (twilio *TwilioRestClient) callTwilio(method string, formValues url.Values, twilioUrl string) ([]byte, int, error) {
	req, httperr := http.NewRequest(method, twilioUrl, strings.NewReader(formValues.Encode()))
	if httperr != nil {
		return nil, httperr
	}

	// use basic auth to connect
	req.SetBasicAuth(twilio.AccountSid, twilio.AuthToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	c := &http.Client{}

	res, err := c.Do(req)

	body, callerr := ioutil.ReadAll(res.Body)

	return body, res.status, callerr
}

func NewTwilioRestClient(accountSid, authToken, endPoint string) {

	trc := &TwilioRestClient{accountSid, authToken, endPoint}
}
