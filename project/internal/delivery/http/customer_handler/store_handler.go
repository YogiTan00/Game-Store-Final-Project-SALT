package customer_handler

import (
	"encoding/json"
	"fmt"
	"game-store-final-project/project/domain/entity/customer"
	"game-store-final-project/project/internal/delivery/http_request"
	"game-store-final-project/project/internal/delivery/http_response"
	"net/http"
)

func (s_handler *CustomerHandlerInteractor) StoreController(w http.ResponseWriter, r *http.Request) {
	var (
		req     http_request.RequestCustomer
		decoder = json.NewDecoder(r.Body)
	)

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}

	// build dto
	dataCustomer := customer.DTOCustomer{
		Nik:          req.Nik,
		Nama:         req.Nama,
		Alamat:       req.Alamat,
		NoTlp:        req.NoTlp,
		JenisKelamin: req.JenisKelamin,
	}

	_, errStoreCustomerFromUseCase := s_handler.CustomerUseCase.StoreCustomer(dataCustomer)
	if errStoreCustomerFromUseCase != nil {
		fmt.Println(errStoreCustomerFromUseCase)
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
