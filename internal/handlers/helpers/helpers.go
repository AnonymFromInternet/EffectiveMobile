package helpers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AnonymFromInternet/EffectiveMobile/internal/models"
)

type Payload struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type Response struct {
	Status  int           `json:"status"`
	Error   bool          `json:"error"`
	Message string        `json:"message"`
	Songs   []models.Song `json:"songs"`
}

func SendResponse(w http.ResponseWriter, statusHeader int, payload any) {
	w.WriteHeader(statusHeader)
	w.Header().Set("Content-Type", "application/json")

	if e := json.NewEncoder(w).Encode(payload); e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}
}

func GetResponse(status int, e bool, songs []models.Song, message string) Response {
	return Response{
		Error:   e,
		Status:  status,
		Songs:   songs,
		Message: message,
	}
}

func ConvertDate(date string) (string, error) {
	parsedDate, e := time.Parse("02.01.2006", date)
	if e != nil {
		return "", e
	}

	return parsedDate.Format("2006.01.02"), nil

}
