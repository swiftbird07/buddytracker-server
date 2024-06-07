package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type Buddy struct {
	User       User   `json:"user"`
	StatusText string `json:"statusText"`
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func ListBuddies(w http.ResponseWriter, r *http.Request) {
	data := []Buddy{
		{
			User: User{
				Id:   "1",
				Name: "Ingrid Fowler",
			},
			StatusText: "Offline",
		},
		{
			User: User{
				Id:   "2",
				Name: "Thor Crane",
			},
			StatusText: "Online",
		},
		{
			User: User{
				Id:   "3",
				Name: "Athena Carter",
			},
			StatusText: "Online",
		},
		{
			User: User{
				Id:   "4",
				Name: "Zena Reese",
			},
			StatusText: "Offline",
		},
		{
			User: User{
				Id:   "5",
				Name: "Simon Stokes",
			},
			StatusText: "Offline",
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
