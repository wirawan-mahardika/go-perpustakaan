package main

import (
	"net/http"
	"perpustakaan/app"
	"perpustakaan/controller"
	"perpustakaan/service"

	"github.com/julienschmidt/httprouter"
)

var apiPathStart = "/api/v1"

func main() {
	db := app.ConnectDB()

	bukuService := service.NewBukuService(db)
	bukuController := controller.NewBukuController(bukuService)

	router := httprouter.New()
	router.GET(apiPathStart+"/buku", bukuController.FindAllOrSearch)
	router.GET(apiPathStart+"/buku/:id_buku", bukuController.FindById)

	server := http.Server{
		Addr:    "localhost:1000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
