package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", RootRouteHandler)
	http.ListenAndServe(":8080", nil)
}

// AppHealth struct holds info about current application health
type AppHealth struct {
	Healthy bool `json:"healthy"`
}

// RootRouteHandler responds to `GET /` request
func RootRouteHandler(w http.ResponseWriter, r *http.Request) {
	response := AppHealth{Healthy: true}

	json.NewEncoder(w).Encode(response)
}
