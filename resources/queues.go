package resources

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

type Queue struct {
	Resource
	FriendlyName    string `json:friendly_name"`
	CurrentSize     int    `json:"current_size"`
	MaxSize         int    `json:"max_size"`
	AverageWaitTime int    `json:"average_wait_time"`
}

type QueueListResponse struct {
	ListResponse
	List []Queue `json:"queues"`
}

type QueueParams struct {
	PagingParams
	SubAccountSid string
}

func (q *QueueParams) name() url.Values {
	return q.PagingParams.AsValues()
}

type Queues struct {
	Name       string
	Connection *TwilioConnection
}

func (res *Queues) GetList(params QueueParams) (*QueueListResponse, error) {
	if len(params.SubAccountSid) == 0 {
		params.SubAccountSid = res.Connection.Credentials.AccountSid
	}

	resource := "Queues"

	resp, err := res.Connection.Get(params.AsValues(), params.SubAccountSid, resource)

	defer resp.Body.Close()

	if err != nil {
		log.Fatalf("Error getting queues from twilio :%s", err)
	}

	if resp.StatusCode != 200 {
		return nil, convertToTwilioError(resp)
	}

	decoder := json.NewDecoder(resp.Body)

	listResponse := new(QueueListResponse)

	err = decoder.Decode(listResponse)

	return listResponse, err
}

func (res *Queues) Get(queueSid, accountSid string) (*Queue, error) {
	// use master account if no sub account selected
	if len(accountSid) == 0 {
		accountSid = res.Connection.Credentials.AccountSid
	}

	theUrl := fmt.Sprintf("Queues/%s", url.QueryEscape(queueSid))

	resp, err := res.Connection.Get(url.Values{}, accountSid, theUrl)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, convertToTwilioError(resp)
	}

	decoder := json.NewDecoder(resp.Body)

	response := new(Queue)

	err = decoder.Decode(response)

	return response, err
}
