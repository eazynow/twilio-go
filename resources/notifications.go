package resources

type Notification struct {
}

type NotificationSummary struct {
}

type NotificationParams struct {
	PagingParams
	LogLevel   string //log level needs to be string as 0 is a valid level in twilio
	DateAfter  string
	DateBefore string
}

type NotificationListResponse struct {
	ListResponse
	list []NotificationSummary
}

type Notifications struct {
	Name       string
	Connection *TwilioConnection
}

func (nots *Notifications) GetBySid(sid string) Notification {
	n := Notification{}

	return n
}

func (nots *Notifications) GetList(params NotificationParams) NotificationListResponse {
	return NotificationListResponse{}
}
