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
	// bagaimana cara kirim reqItem sebagai param ???
	// fmt.Println(reqItem)

	items := make([]*transaction.DTOItemPembelian, 0)
	// loop and append to DTO
	for _, item := range *reqItem {
		dataItem := &transaction.DTOItemPembelian{
			ItemId:          item.ItemId,
			JumlahPembelian: item.JumlahPembelian,
		}
		items = append(items, dataItem)
	}

	_, errStoreTrxFromUseCase := s_handler.TransactionUseCase.StoreTransaction(req.CustomerId, req.TanggalPembelian, req.Voucher, items)
	if errStoreTrxFromUseCase != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error on Usecase"))
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
