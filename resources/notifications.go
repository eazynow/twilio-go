// Package resources contains all the key twilio resource objects
package resources

import (
	"encoding/json"
	"log"
	"net/url"
)

// NotificationBasic represents the core fields relating to a twilio notification.
// This is a subset of the full Notification object to reduce size in list responses.
type NotificationBasic struct {
	Sid           string `json:"sid"`
	AccountSid    string `json:"account_sid"`
	CallSid       string `json:"call_sid"`
	Log           string `json:"log"`
	ErrorCode     string `json:"error_code"`
	MoreInfo      string `json:"more_info"`
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
	DateAfter     string
	DateBefore    string
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

func (nots *Notifications) GetBySid(sid string) Notification {
	n := Notification{}

	return n
}

func (nots *Notifications) GetList(params NotificationParams) (*NotificationListResponse, error) {
	if len(params.SubAccountSid) == 0 {
		params.SubAccountSid = nots.Connection.Credentials.AccountSid
	}

	resp, err := nots.Connection.Get(url.Values{}, params.SubAccountSid, "Notifications")

	defer resp.Body.Close()

	if err != nil {
		log.Fatalf("Error getting notifications from twilio :%s", err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("Expected a 200 from twilio but got a %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)

	listReponse := new(NotificationListResponse)

	err = decoder.Decode(listReponse)

	return listReponse, err
}

func (nots *Notifications) DeleteBySid(sid string) {

}
