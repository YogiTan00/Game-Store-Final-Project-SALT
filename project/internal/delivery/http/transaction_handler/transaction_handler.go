package transaction_handler

import (
	"game-store-final-project/project/internal/delivery/http_response"
	"net/http"
)

func (h *TransactionHandler) GetAllTransaction(w http.ResponseWriter, r *http.Request) {
	listTransaction, err := h.repoTransaction.GetAllTransaction(h.ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	response, errMap := http_response.MapResponseListTransaction(listTransaction, 200, "Success")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}

	w.WriteHeader(200)
	w.Write(response)
}
