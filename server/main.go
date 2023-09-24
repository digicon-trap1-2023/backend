package main

import (
	"log"

	"github.com/digicon-trap1-2023/backend/handler"
	"github.com/digicon-trap1-2023/backend/infrastructure"
	"github.com/digicon-trap1-2023/backend/interfaces/repository"
	"github.com/digicon-trap1-2023/backend/usecases/service"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := infrastructure.NewGormDB()
	if err != nil {
		log.Fatal(err)
	}

	documentRepository := repository.NewDocumentRepository(db)

	documentService := service.NewDocumentService(documentRepository)

	api := handler.NewAPI(handler.NewPingHandler(), handler.NewAuthHandler(), handler.NewBookMarkHandler(), handler.NewDocumentHandler(documentService))

	handler.SetUpRouter(e, api)
}
