package transaction_handler

import (
	"encoding/json"
	"fmt"
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
		return
	}

	// create DTO items
	// bagaimana cara kirim reqItem sebagai param ???
	// fmt.Println(reqItem)

	_, errStoreTrxFromUseCase := s_handler.TransactionUseCase.StoreTransaction(req.CustomerId, req.TanggalPembelian, req.Voucher)
	if errStoreTrxFromUseCase != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error on Usecase"))
		return
	}

	w.WriteHeader(200)
	fmt.Fprintf(w, "SUCCESS INSERT DATA")

}
