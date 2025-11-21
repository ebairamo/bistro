package handler

import (
	"bistro/internal/dal"
	"bistro/internal/service"
	"bistro/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddInventoryItem(w http.ResponseWriter, r *http.Request, repo *dal.InventoryRepository) {
	if r.Method != http.MethodPost {
		sendError(w, http.StatusMethodNotAllowed, "Status Method Not Allowed", "Use post")
		return
	}
	var item models.InventoryItem
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = service.SaveItem(item, repo)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Status Internal Server Error", err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
	// TODO: проверить что метод POST

	// TODO: прочитать и распарсить JSON

	// TODO: вызвать service для сохранения

	// TODO: вернуть успешный ответ
}
