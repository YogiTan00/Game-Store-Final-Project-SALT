package customer_handler

import (
	"encoding/json"
	"fmt"
	"game-store-final-project/project/domain/entity"
	"game-store-final-project/project/internal/delivery/http_request"
	"net/http"
)

func (s_handler *CustomerHandler) Store(w http.ResponseWriter, r *http.Request) {
	var (
		req     http_request.RequestCustomer
		decoder = json.NewDecoder(r.Body)
	)

	fmt.Println(req)

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}

	customer, err := entity.NewCustomer(req.Nik, req.Nama, req.Alamat, req.NoTlp, req.JenisKelamin)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error build data"))
		return
	}

	errInsert := s_handler.repoCustomer.StoreCustomer(s_handler.ctx, customer)
	if errInsert != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errInsert.Error()))
		return
	}

	w.WriteHeader(200)
	fmt.Fprintf(w, "SUCCESS INSERT DATA")

}
