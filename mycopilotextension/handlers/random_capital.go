package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type CapitalResponse struct {
	Capital string `json:"capital"`
	Country string `json:"country"`
}

var capitals = []CapitalResponse{
	{"Paris", "France"},
	{"Tokyo", "Japan"},
	{"Ottawa", "Canada"},
	{"Canberra", "Australia"},
	{"Brasília", "Brazil"},
	{"Berlin", "Germany"},
	{"New Delhi", "India"},
	{"Pretoria", "South Africa"},
	{"Buenos Aires", "Argentina"},
	{"Cairo", "Egypt"},
}

func RandomCapital(w http.ResponseWriter, r *http.Request) {
	randomIndex := rand.Intn(len(capitals))
	response := capitals[randomIndex]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
