package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleError(w http.ResponseWriter, err error) {
	fmt.Printf("Error: %v", err)
	w.WriteHeader(http.StatusInternalServerError)
}

func MakeHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/sessions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		dec := json.NewDecoder(r.Body)
		
	})
}
