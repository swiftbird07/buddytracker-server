package handler

import (
	"net/http"

	"github.com/swiftbird07/buddytracker-server/internal/controller"
)

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	_ = r.Form.Get("password")

	_ = controller.NewUser(name)

	w.WriteHeader(http.StatusOK)
}
