package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var sponsorSlice []Sponsor

type Sponsor struct{
	Id string
	Name string
	ImagePath string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", GetSponsorsHandler).Methods("GET")
	r.HandleFunc("/", NewSponsorHandler).Methods("POST")

	sponsorSlice = make([]Sponsor, 0)
	sponsorSlice = append(sponsorSlice, Sponsor{
		Id:        "1",
		Name:      "Gold Star",
		ImagePath: "goldstar.png",
	})

	log.Printf("Sponsors Service Started!")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetSponsorsHandler(writer http.ResponseWriter, request *http.Request) {
	respondWithJSON(writer, http.StatusOK, &sponsorSlice)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func NewSponsorHandler(writer http.ResponseWriter, request *http.Request) {
	var sponsor Sponsor
	err := json.NewDecoder(request.Body).Decode(&sponsor)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Got a Sponsor: %s", sponsor)
	sponsorSlice = append(sponsorSlice, sponsor)
	respondWithJSON(writer, http.StatusOK, &sponsor)
}


