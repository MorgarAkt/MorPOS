package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshall JSON Response %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responded with internal server error: %v", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	RespondWithJson(w, code, errResponse{
		Error: msg,
	})
}
