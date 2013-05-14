package twiliogo

// add a constant for version

type TwilioRestClient struct {
	AccountSid          string
	AuthTokenredentials string
	Endpoint            string
	DefaultVersion      string
	NumRetries          int
}

func (twilio *Twilio) post(formValues url.Values, twilioUrl string) (*http.Response, error) {
	req, err := http.NewRequest("POST", twilioUrl, strings.NewReader(formValues.Encode()))
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
