package resources

type Usage struct {
	Records  UsageRecords
	Triggers UsageTriggers
}

type UsageRecords struct {
	Connection *TwilioConnection
}

type UsageTriggers struct {
	Connection *TwilioConnection
}
