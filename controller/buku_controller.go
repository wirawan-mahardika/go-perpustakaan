package controller

import (
	"encoding/json"
	"net/http"
	"perpustakaan/model/web"
	"perpustakaan/service"

	"github.com/julienschmidt/httprouter"
)

type BukuController interface {
	FindAllOrSearch(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

func NewBukuController(service service.BukuService) BukuController {
	return &bukuControllerImpl{service: service}
}

type bukuControllerImpl struct {
	service service.BukuService
}

func (controller *bukuControllerImpl) FindAllOrSearch(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	request := web.RequestBukuFindAll{Judul: r.URL.Query().Get("judul")}
	bukus := controller.service.FindAllOrSearch(r.Context(), request)

	webResponse := web.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    bukus,
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
	if err != nil {
		panic(err)
	}
}
