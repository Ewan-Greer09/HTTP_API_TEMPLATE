package types

type JobListing struct {
	JobID          string  `json:"jobId" Validate:"required"`
	JobTitle       string  `json:"jobTitle" Validate:"required"`
	JobDescription string  `json:"jobDescription" Validate:"required"`
	JobLocation    string  `json:"jobLocation" Validate:"required"`
	JobSalary      float64 `json:"jobSalary" Validate:"required, numeric"`
	JobCompany     string  `json:"jobCompany" Validate:"required"`
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
