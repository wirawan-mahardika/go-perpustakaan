package middleware

import (
	"encoding/json"
	"net/http"
	"perpustakaan/app"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	config := app.ViperConfig("config", "json", ".")
	configApiKey := config.GetString("security.api_key")
	apiKey := r.Header.Get("X-API-Key")

	if apiKey != configApiKey {
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.Encode(map[string]any{
			"code":    401,
			"message": "Request is not allowed",
			"data":    nil,
		})

		return
	}

	middleware.Handler.ServeHTTP(w, r)
}
