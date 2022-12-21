package transaction_handler

import (
	"context"
	"game-store-final-project/project/domain/entity/transaction"
	"game-store-final-project/project/internal/delivery/http_response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *TransactionHandler) GetTransactionByIDHandler(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		ctx  = context.Background()
	)
	transaction, err := h.useCaseTransaction.GetTransactionByID(ctx, vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	if transaction != nil {

		listTransDetail, errDetail := h.useCaseTransDetail.GetAllTransactionDetailByID(ctx, strconv.Itoa(transaction.GetID()))
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
		response, errMap := http_response.MapResponse(200, "TRANSACTION NOT FOUND")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}

		w.WriteHeader(404)
		w.Write(response)
	}
}

func (h *TransactionHandler) GetAllTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.Background()
	)

	listTransaction, err := h.useCaseTransaction.GetAllTransaction(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	if len(listTransaction) > 0 {
		dataTrans := make([]*transaction.Transaction, 0)
		for _, data := range listTransaction {
			listTransDetail, errDetail := h.useCaseTransDetail.GetAllTransactionDetailByID(ctx, strconv.Itoa(data.GetID()))
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
	} else {
		response, errMap := http_response.MapResponse(200, "TRANSACTION NOT FOUND")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}

		w.WriteHeader(404)
		w.Write(response)
	}
}

func (h *TransactionHandler) GetAllTransactionByCustomerIDHandler(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = context.Background()
		vars = mux.Vars(r)
	)

	listTransaction, err := h.useCaseTransaction.GetAllTransactionByCustomerID(ctx, vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	if len(listTransaction) > 0 {
		dataTrans := make([]*transaction.Transaction, 0)
		for _, data := range listTransaction {
			listTransDetail, errDetail := h.useCaseTransDetail.GetAllTransactionDetailByID(ctx, strconv.Itoa(data.GetID()))
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
	} else {
		response, errMap := http_response.MapResponse(200, "TRANSACTION NOT FOUND")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}

		w.WriteHeader(404)
		w.Write(response)
	}
}
