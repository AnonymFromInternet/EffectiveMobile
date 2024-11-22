package helpers

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, statusHeader int, payload any) {
	w.WriteHeader(statusHeader)
	w.Header().Set("Content-Type", "application/json")

	if e := json.NewEncoder(w).Encode(payload); e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}
}
