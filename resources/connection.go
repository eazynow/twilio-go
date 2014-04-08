package resources

import (
	"fmt"
	"io/ioutil"
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

func (tc *TwilioConnection) callTwilio(method string, formValues url.Values, sid, twilioUrl string) (string, int, error) {
	// req, httperr := http.NewRequest(method, twilioUrl, strings.NewReader(formValues.Encode()))
	fullUrl := fmt.Sprintf("%s/Accounts/%s/%s.json", tc.Endpoint, url.QueryEscape(sid), twilioUrl)

	fmt.Println(fullUrl)
	req, err := http.NewRequest(method, fullUrl, strings.NewReader(formValues.Encode()))

	if err != nil {
		log.Fatalf("twilio-go: error building request: %s", err)
	}

	// use basic auth to connect
	req.SetBasicAuth(tc.Credentials.AccountSid, tc.Credentials.AuthToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	c := &http.Client{}

	//res, err := c.Do(req)
	res, err := c.Do(req)

	if err != nil {
		log.Fatalf("twilio-go: error doing the request: %s", err)
	}

	// body, callerr := ioutil.ReadAll(res.Body)
	response, callerr := ioutil.ReadAll(res.Body)

	fmt.Println(string(response))

	return string(response), res.StatusCode, callerr
}

func (tc *TwilioConnection) Post(formValues url.Values, sid, twilioUrl string) (string, int, error) {
	return tc.callTwilio("POST", formValues, sid, twilioUrl)
}

func (tc *TwilioConnection) Get(formValues url.Values, sid, twilioUrl string) (string, int, error) {
	return tc.callTwilio("GET", formValues, sid, twilioUrl)
}
