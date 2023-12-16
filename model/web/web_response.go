package web

type WebResponse struct {
	Code    int16  `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
