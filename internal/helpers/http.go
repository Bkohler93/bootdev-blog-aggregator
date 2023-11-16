package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to encode payload", err)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to write JSON to client")
	}
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {

	payload := map[string]interface{}{
		"error": msg,
	}

	RespondWithJSON(w, code, payload)
}
