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

	listTransaction, err := h.useCaseTransactionDetail.UcGetAllTransactionDetail(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	response, errMap := http_response.MapResponseListTransactionDetail(listTransaction, 200, "Success")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}

	w.WriteHeader(200)
	w.Write(response)
}

func (h *TransactionDetailHandler) GetTransactionDetailByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var (
		ctx               = context.Background()
		transactionDetail *transaction_detail.TransactionDetail
		errGet            error
	)

	transactionDetail, errGet = h.useCaseTransactionDetail.UcGetTransactionDetailByID(ctx, vars["id"])
	if errGet != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errGet.Error()))
	}

	response, errMap := http_response.MapResponseTransactionDetail(transactionDetail, 200, "Success")

	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}
	w.WriteHeader(200)
	w.Write(response)
}

func (h *TransactionDetailHandler) GeAllTransactionDetailByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var (
		ctx               = context.Background()
		transactionDetail []*transaction_detail.TransactionDetail
		errGet            error
	)

	transactionDetail, errGet = h.useCaseTransactionDetail.UcGetAllTransactionDetailByID(ctx, vars["id"])
	if errGet != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errGet.Error()))
	}

	response, errMap := http_response.MapResponseListTransactionDetail(transactionDetail, 200, "Success")

	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}
	w.WriteHeader(200)
	w.Write(response)
}
