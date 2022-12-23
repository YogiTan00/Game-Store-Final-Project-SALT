package customer_handler

import (
	"encoding/json"
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

	result, errStoreCustomerFromUseCase := s_handler.CustomerUseCase.StoreCustomer(dataCustomer)

	if errStoreCustomerFromUseCase != nil {
		response, errMap := http_response.MapResponse(200, errStoreCustomerFromUseCase.Error())
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.WriteHeader(417) // 417 Expectation Failed
		w.Write(response)
		return
	}

	if result != nil {
		response, errMap := http_response.MapResponse(200, "Success")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.WriteHeader(200)
		w.Write(response)
	} else {
		response, errMap := http_response.MapResponse(200, "Data customer already")
		if errMap != nil {
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("Error mapping data"))
		}
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write(response)
	}

}
