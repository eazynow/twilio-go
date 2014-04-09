// Package resources contains all the key twilio resource objects
package resources

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
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
		return nil, nots.returnError(resp)
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

	queryVals := url.Values{}

	if params.PageSize > 0 {
		queryVals.Add("PageSize", strconv.Itoa(params.PageSize))
	}

	if len(params.Date) > 0 {
		queryVals.Add("MessageDate", params.Date)
	}

	if len(params.DateFrom) > 0 {
		queryVals.Add("MessageDate>", params.DateFrom)
	}

	if len(params.DateTo) > 0 {
		queryVals.Add("MessageDate<", params.DateTo)
	}

	queryVals.Add("Page", strconv.Itoa(params.Page))

	resp, err := nots.Connection.Get(queryVals, params.SubAccountSid, "Notifications")

	defer resp.Body.Close()

	if err != nil {
		log.Fatalf("Error getting notifications from twilio :%s", err)
	}

	if resp.StatusCode != 200 {
		return nil, nots.returnError(resp)
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
		return nots.returnError(resp)
	}

	return nil
}

func (nots *Notifications) returnError(resp *http.Response) error {
	decoder := json.NewDecoder(resp.Body)

	twilioErr := new(TwilioError)

	err := decoder.Decode(twilioErr)
	if err == nil {
		err = twilioErr
	}
	return err
}
