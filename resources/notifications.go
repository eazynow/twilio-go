// Package resources contains all the key twilio resource objects
package resources

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

type NotificationResponse struct {
}

// NotificationBasic represents the core fields relating to a twilio notification.
// This is a subset of the full Notification object to reduce size in list responses.
type NotificationBasic struct {
	NotificationResponse
	Sid           string `json:"sid"`
	AccountSid    string `json:"account_sid"`
	CallSid       string `json:"call_sid"`
	Log           string `json:"log"`
	ErrorCode     string `json:"error_code"`
	MessageText   string `json:"message_text"`
	MessageDate   string `json:"message_date"`
	RequestMethod string `json:"request_method"`
	RequestUrl    string `json:"request_url"`
	DateCreated   string `json:"date_created"`
	ApiVersion    string `json:"api_version"`
	DateUpdated   string `json:"date_updated"`
	Uri           string `json:"uri"`
}

// Notification represents a twilio notification in detail
type Notification struct {
	NotificationBasic
	ResponseBody     string `json:"response_body"`
	RequestVariables string `json:"request_variables"`
	ResponseHeaders  string `json:"response_headers"`
}

// NotificationParams represents the query parameters that can be used
// when calling the Notifications list object
type NotificationParams struct {
	PagingParams
	SubAccountSid string
	LogLevel      string //log level needs to be string as 0 is a valid level in twilio
	Date          string
	DateFrom      string
	DateTo        string
}

func (np *NotificationParams) AsValues() url.Values {

	queryVals := np.PagingParams.AsValues()

	if len(np.Date) > 0 {
		queryVals.Add("MessageDate", np.Date)
	}

	if len(np.DateFrom) > 0 {
		queryVals.Add("MessageDate>", np.DateFrom)
	}

	if len(np.DateTo) > 0 {
		queryVals.Add("MessageDate<", np.DateTo)
	}

	if len(np.LogLevel) > 0 {
		queryVals.Add("Log", np.LogLevel)
	}

	return queryVals
}

// NotificationListResponse represents the response from twilio for a list of notifications
type NotificationListResponse struct {
	ListResponse
	List []NotificationBasic `json:"notifications"`
}

// Notifications represents the notifications resource type in twilio and any actions related to it
type Notifications struct {
	Name       string
	Connection *TwilioConnection
}

func (nots *Notifications) Get(notificationSid, accountSid string) (*Notification, error) {

	// use master account if no sub account selected
	if len(accountSid) == 0 {
		accountSid = nots.Connection.Credentials.AccountSid
	}

	notUrl := fmt.Sprintf("Notifications/%s", url.QueryEscape(notificationSid))

	resp, err := nots.Connection.Get(url.Values{}, accountSid, notUrl)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, convertToTwilioError(resp)
	}

	decoder := json.NewDecoder(resp.Body)

	notResponse := new(Notification)

	err = decoder.Decode(notResponse)

	return notResponse, err
}

func (nots *Notifications) GetList(params NotificationParams) (*NotificationListResponse, error) {
	if len(params.SubAccountSid) == 0 {
		params.SubAccountSid = nots.Connection.Credentials.AccountSid
	}

	resp, err := nots.Connection.Get(params.AsValues(), params.SubAccountSid, "Notifications")

	defer resp.Body.Close()

	if err != nil {
		log.Fatalf("Error getting notifications from twilio :%s", err)
	}

	if resp.StatusCode != 200 {
		return nil, convertToTwilioError(resp)
	}

	decoder := json.NewDecoder(resp.Body)

	listResponse := new(NotificationListResponse)

	err = decoder.Decode(listResponse)

	return listResponse, err
}

func (nots *Notifications) Delete(notificationSid string, accountSid string) error {
	// use master account if no sub account selected
	if len(accountSid) == 0 {
		accountSid = nots.Connection.Credentials.AccountSid
	}

	notUrl := fmt.Sprintf("Notifications/%s", url.QueryEscape(notificationSid))

	resp, _ := nots.Connection.Delete(url.Values{}, accountSid, notUrl)

	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return convertToTwilioError(resp)
	}

	return nil
}
