package handler

import (
	"bistro/internal/dal"
	"bistro/internal/service"
	"bistro/models"
	"encoding/json"
	"net/http"
)

func PostOrder(w http.ResponseWriter, r *http.Request, ordersRepo *dal.OrdersRepository) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		sendError(w, http.StatusBadRequest, "StatusBadRequest", err.Error())
		return
	}
	err = service.PostOrder(order, ordersRepo)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "StatusInternalServerError", err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)

}
