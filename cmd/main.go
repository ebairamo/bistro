package main

import (
	"bistro/internal/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/inventory", inventoryHandler)
	http.ListenAndServe(":8000", nil)
}

func inventoryHandler(w http.ResponseWriter, r *http.Request) {

	handler.AddInventoryItem(w, r)
}
