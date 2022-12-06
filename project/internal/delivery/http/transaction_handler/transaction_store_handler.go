package transaction_handler

import (
	"encoding/json"
	"fmt"
	"game-store-final-project/project/domain/entity/transaction"
	"game-store-final-project/project/internal/delivery/http_request"
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
		fmt.Println(errDecode)
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
		fmt.Println(errStoreTrxFromUseCase)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error on Usecase"))
		return
	}

	w.WriteHeader(200)
	fmt.Fprintf(w, "SUCCESS INSERT DATA")

}
