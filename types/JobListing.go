package types

type JobListing struct {
	ID          string            `json:"id"`
	Position    string            `json:"position"`
	Description string            `json:"description"`
	Location    string            `json:"location"`
	Pay         float64           `json:"pay"`
	Company     string            `json:"company"`
	Salaried    bool              `json:"salaried"`
	Remote      bool              `json:"remote"`
	Datafields  map[string]string `json:"datafields"`
}

func NewJobListing() JobListing {
	return JobListing{}
}

func NewJobListingPerams(jobID, jobTitle, jobDescription, jobLocation, jobCompany string, jobSalary float64, isSalaried, isRemote bool, metadata map[string]string) JobListing {
	return JobListing{
		ID:          jobID,
		Position:    jobTitle,
		Description: jobDescription,
		Location:    jobLocation,
		Pay:         jobSalary,
		Company:     jobCompany,
		Salaried:    isSalaried,
		Remote:      isRemote,
		Datafields:  metadata,
	}
}
