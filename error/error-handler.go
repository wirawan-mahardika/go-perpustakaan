package exception

import (
	"encoding/json"
	"net/http"
	"perpustakaan/model/web"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, i interface{}) {

	internalServerError(w, r, i)
}

func internalServerError(w http.ResponseWriter, r *http.Request, i interface{}) {
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
