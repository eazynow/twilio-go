package resources

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Call struct {
	Resource
	ParentCallSid  string `json:"parent_call_sid"`
	To             string `json:"to"`
	ToFormatted    string `json:"to_formatted"`
	From           string `json:"from"`
	FromFormatted  string `json:"from_formatted"`
	PhoneNumberSid string `json:"phone_number_sid"`
	Status         string `json:"status"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
	Duration       string `json:"duration"`
	Price          string `json:"price"`
	PriceUnit      string `json:"price_unit"`
	Direction      string `json:"direction"`
	AnsweredBy     string `json:"answered_by"`
	Annotation     string `json:"annotation"`
	ForwardedFrom  string `json:"forwarded_from"`
	GroupSid       string `json:"group_sid"`
	CallerName     string `json:"caller_name"`
	//TODO: Sort these out - probably easy KV pair of some sort
	//SubResourceUris  string `json:"to"`
	//      "subresource_uris": {
	//        "notifications": "/2010-04-01/Accounts/AC1ab10b47c2bb2d804d3dcee408ddf8ce/Calls/CA70c0075ac7b91a28b6dbbe9fe1dd1e19/Notifications.json",
	//        "recordings": "/2010-04-01/"
}

// func (call *Call) GetNotificationList(params NotificationParams) (*NotificationListResponse, error) {
// 	if len(call.Sid) > 0 {
// 		params.CallSid = call.Sid
// 	}
// 	return getNotifications(call.Connection, params)
// }

// func (calls *Calls) GetRecordingList(params RecordingParams) (*CallRecordingListResponse, error) {
// 	if len(call.Sid) > 0 {
// 		params.CallSid = call.Sid
// 	}
// 	return getRecordings(calls.Connection, params)
// }

// CallListResponse represents the response from twilio for a list of calls
type CallListResponse struct {
	ListResponse
	List []Call `json:"calls"`
}

type CallRecordingListResponse struct {
}

type CallParams struct {
	PagingParams
	From          string
	To            string
	StartTime     string
	ParentCallSid string
	Status        string
}

func (cp *CallParams) AsValues() url.Values {

	queryVals := cp.PagingParams.AsValues()
	addParam(&queryVals, "From", cp.From)
	addParam(&queryVals, "To", cp.To)
	addParam(&queryVals, "StartTime", cp.StartTime)
	addParam(&queryVals, "ParentCallSid", cp.ParentCallSid)
	addParam(&queryVals, "Status", cp.Status)

	return queryVals
}

type Calls struct {
	Name       string
	Connection *TwilioConnection
}

func (calls *Calls) Get(callSid, accountSid string) (*Call, error) {
	// use master account if no sub account selected
	if len(accountSid) == 0 {
		accountSid = calls.Connection.Credentials.AccountSid
	}

	notUrl := fmt.Sprintf("Calls/%s", url.QueryEscape(callSid))

	resp, err := calls.Connection.Get(url.Values{}, accountSid, notUrl)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, convertToTwilioError(resp)
	}

	decoder := json.NewDecoder(resp.Body)

	response := new(Call)

	err = decoder.Decode(response)

	return response, err
}

func (calls *Calls) GetList(params CallParams) (*CallListResponse, error) {
	return nil, nil
}

func (calls *Calls) Delete(callSid, accountSid string) error {
	return nil
}

func (calls *Calls) Create() error {
	return nil
}

func (calls *Calls) Update(callSid, method, status string) (*Call, error) {
	return nil, nil
}

func (calls *Calls) GetRecordingList(callSid, accountSid string) (*CallRecordingListResponse, error) {
	return nil, nil
}

func (calls *Calls) GetNotificationList(callSid string, params NotificationParams) (*NotificationListResponse, error) {
	if len(callSid) > 0 {
		// override anything in the params
		params.CallSid = callSid
	}
	return getNotifications(calls.Connection, params)
}
