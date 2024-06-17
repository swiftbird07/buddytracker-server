package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/swiftbird07/buddytracker-server/internal/controller"
)

type ReqUDIDRegister struct {
	UDID string `json:"udid"`
	Name string `json:"name"`
}

type ResUDIDRegister struct {
	Token string `json:"token"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	req := ReqUDIDRegister{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newUser, err := controller.NewUser(req.UDID, req.Name)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := ResUDIDRegister{
		Token: newUser.GetToken(),
	}

	bRes, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bRes)
	w.WriteHeader(http.StatusOK)
}
