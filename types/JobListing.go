package types

type JobListing struct {
	JobID          string  `json:"jobId"`
	JobTitle       string  `json:"jobTitle"`
	JobDescription string  `json:"jobDescription"`
	JobLocation    string  `json:"jobLocation"`
	JobSalary      float64 `json:"jobSalary"`
	JobCompany     string  `json:"jobCompany"`
}

func NewJobListing() JobListing {
	return JobListing{}
}

func NewJobListingPerams(jobID, jobTitle, jobDescription, jobLocation, jobCompany string, jobSalary float64) JobListing {
	return JobListing{
		JobID:          jobID,
		JobTitle:       jobTitle,
		JobDescription: jobDescription,
		JobLocation:    jobLocation,
		JobSalary:      jobSalary,
		JobCompany:     jobCompany,
	}
}
