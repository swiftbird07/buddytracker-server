package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type Buddy struct {
	User User `json:"user"`
}

type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status Status `json:"status"`
}

type Status struct {
	Status           string     `json:"status"`
	Activities       []Activity `json:"activities"`
	Location         Location   `json:"location"`
	LocationAccuracy int        `json:"locationAccuracy"`
	ExpiresAt        int64      `json:"expiresAt"`
}

type Activity struct {
	Custom   bool   `json:"custom"`
	Activity string `json:"activity"`
}

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

func ListBuddies(w http.ResponseWriter, r *http.Request) {
	data := []Buddy{
		{
			User: User{
				Id:   "1",
				Name: "Bob Smith",
				Status: Status{
					Status: "Open for Activities",
					Activities: []Activity{{
						Custom:   true,
						Activity: "Going to the Beach",
					},
						{
							Custom:   true,
							Activity: "Clubbing",
						},
					},
					Location: Location{
						Latitude:  54.118751,
						Longitude: 12.201846,
					},
					LocationAccuracy: 0,
					ExpiresAt:        1723049469,
				},
			},
		},
		{
			User: User{
				Id:   "2",
				Name: "Charlie Miles",
				Status: Status{
					Status: "Join me!",
					Activities: []Activity{
						{
							Custom:   false,
							Activity: "studying",
						},
					},
					Location: Location{
						Latitude:  49.405880,
						Longitude: 8.688515,
					},
					LocationAccuracy: 0,
					ExpiresAt:        1720371069,
				},
			},
		},
		{
			User: User{
				Id:   "3",
				Name: "Alice Johnson",
				Status: Status{
					Status:     "Open for Activities",
					Activities: []Activity{},
					Location: Location{
						Latitude:  51.360806,
						Longitude: 13.024071,
					},
					LocationAccuracy: 0,
					ExpiresAt:        1720716669,
				},
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
