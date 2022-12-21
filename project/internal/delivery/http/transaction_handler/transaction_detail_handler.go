package transaction_handler

import (
	"context"
	"game-store-final-project/project/domain/entity/transaction_detail"
	"game-store-final-project/project/internal/delivery/http_response"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *TransactionDetailHandler) GetAllTransactionDetailHandler(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.Background()
	)

	listTransaction, err := h.useCaseTransactionDetail.GetAllTransactionDetail(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	if len(listTransaction) > 0 {
		response, errMap := http_response.MapResponseListTransactionDetail(listTransaction, 200, "Success")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}

		w.WriteHeader(200)
		w.Write(response)
	} else {
		response, errMap := http_response.MapResponse(200, "TRANSACTION DETAIL NOT FOUND")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}

		w.WriteHeader(404)
		w.Write(response)
	}
}

func (h *TransactionDetailHandler) GetTransactionDetailByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var (
		ctx               = context.Background()
		transactionDetail *transaction_detail.TransactionDetail
		errGet            error
	)

	transactionDetail, errGet = h.useCaseTransactionDetail.GetTransactionDetailByID(ctx, vars["id"])
	if errGet != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errGet.Error()))
	}
	if transactionDetail != nil {
		response, errMap := http_response.MapResponseTransactionDetail(transactionDetail, 200, "Success")

		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.WriteHeader(200)
		w.Write(response)
	} else {
		response, errMap := http_response.MapResponse(200, "TRANSACTION DETAIL NOT FOUND")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}

		w.WriteHeader(404)
		w.Write(response)
	}
}

func (h *TransactionDetailHandler) GetAllTransactionDetailByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var (
		ctx               = context.Background()
		transactionDetail []*transaction_detail.TransactionDetail
		errGet            error
	)

	transactionDetail, errGet = h.useCaseTransactionDetail.GetAllTransactionDetailByID(ctx, vars["id"])
	if errGet != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errGet.Error()))
	}
	if len(transactionDetail) > 0 {
		response, errMap := http_response.MapResponseListTransactionDetail(transactionDetail, 200, "Success")

		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.WriteHeader(200)
		w.Write(response)
	} else {
		response, errMap := http_response.MapResponse(200, "TRANSACTION DETAIL NOT FOUND")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}

		w.WriteHeader(404)
		w.Write(response)
	}
}
