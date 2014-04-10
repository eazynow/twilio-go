package resources

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
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

type CallSetupParams struct {
	Method               string
	FallbackUrl          string
	FallbackMethod       string
	StatusCallback       string
	StatusCallbackMethod string
	SendDigits           string
	IfMachine            string
	Timeout              int
	Record               bool
	SubAccountSid        string
}

func (csp *CallSetupParams) AsValues() url.Values {
	queryVals := url.Values{}

	addParam(&queryVals, "Method", csp.Method)
	addParam(&queryVals, "FallbackUrl", csp.FallbackUrl)
	addParam(&queryVals, "FallbackMethod", csp.FallbackMethod)
	addParam(&queryVals, "StatusCallback", csp.StatusCallback)
	addParam(&queryVals, "StatusCallbackMethod", csp.StatusCallbackMethod)
	addParam(&queryVals, "SendDigits", csp.SendDigits)
	addParam(&queryVals, "IfMachine", csp.IfMachine)
	addParam(&queryVals, "Timeout", strconv.Itoa(csp.Timeout))
	addParam(&queryVals, "Record", strconv.FormatBool(csp.Record))
	addParam(&queryVals, "SubAccountSid", csp.SubAccountSid)

	return queryVals
}

type CallParams struct {
	PagingParams
	SubAccountSid string
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

func getCalls(con *TwilioConnection, params CallParams) (*CallListResponse, error) {
	if len(params.SubAccountSid) == 0 {
		params.SubAccountSid = con.Credentials.AccountSid
	}

	resource := "Calls"

	resp, err := con.Get(params.AsValues(), params.SubAccountSid, resource)

	defer resp.Body.Close()

	if err != nil {
		log.Fatalf("Error getting calls from twilio :%s", err)
	}

	if resp.StatusCode != 200 {
		return nil, convertToTwilioError(resp)
	}

	decoder := json.NewDecoder(resp.Body)

	listResponse := new(CallListResponse)

	err = decoder.Decode(listResponse)

	return listResponse, err
}

func (calls *Calls) GetList(params CallParams) (*CallListResponse, error) {
	return getCalls(calls.Connection, params)
}

func (calls *Calls) Delete(callSid, accountSid string) error {
	// use master account if no sub account selected
	if len(accountSid) == 0 {
		accountSid = calls.Connection.Credentials.AccountSid
	}

	callUrl := fmt.Sprintf("Calls/%s", url.QueryEscape(callSid))

	resp, _ := calls.Connection.Delete(url.Values{}, accountSid, callUrl)

	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return convertToTwilioError(resp)
	}

	return nil
}

func (calls *Calls) Create(from, to string, params CallSetupParams) (*Call, error) {
	if len(params.SubAccountSid) == 0 {
		params.SubAccountSid = calls.Connection.Credentials.AccountSid
	}

	resource := "Calls"

	resp, err := calls.Connection.Post(params.AsValues(), params.SubAccountSid, resource)

	if resp.StatusCode != 201 {
		return nil, convertToTwilioError(resp)
	}

	decoder := json.NewDecoder(resp.Body)

	callResponse := new(Call)

	err = decoder.Decode(callResponse)

	return callResponse, err
}

func (calls *Calls) Update(callSid, method, status string) (*Call, error) {
	return nil, nil
}

func (calls *Calls) GetRecordingList(callSid string, params RecordingParams) (*RecordingListResponse, error) {
	if len(callSid) > 0 {
		// override anything in the params
		params.CallSid = callSid
	}
	return getRecordings(calls.Connection, params)
}

func (calls *Calls) GetNotificationList(callSid string, params NotificationParams) (*NotificationListResponse, error) {
	if len(callSid) > 0 {
		// override anything in the params
		params.CallSid = callSid
	}
	return getNotifications(calls.Connection, params)
}
