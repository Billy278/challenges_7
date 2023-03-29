package server

import (
	"challenges_7/config"
	"challenges_7/module/controller"
	repository "challenges_7/module/repository/book"
	service "challenges_7/module/service/book"
)

type Ctrs struct {
	BookCtr controller.BookController
}

func InitControllers() Ctrs {
	//why err if i do this?
	// dataStore := config.DataStore{}
	dataStore := config.ConnectDataStore()
	bookRepo := repository.NewBookRepositoryImpl(dataStore)
	bookServ := service.NewBookServiceImpl(bookRepo)
	bookCtr := controller.NewBookControllerImpl(bookServ)

	return Ctrs{
		BookCtr: bookCtr,
	}
}
