package middleware

import (
	"net/http"

	"github.com/masrayfa/go-basic-rest/helper"
	"github.com/masrayfa/go-basic-rest/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware AuthMiddleware) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	if "RAHASIA" == req.Header.Get("X-API-Key") {
		// ok
		middleware.Handler.ServeHTTP(writer, req)
	} else {
		// error
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
