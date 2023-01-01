package voucher_handler

import (
	"context"
	"game-store-final-project/project/internal/delivery/http_response"
	"github.com/gorilla/mux"
	"net/http"
)

func (v *VoucherHandler) GetVoucherCustomerByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var (
		ctx    = context.Background()
		errGet error
	)

	VoucherCustomer, errGet := v.useCaseVoucher.GetVoucherByCustomerId(ctx, vars["id"])
	if VoucherCustomer != nil {
		if errGet != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errGet.Error()))
		}

		response, errMap := http_response.MapResponseVoucherCustomer(VoucherCustomer, 200, "Success")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.WriteHeader(200)
		w.Write(response)
	} else {
		response, errMap := http_response.MapResponse(200, "VOUCHER CUSTOMER NOT FOUND")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}

		w.WriteHeader(404)
		w.Write(response)
	}
}
