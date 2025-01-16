package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type HistoricalFactResponse struct {
	Fact string `json:"fact"`
}

var facts = []string{
	"The Great Fire of London in 1666 destroyed much of the city but only six people were recorded as dying.",
	"Albert Einstein was offered the presidency of Israel in 1952 but declined.",
	"The Eiffel Tower can be 15 cm taller during hot days due to the expansion of iron.",
	"Ancient Egyptians used honey as an antibiotic and for wound healing.",
	"The first recorded Olympic Games were held in 776 BCE in Olympia, Greece.",
}

func RandomHistoricalFact(w http.ResponseWriter, r *http.Request) {
	randomIndex := rand.Intn(len(facts))
	response := HistoricalFactResponse{
		Fact: facts[randomIndex],
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
