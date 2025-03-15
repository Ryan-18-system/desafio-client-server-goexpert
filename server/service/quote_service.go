package service

import (
	"encoding/json"
	"net/http"

	"github.com/Ryan-18-system/desafio-client-server-goexpert/server/database"
)

func ProcessCurrentQuote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	responseApi, err := GetCurrentQuote()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = database.InsertQuote(*responseApi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseApi.USDBRL.Bid)
}
