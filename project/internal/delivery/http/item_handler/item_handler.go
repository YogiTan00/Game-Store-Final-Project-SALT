package item_handler

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

	item, errGet = h.useCaseItem.GetAllItem(ctx)
	if len(item) > 0 {
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
	} else {
		response, errMap := http_response.MapResponse(200, "ITEM NOT FOUND")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}

		w.WriteHeader(404)
		w.Write(response)
	}
}

func (h *ItemHandler) GetItemByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var (
		ctx    = context.Background()
		item   *item2.Item
		errGet error
	)

	item, errGet = h.useCaseItem.GetItemByID(ctx, vars["id"])
	if item != nil {
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
	} else {
		response, errMap := http_response.MapResponse(200, "ITEM NOT FOUND")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}

		w.WriteHeader(404)
		w.Write(response)
	}
}
