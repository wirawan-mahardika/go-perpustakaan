package main

import (
	"fmt"
	"net/http"
	"perpustakaan/database"

	"github.com/julienschmidt/httprouter"
)

var apiPathStart = "/api/v1"

func main() {
	db := database.ConnectToDb()
	defer db.Close()

	router := httprouter.New()
	router.GET(apiPathStart, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello world")
	})

	server := http.Server{
		Addr: "localhost:1000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil { panic(err) }
}