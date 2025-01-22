package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sneha-natekar/copilotextensionssneha/shared/handlers"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	http.HandleFunc("/random-joke", handlers.RandomJoke)
	http.HandleFunc("/random-capital", handlers.RandomCapital)
	http.HandleFunc("/random-historical-fact", handlers.RandomHistoricalFact)
	http.HandleFunc("/_ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.ListenAndServe(":8080", nil)
	return nil
}
