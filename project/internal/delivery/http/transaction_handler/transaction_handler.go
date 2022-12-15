package transaction_handler

import (
	"context"
	"game-store-final-project/project/internal/delivery/http_response"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *TransactionHandler) GetTransactionByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var (
		ctx = context.Background()
	)
	Transaction, err := h.repoTransaction.GetTransactionByID(ctx, vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	listTransDetail, errDetail := h.repoTransactionDetail.GetAllTransactionDetailByID(ctx, vars["id"])
	if errDetail != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errDetail.Error()))
	}

	Transaction = Transaction.AddTransDetail(listTransDetail)

	response, errMap := http_response.MapResponseTransaction(Transaction, 200, "Success")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}

	w.WriteHeader(200)
	w.Write(response)
}

func (h *TransactionHandler) GetAllTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.Background()
	)

	listTransaction, err := h.repoTransaction.GetAllTransaction(ctx)
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

func (h *TransactionHandler) GetAllTransactionByCustomerIDHandler(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = context.Background()
		vars = mux.Vars(r)
	)

	listTransaction, err := h.repoTransaction.GetAllTransactionByCustomerID(ctx, vars["id"])
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
