package exception

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/masrayfa/go-basic-rest/helper"
	"github.com/masrayfa/go-basic-rest/model/web"
)

func ErrorHandler(writer http.ResponseWriter, req *http.Request, err interface{}) {

	if notFoundError(writer, req, err) {
		return
	}

	if validationError(writer, req, err) {
		return
	}

	internalServerError(writer, req, err)
}

func validationError(writer http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, req *http.Request, err interface{}) bool {

	exception, ok := err.(NotFoundError)
	if ok {

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, req *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
