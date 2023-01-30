package types

type Request struct {
	RequestType string `json:"requestType"`
	JobListing  JobListing
}

func NewRequest() Request {
	return Request{}
}

func NewRequestPerams(requestType string, jobListing JobListing) Request {
	return Request{
		RequestType: requestType,
		JobListing:  jobListing,
	}
}
