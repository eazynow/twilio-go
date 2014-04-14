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

type CallModifyParams struct {
	Url                  string
	Method               string
	Status               string
	FallbackUrl          string
	FallbackMethod       string
	StatusCallback       string
	StatusCallbackMethod string
	SubAccountSid        string
}

func (cmp *CallModifyParams) AsValues() url.Values {
	queryVals := url.Values{}

	addParam(&queryVals, "Url", cmp.Url)
	addParam(&queryVals, "Method", cmp.Method)
	addParam(&queryVals, "Status", cmp.Status)
	addParam(&queryVals, "FallbackUrl", cmp.FallbackUrl)
	addParam(&queryVals, "FallbackMethod", cmp.FallbackMethod)
	addParam(&queryVals, "StatusCallback", cmp.StatusCallback)
	addParam(&queryVals, "StatusCallbackMethod", cmp.StatusCallbackMethod)
	addParam(&queryVals, "SubAccountSid", cmp.SubAccountSid)

	return queryVals
}

type CallSetupParams struct {
	CallModifyParams
	SendDigits string
	IfMachine  string
	Timeout    int
	Record     bool
}

func (csp *CallSetupParams) AsValues() url.Values {
	queryVals := csp.CallModifyParams.AsValues()

	addParam(&queryVals, "SendDigits", csp.SendDigits)
	addParam(&queryVals, "IfMachine", csp.IfMachine)
	addParam(&queryVals, "Timeout", strconv.Itoa(csp.Timeout))
	addParam(&queryVals, "Record", strconv.FormatBool(csp.Record))

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
	callUrl := fmt.Sprintf("Calls/%s", url.QueryEscape(callSid))

	resp, _ := calls.Connection.Delete(url.Values{}, accountSid, callUrl)

	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return convertToTwilioError(resp)
	}

	return nil
}

func (calls *Calls) Create(from, to string, params CallSetupParams) (*Call, error) {
	callParams := params.AsValues()

	addParam(&callParams, "From", from)
	addParam(&callParams, "To", from)

	resource := "Calls"

	resp, err := calls.Connection.Post(callParams, params.SubAccountSid, resource)

	if resp.StatusCode != 201 {
		return nil, convertToTwilioError(resp)
	}

	decoder := json.NewDecoder(resp.Body)

	callResponse := new(Call)

	err = decoder.Decode(callResponse)

	return callResponse, err
}

func (calls *Calls) Update(callSid string, params CallModifyParams) (*Call, error) {

	callParams := params.AsValues()

	resource := fmt.Sprintf("Calls/%s", url.QueryEscape(callSid))

	resp, err := calls.Connection.Post(callParams, params.SubAccountSid, resource)

	if resp.StatusCode != 200 {
		return nil, convertToTwilioError(resp)
	}

	decoder := json.NewDecoder(resp.Body)

	callResponse := new(Call)

	err = decoder.Decode(callResponse)

	return callResponse, err
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
