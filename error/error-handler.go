package exception

import (
	"encoding/json"
	"log"
	"net/http"
	"perpustakaan/model/web"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, i interface{}) {
	if unauthorize(w, r, i) {
		return
	}

	if notFound(w, r, i) {
		return
	}

	internalServerError(w, r, i)
}

func unauthorize(w http.ResponseWriter, r *http.Request, i interface{}) bool {
	data, ok := i.(Unauthorize)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		response := web.WebResponse{
			Code:    401,
			Message: "UNAUTHORIZE",
			Data:    data.Message,
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

func notFound(w http.ResponseWriter, r *http.Request, i interface{}) bool {
	data, ok := i.(NotFound)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		response := web.WebResponse{
			Code:    404,
			Message: "NOT FOUND",
			Data:    data.Message,
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
