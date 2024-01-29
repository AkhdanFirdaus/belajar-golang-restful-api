package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/AkhdanFirdaus/belajar-golang-restful-api/app"
	"github.com/AkhdanFirdaus/belajar-golang-restful-api/controller"
	"github.com/AkhdanFirdaus/belajar-golang-restful-api/helper"
	"github.com/AkhdanFirdaus/belajar-golang-restful-api/repository"
	"github.com/AkhdanFirdaus/belajar-golang-restful-api/service"
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}