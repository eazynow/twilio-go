package twiliogo

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// add a constant for version

type TwilioRestClient struct {
	AccountSid     string
	AuthToken      string
	Endpoint       string
	DefaultVersion string
	NumRetries     int
}

func (twilio *TwilioRestClient) callTwilio(method string, formValues url.Values, twilioUrl string) (string, int, error) {
	// req, httperr := http.NewRequest(method, twilioUrl, strings.NewReader(formValues.Encode()))
	req, _ := http.NewRequest(method, twilioUrl, strings.NewReader(formValues.Encode()))

	// use basic auth to connect
	req.SetBasicAuth(twilio.AccountSid, twilio.AuthToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	c := &http.Client{}

	//res, err := c.Do(req)
	res, _ := c.Do(req)

	// body, callerr := ioutil.ReadAll(res.Body)
	_, callerr := ioutil.ReadAll(res.Body)

	return "body", res.StatusCode, callerr
}

func (twilio *TwilioRestClient) Post(formValues url.Values, twilioUrl string) (string, int, error) {
	return twilio.callTwilio("POST", formValues, twilioUrl)
}

func (twilio *TwilioRestClient) Get(formValues url.Values, twilioUrl string) (string, int, error) {
	return twilio.callTwilio("GET", formValues, twilioUrl)
}

func NewTwilioRestClient(accountSid, authToken, endPoint string) *TwilioRestClient {
	twilioUrl := "https://api.twilio.com/2010-04-01"
	retries := 1
	return &TwilioRestClient{
		AccountSid:     accountSid,
		AuthToken:      authToken,
		Endpoint:       endPoint,
		DefaultVersion: twilioUrl,
		NumRetries:     retries}
}
