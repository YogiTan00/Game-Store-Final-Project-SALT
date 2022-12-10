package item

import (
	"context"
	item2 "game-store-final-project/project/domain/entity/item"
	"game-store-final-project/project/internal/delivery/http_response"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *ItemHandler) GetAllItemHandler(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = context.Background()
		item   []*item2.Item
		errGet error
	)

	item, errGet = h.useCaseItem.UcGetAllItem(ctx)

	if errGet != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errGet.Error()))
	}

	response, errMap := http_response.MapResponseListItem(item, 200, "Success")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}
	w.WriteHeader(200)
	w.Write(response)
}

func (h *ItemHandler) GetItemByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var (
		ctx    = context.Background()
		item   *item2.Item
		errGet error
	)

	item, errGet = h.useCaseItem.UcGetItemByID(ctx, vars["id"])

	if errGet != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errGet.Error()))
	}

	response, errMap := http_response.MapResponseItem(item, 200, "Success")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}
	w.WriteHeader(200)
	w.Write(response)
}
