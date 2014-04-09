package resources

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
