package jobboard

type Request struct {
	RequestType string     `json:"requestType"`
	JobListing  JobListing `json:"jobListing"`
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
