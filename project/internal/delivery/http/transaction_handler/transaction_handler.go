package transaction_handler

import (
	"context"
	"game-store-final-project/project/domain/entity/transaction"
	"game-store-final-project/project/internal/delivery/http_response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *TransactionHandler) GetTransactionByIDHandler(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		ctx  = context.Background()
	)
	transaction, err := h.useCaseTransaction.UcGetTransactionByID(ctx, vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	if transaction != nil {

		listTransDetail, errDetail := h.useCaseTransDetail.UcGetAllTransactionDetailByID(ctx, strconv.Itoa(transaction.GetID()))
		if errDetail != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errDetail.Error()))
		}

		transaction = transaction.AddTransDetail(listTransDetail)

		response, errMap := http_response.MapResponseTransaction(transaction, 200, "Success")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}

		w.WriteHeader(200)
		w.Write(response)
	} else {
		w.WriteHeader(404)
	}
}

func (h *TransactionHandler) GetAllTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.Background()
	)

	listTransaction, err := h.useCaseTransaction.UcGetAllTransaction(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	dataTrans := make([]*transaction.Transaction, 0)
	for _, data := range listTransaction {
		listTransDetail, errDetail := h.useCaseTransDetail.UcGetAllTransactionDetailByID(ctx, strconv.Itoa(data.GetID()))
		if errDetail != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errDetail.Error()))
		}
		data = data.AddTransDetail(listTransDetail)
		dataTrans = append(dataTrans, data)
	}

	response, errMap := http_response.MapResponseListTransaction(dataTrans, 200, "Success")
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

	listTransaction, err := h.useCaseTransaction.UcGetAllTransactionByCustomerID(ctx, vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	dataTrans := make([]*transaction.Transaction, 0)
	for _, data := range listTransaction {
		listTransDetail, errDetail := h.useCaseTransDetail.UcGetAllTransactionDetailByID(ctx, strconv.Itoa(data.GetID()))
		if errDetail != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errDetail.Error()))
		}
		data = data.AddTransDetail(listTransDetail)
		dataTrans = append(dataTrans, data)
	}

	response, errMap := http_response.MapResponseListTransaction(listTransaction, 200, "Success")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}

	w.WriteHeader(200)
	w.Write(response)
}
