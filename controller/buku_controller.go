package controller

import (
	"encoding/json"
	"net/http"
	"perpustakaan/helper"
	"perpustakaan/model/web"
	"perpustakaan/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BukuController interface {
	FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

type BukuControllerImpl struct {
	Service service.BukuService
}

func NewBukuController(service service.BukuService) BukuController {
	return &BukuControllerImpl{Service: service}
}

func (controller *BukuControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	idBuku, err := strconv.Atoi(p.ByName("id_buku"))
	helper.PanicIfError(err)
	request := web.RequestBukuFindById{Id: idBuku}

	responseBuku := controller.Service.FindById(r.Context(), request)
	webResponse := helper.ToWebResponse(200, "Berhasil mengambil buku dengan id "+strconv.Itoa(request.Id), responseBuku)

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
