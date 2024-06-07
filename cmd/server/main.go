package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/swiftbird07/buddytracker-server/internal/router"
)

func main() {
	r := router.NewRouter()

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", "0.0.0.0", 3001), r)
	if err != nil {
		log.Fatalln(err)
	}
}
