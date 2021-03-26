package handlers

import (
	"encoding/json"
	"net/http"
	"phrases-server/database"
	"strings"
)

func GetReviewPhrasesEndpoint(w http.ResponseWriter, r *http.Request) {
	sortOrder := _getSortOrder(r)

	payload := database.GetReviewPhrases(sortOrder)
	json.NewEncoder(w).Encode(payload)
}

func _getSortOrder(r *http.Request) database.SortOrder {
	sortOrderArgument := r.URL.Query().Get("sortByFrequency")
	var sortOrder database.SortOrder
	sortOrder = database.Descending
	if strings.ToUpper(sortOrderArgument) == "ASC" {
		sortOrder = database.Ascending
	}

	return sortOrder
}
