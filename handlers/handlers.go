package handlers

import (
	"encoding/json"
	"net/http"
	"phrases-server/database"
	"strconv"
	"strings"
)

func GetReviewPhrasesEndpoint(w http.ResponseWriter, r *http.Request) {
	//TODO Validate input
	sortOrder := _getSortOrder(r)
	page := _getPage(r)
	limit := _getLimit(r)

	payload := database.GetReviewPhrases(sortOrder, page, limit)
	json.NewEncoder(w).Encode(payload)
}

func _getLimit(r *http.Request) int64 {
	if r.URL.Query().Get("limit") != "" {
		limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
		return limit
	}

	return 0
}

func _getPage(r *http.Request) int64 {
	if r.URL.Query().Get("page") != "" {
		page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 32)
		return page
	}

	return 1
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
