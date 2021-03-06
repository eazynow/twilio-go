package resources

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

type Recording struct {
	Resource
	CallSid  string `json:"call_sid"`
	Duration string `json:"duration"`
}

type RecordingListResponse struct {
	ListResponse
	List []Recording `json:"recordings"`
}

// NotificationParams represents the query parameters that can be used
// when calling the Notifications list object
type RecordingParams struct {
	PagingParams
	CallSid       string
	SubAccountSid string
	Date          string
	DateFrom      string
	DateTo        string
}

func (np *RecordingParams) AsValues() url.Values {

	queryVals := np.PagingParams.AsValues()
	addParam(&queryVals, "CallSid", np.CallSid)
	addParam(&queryVals, "DateCreated", np.Date)
	addParam(&queryVals, "DateCreated>", np.DateFrom)
	addParam(&queryVals, "DateCreated<", np.DateTo)

	return queryVals
}

type Recordings struct {
	Name       string
	Connection *TwilioConnection
}

// getRecordings is a private function that returns notifications based on parameters
func getRecordings(con *TwilioConnection, params RecordingParams) (*RecordingListResponse, error) {
	var resource string

	if len(params.CallSid) > 0 {
		resource = fmt.Sprintf("Calls/%s/", url.QueryEscape(params.CallSid))
	}

	resource += "Recordings"

	resp, err := con.Get(params.AsValues(), params.SubAccountSid, resource)

	defer resp.Body.Close()

	if err != nil {
		log.Fatalf("Error getting recordings from twilio :%s", err)
	}

	if resp.StatusCode != 200 {
		return nil, convertToTwilioError(resp)
	}

	decoder := json.NewDecoder(resp.Body)

	listResponse := new(RecordingListResponse)

	err = decoder.Decode(listResponse)

	return listResponse, err
}

func (recs *Recordings) Get(recordingSid, accountSid string) (*Recording, error) {

	recUrl := fmt.Sprintf("Recordings/%s", url.QueryEscape(recordingSid))

	resp, err := recs.Connection.Get(url.Values{}, accountSid, recUrl)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, convertToTwilioError(resp)
	}

	decoder := json.NewDecoder(resp.Body)

	response := new(Recording)

	err = decoder.Decode(response)

	return response, err
}

func (recs *Recordings) GetList(params RecordingParams) (*RecordingListResponse, error) {
	return getRecordings(recs.Connection, params)
}

func (recs *Recordings) Delete(sid string, accountSid string) error {

	theUrl := fmt.Sprintf("Recordings/%s", url.QueryEscape(sid))

	resp, _ := recs.Connection.Delete(url.Values{}, accountSid, theUrl)

	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return convertToTwilioError(resp)
	}

	return nil
}

func (recs *Recordings) GetTranscriptionsList(recordingSid string, params TranscriptionParams) (*TranscriptionListResponse, error) {
	return getTranscriptions(recs.Connection, params)
}
