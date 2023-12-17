package main

import (
	"net/http"
	"perpustakaan/app"
	"perpustakaan/controller"
	exception "perpustakaan/error"
	"perpustakaan/middleware"
	"perpustakaan/service"

	"github.com/julienschmidt/httprouter"
)

func main() {
	var apiPathStart = "/api/v1"
	db := app.ConnectDB()

	bukuService := service.NewBukuService(db)
	bukuController := controller.NewBukuController(bukuService)

	router := httprouter.New()
	router.PanicHandler = exception.ErrorHandler
	router.GET(apiPathStart+"/buku", bukuController.FindAllOrSearch)
	router.GET(apiPathStart+"/buku/:id_buku", bukuController.FindById)
	router.POST(apiPathStart+"/buku", bukuController.Insert)
	router.PATCH(apiPathStart+"/buku/:id_buku", bukuController.Update)
	router.DELETE(apiPathStart+"/buku/:id_buku", bukuController.Delete)

	AuthMiddleware := &middleware.AuthMiddleware{Handler: router}

	server := http.Server{
		Addr:    "localhost:1000",
		Handler: AuthMiddleware,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
