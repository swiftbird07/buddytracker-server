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

type ReqUDIDLogin struct {
	UDID string `json:"udid"`
}

type ResUDIDLogin struct {
	Token string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	req := ReqUDIDLogin{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := controller.GetUser(req.UDID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := ResUDIDLogin{
		Token: user.GetToken(),
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
