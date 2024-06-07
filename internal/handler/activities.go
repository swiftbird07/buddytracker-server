package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type DefaultActivity struct {
	Id   string            `json:"id"`
	Text map[string]string `json:"text"` // Key is lang code
}

func ListActivities(w http.ResponseWriter, r *http.Request) {
	data := []DefaultActivity{
		{
			Id: "studying",
			Text: map[string]string{
				"en": "Studying",
				"de": "Lernen",
			},
		},
		{
			Id: "swimming",
			Text: map[string]string{
				"en": "Swimming",
				"de": "Schwimmen",
			},
		},
	}

	dataB, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(dataB)
}
