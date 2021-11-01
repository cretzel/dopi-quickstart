package router

import (
	"dopi-quickstart/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ItemRouter struct {
	MuxRouter *mux.Router
	service   model.ItemService
}

func NewItemRouter(service model.ItemService) *ItemRouter {

	muxRouter := mux.NewRouter()
	amw := AuthMiddleware{}
	muxRouter.Use(amw.authMiddleware)
	quickstartRouter := ItemRouter{MuxRouter: muxRouter, service: service}

	muxRouter.HandleFunc("/api/quickstart/items", quickstartRouter.GetItems).Methods("GET")
	muxRouter.HandleFunc("/api/quickstart/items", quickstartRouter.CreateItem).Methods("POST")

	return &quickstartRouter
}

func (router *ItemRouter) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := router.service.GetItems()
	if err != nil {
		StatusError(w, err)
		return
	}

	response := toItemDtos(items)
	json.NewEncoder(w).Encode(response)
}

func (router *ItemRouter) CreateItem(w http.ResponseWriter, r *http.Request) {
	var itemDto CreateItemDto
	err := json.NewDecoder(r.Body).Decode(&itemDto)
	if err != nil {
		Error(w, 400, "Bad Request")
		return
	}

	item := createItemDtoToItem(&itemDto)
	username := router.claims(r).Username
	item.User = username
	newItem, err := router.service.CreateItem(item)
	if err != nil {
		StatusError(w, err)
		return
	}

	json.NewEncoder(w).Encode(toItemDto(newItem))
}

func (router *ItemRouter) claims(r *http.Request) *DopiClaims {
	return (r.Context().Value("claims")).(*DopiClaims)
}
