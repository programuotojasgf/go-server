package handlers

import (
	"encoding/json"
	"net/http"
	"phrases-server/dao"
)

func GetReviewPhrasesEndpoint(w http.ResponseWriter, r *http.Request) {
	payload := dao.GetReviewPhrases()
	json.NewEncoder(w).Encode(payload)
}
