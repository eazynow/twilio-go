package resources

import (
	"net/url"
	"strconv"
)

// The object representation of an instance response of twilio api
type InstanceResource struct {
	Resource
}

type PagingParams struct {
	PageSize int `json:"page_size"`
	Page     int `json:"page"`
}

func (pp *PagingParams) AsValues() url.Values {
	queryVals := url.Values{}

	if pp.PageSize > 0 {
		addParam(&queryVals, "PageSize", strconv.Itoa(pp.PageSize))
	}

	addParam(&queryVals, "Page", strconv.Itoa(pp.Page))

	return queryVals
}

type ListResponse struct {
	PagingParams
	NumPages        int    `json:"num_pages"`
	Total           int    `json:"total"`
	Start           int    `json:"start"`
	End             int    `json:"end"`
	Uri             string `json:"uri"`
	FirstPageUri    string `json:"first_page_uri"`
	NextPageUri     string `json:"next_page_uri"`
	PreviousPageUri string `json:"previous_page_uri"`
	LastPageUri     string `json:"last_page_uri"`
}

// A REST resource
type Resource struct {
	Sid         string `json:"sid"`
	AccountSid  string `json:"account_sid"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
	Uri         string `json:"uri"`
	ApiVersion  string `json:"api_version"`
}

func addParam(params *url.Values, pName, pValue string) {
	if len(pValue) > 0 {
		params.Add(pName, pValue)
	}
}
