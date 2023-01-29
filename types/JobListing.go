package types

type JobListing struct {
	JobID          string  `json:"jobId" validate:"required"`
	JobTitle       string  `json:"jobTitle" validate:"required"`
	JobDescription string  `json:"jobDescription" validate:"required"`
	JobLocation    string  `json:"jobLocation" validate:"required"`
	JobSalary      float64 `json:"jobSalary" validate:"required, numeric"`
	JobCompany     string  `json:"jobCompany" validate:"required"`
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
