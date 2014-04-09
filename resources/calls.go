package resources

import (
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

// CallListResponse represents the response from twilio for a list of calls
type CallListResponse struct {
	ListResponse
	List []Call `json:"calls"`
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
