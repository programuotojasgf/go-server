package handlers

import (
	"encoding/json"
	"net/http"
	"phrases-server/database"
)

func GetReviewPhrasesEndpoint(w http.ResponseWriter, r *http.Request) {
	payload := database.GetReviewPhrases()
	json.NewEncoder(w).Encode(payload)
}
