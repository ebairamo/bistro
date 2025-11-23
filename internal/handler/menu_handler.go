package handler

import (
	"bistro/internal/dal"
	"bistro/internal/service"
	"bistro/models"
	"encoding/json"
	"net/http"
)

func AddMenuItem(w http.ResponseWriter, r *http.Request, menuRepo *dal.MenuRepository) {
	menu := models.MenuItem{}
	err := json.NewDecoder(r.Body).Decode(&menu)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "StatusInternalServerError", err.Error())
		return
	}
	err = service.AddMenuItem(menuRepo, menu)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Status Internal Server Error", err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(menu)
}
