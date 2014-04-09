package resources

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// A REST resource
type Resource struct {
}

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
		queryVals.Add("PageSize", strconv.Itoa(pp.PageSize))
	}

	queryVals.Add("Page", strconv.Itoa(pp.Page))

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

type TwilioError struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int    `json:"status"`
}

func (te *TwilioError) Error() string {
	return fmt.Sprintf("Twilio Error: (%d) %s", te.Code, te.Message)
}

func convertToTwilioError(resp *http.Response) error {
	decoder := json.NewDecoder(resp.Body)

	twilioErr := new(TwilioError)

	err := decoder.Decode(twilioErr)
	if err == nil {
		err = twilioErr
	}
	return err
}
