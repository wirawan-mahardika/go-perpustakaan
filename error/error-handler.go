package exception

import (
	"encoding/json"
	"log"
	"net/http"
	"perpustakaan/model/web"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, i interface{}) {
	if notFound(w, r, i) {
		return
	}

	internalServerError(w, r, i)
}

func notFound(w http.ResponseWriter, r *http.Request, i interface{}) bool {
	message, ok := i.(NotFound)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		response := web.WebResponse{
			Code:    404,
			Message: "NOT FOUND",
			Data:    message,
		}

		encoder := json.NewEncoder(w)
		err := encoder.Encode(response)
		if err != nil {
			panic(err)
		}
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, i interface{}) {
	log.Println(i)
	w.Header().Set("Content-Type", "application/json")
	response := web.WebResponse{
		Code:    500,
		Message: "INTERNAL SERVER ERROR",
		Data:    nil,
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		panic(err)
	}
}
