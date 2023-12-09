package helper

import "perpustakaan/model/web"

func ToWebResponse(code int, message string, data any) (response web.WebResponse) {
	response = web.WebResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}

	return response
}
