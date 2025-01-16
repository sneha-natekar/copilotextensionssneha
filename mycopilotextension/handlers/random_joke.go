package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type JokeResponse struct {
	Joke string `json:"joke"`
}

func RandomJoke(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Random Joke Called")
	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, "https://official-joke-api.appspot.com/jokes/random", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var jokeData map[string]interface{}
	json.Unmarshal(body, &jokeData)

	response := JokeResponse{
		Joke: fmt.Sprintf("%v: %v", jokeData["setup"], jokeData["punchline"]),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
