package resources

// A REST resource
type Resource struct {
}

// The object representation of an instance response of twilio api
type InstanceResource struct {
	Resource
}

type PagingParams struct {
	PageSize int
	Page     int
}

type ListResponse struct {
	PagingParams
	NumPages        int
	Total           int
	Start           int
	End             int
	Uri             string
	FirstPageUri    string
	NextPageUri     string
	PreviousPageUri string
	LastPageUri     string
}
