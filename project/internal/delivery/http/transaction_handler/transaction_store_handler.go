package transaction_handler

import (
	"encoding/json"
	"game-store-final-project/project/domain/entity/transaction"
	"game-store-final-project/project/internal/delivery/http_request"
	"game-store-final-project/project/internal/delivery/http_response"
	"net/http"
)

type test_struct struct {
	Test string
}

func (s_handler *TransactionHandlerInteractor) StoreController(w http.ResponseWriter, r *http.Request) {
	var (
		req     http_request.RequestTransaction
		decoder = json.NewDecoder(r.Body)
	)

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}

	// create DTO items
	reqItem := req.DetailTransaction
	items := make([]*transaction.DTOItemPembelian, 0)
	// loop and append to DTO
	for _, item := range *reqItem {
		dataItem := &transaction.DTOItemPembelian{
			ItemId:          item.ItemId,
			JumlahPembelian: item.JumlahPembelian,
		}
		items = append(items, dataItem)
	}

	resultUseCase, errStoreTrxFromUseCase := s_handler.TransactionUseCase.StoreTransaction(req.CustomerId, req.TanggalPembelian, req.Voucher, items)
	if errStoreTrxFromUseCase != nil {
		response, _ := http_response.MapResponse(200, errStoreTrxFromUseCase.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}

	if resultUseCase == nil {
		response, errMap := http_response.MapResponse(200, "DATA CUSTOMER NOT FOUND")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}

		w.WriteHeader(404)
		w.Write(response)
		return
	}

	response, errMap := http_response.MapResponse(200, "Success")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}

	w.WriteHeader(200)
	w.Write(response)
}
