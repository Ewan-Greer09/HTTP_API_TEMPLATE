package storage

import (
	"log"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
	"github.com/enescakir/emoji"
)

// *TEMP storage for job listings
var storage = make(map[string]types.JobListing)

type Storage struct {
	storage map[string]types.JobListing
}

func PopulateStorage() map[string]types.JobListing {
	storage = map[string]types.JobListing{
		"123": {
			ID:          "123",
			Position:    "Software Engineer",
			Description: "Write code",
			Location:    "London",
			Pay:         100000,
			Company:     "Google",
		},
		"456": {
			ID:          "456",
			Position:    "Software Engineer",
			Description: "Write code",
			Location:    "London",
			Pay:         100000,
			Company:     "Google",
		},
		"789": {
			ID:          "789",
			Position:    "Software Engineer",
			Description: "Write code",
			Location:    "London",
			Pay:         100000,
			Company:     "Google",
		},
	}
	log.Println(emoji.Rocket, " Populated storage ", emoji.Rocket)
	return storage
}
