package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/masrayfa/go-basic-rest/app"
	"github.com/masrayfa/go-basic-rest/controller"
	"github.com/masrayfa/go-basic-rest/helper"
	"github.com/masrayfa/go-basic-rest/middleware"
	"github.com/masrayfa/go-basic-rest/repository"
	"github.com/masrayfa/go-basic-rest/service"
)

func main() {
	db := app.NewDB()

	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
