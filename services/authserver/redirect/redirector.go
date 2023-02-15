package redirect

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
	"github.com/carlmjohnson/requests"
)

const (
	url = "http://localhost:8080/api"
)

// TODO: FIX THIS METHOD SO IT APPLIES TO ALL REQUEST TYPES
func RedirectRequestToAPI(request types.JobListing, requestType, id string) error {
	log.Println("Preparing to redirect request to API")

	//convert request to json
	requestJSON, err := json.Marshal(request)
	if err != nil {
		return err
	}

	switch requestType {
	case "ping":
		log.Println("Redirecting ping request to API")
		err := requests.URL("http://localhost:8080/api/ping").
			Method("GET").
			Fetch(nil)

		if err != nil {
			return err
		}
	case "getlistingbyid":
		log.Println("Redirecting getlistingbyid request to API")
		err := requests.
			URL("http://localhost:8080/api/getlistingbyid").
			Method("POST").
			Header("content-type", "application/json").
			BodyJSON(&requestJSON).
			Fetch(nil)

		if err != nil {
			return err
		}

	case "createlistingyid":
		log.Println("Redirecting getlistingbyid request to API")
		err := requests.
			URL("http://localhost:8080/api/createlistingbyid").
			Method("POST").
			Header("content-type", "application/json").
			BodyJSON(&requestJSON).
			Fetch(nil)

		if err != nil {
			return err
		}
	default:
		return errors.New("Invalid request type")
	}

	log.Println("Redirecting request to API")

	return nil
}
