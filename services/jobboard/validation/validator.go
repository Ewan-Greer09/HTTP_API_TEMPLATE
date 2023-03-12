package validation

import (
	"encoding/json"
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types/jobboard"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Validatior struct {
}

func NewValidator() *Validatior {
	return &Validatior{}
}

// Validation middleware for jobboard post request
func (v *Validatior) ValidateJobBoardPostRequest(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		jobListing := jobboard.JobListing{}
		err := json.NewDecoder(r.Body).Decode(&jobListing)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		err = validation.ValidateStruct(&jobListing,
			validation.Field(&jobListing.Position, validation.Required, is.ASCII),
			validation.Field(&jobListing.Description, validation.Required),
			validation.Field(&jobListing.Location, validation.Required),
			validation.Field(&jobListing.Pay, validation.Required, is.Int),
			validation.Field(&jobListing.Company, validation.Required),
			validation.Field(&jobListing.Salaried, validation.Required),
			validation.Field(&jobListing.Remote, validation.Required),
		)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	})
}
