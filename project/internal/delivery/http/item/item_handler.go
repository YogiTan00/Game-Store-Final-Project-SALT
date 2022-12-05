package item

import (
	item2 "game-store-final-project/project/domain/entity/item"
	"game-store-final-project/project/internal/delivery/http_response"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *ItemHandler) GetAllItem(w http.ResponseWriter, r *http.Request) {
	var (
		item   []*item2.Item
		errGet error
	)

	item, errGet = h.repoItem.GetAllItem(h.ctx)

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

func (h *ItemHandler) GetAllItemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var (
		item   []*item2.Item
		errGet error
	)

	item, errGet = h.repoItem.GetAllItemByID(h.ctx, vars["id"])

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
