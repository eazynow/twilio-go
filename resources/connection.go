package resources

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type TwilioAuth struct {
	AccountSid string
	AuthToken  string
}

type TwilioConnection struct {
	Credentials TwilioAuth
	Endpoint    string
	NumRetries  int
}

// callTwilio is a private function that makes a http call to twilio
func (tc *TwilioConnection) callTwilio(method string, formValues url.Values, sid, twilioUrl string) (*http.Response, error) {
	fullUrl := fmt.Sprintf("%s/Accounts/%s/%s.json", tc.Endpoint, url.QueryEscape(sid), twilioUrl)

	req, err := http.NewRequest(method, fullUrl, strings.NewReader(formValues.Encode()))

	if err != nil {
		log.Fatalf("twilio-go: error building request: %s", err)
	}

	// use basic auth to connect
	req.SetBasicAuth(tc.Credentials.AccountSid, tc.Credentials.AuthToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	c := &http.Client{}

	return c.Do(req)
}

// callTwilio is a function to make a POST call to twilio
func (tc *TwilioConnection) Post(formValues url.Values, sid, twilioUrl string) (*http.Response, error) {
	return tc.callTwilio("POST", formValues, sid, twilioUrl)
}

// callTwilio is a function to make a GET call to twilio
func (tc *TwilioConnection) Get(formValues url.Values, sid, twilioUrl string) (*http.Response, error) {
	return tc.callTwilio("GET", formValues, sid, twilioUrl)
}

// callTwilio is a function to make a DELETE call to twilio
func (tc *TwilioConnection) Delete(formValues url.Values, sid, twilioUrl string) (*http.Response, error) {
	return tc.callTwilio("DELETE", formValues, sid, twilioUrl)
}
