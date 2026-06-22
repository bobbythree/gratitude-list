package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (app *App) ListItemsHandler(w http.ResponseWriter, r *http.Request) {
	// cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		var incoming ListItem

		err := json.NewDecoder(r.Body).Decode(&incoming)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// create list item
		err = CreateListItem(app.DB, incoming)
		if err != nil {
			http.Error(w, "Failed to save item", http.StatusInternalServerError)
			return
		}

		// response
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(incoming)
		if err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
			return
		}

		return
	}

	// get
	if r.Method == http.MethodGet {
		items, err := GetListItems(app.DB)
		if err != nil {
			http.Error(w, "Failed to get items", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(items)
		if err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
			return
		}

		return
	}

	// delete
	if r.Method == http.MethodDelete {
		idStr := r.URL.Query().Get("id")

		if idStr == "" {
			http.Error(w, "Missing id", http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}

		err = DeleteListItem(app.DB, id)
		if err != nil {
			http.Error(w, "Failed to delete item", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
