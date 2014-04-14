package resources

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

type Transcription struct {
	Resource
	Duration          string `json:"duration"`
	RecordingSid      string `json:"recording_sid"`
	Price             string `json:"price"`
	Status            string `json:"status"`
	TranscriptionText string `json:"transcription_text"`
	Type              string `json:"type"`
}

type TranscriptionListResponse struct {
}

type TranscriptionParams struct {
	PagingParams
	RecordingSid  string
	SubAccountSid string
}

func (params *TranscriptionParams) AsValues() url.Values {
	queryVals := params.PagingParams.AsValues()
	addParam(&queryVals, "RecordingSid", params.RecordingSid)

	return queryVals
}

// getRecordings is a private function that returns notifications based on parameters
func getTranscriptions(con *TwilioConnection, params TranscriptionParams) (*TranscriptionListResponse, error) {
	if len(params.SubAccountSid) == 0 {
		params.SubAccountSid = con.Credentials.AccountSid
	}

	var resource string

	if len(params.RecordingSid) > 0 {
		resource = fmt.Sprintf("Recordings/%s/", url.QueryEscape(params.RecordingSid))
	}

	resource += "Transcriptions"

	resp, err := con.Get(params.AsValues(), params.SubAccountSid, resource)

	defer resp.Body.Close()

	if err != nil {
		log.Fatalf("Error getting recordings from twilio :%s", err)
	}

	if resp.StatusCode != 200 {
		return nil, convertToTwilioError(resp)
	}

	decoder := json.NewDecoder(resp.Body)

	listResponse := new(TranscriptionListResponse)

	err = decoder.Decode(listResponse)

	return listResponse, err
}
