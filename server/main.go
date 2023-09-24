package main

import (
	"log"

	"github.com/digicon-trap1-2023/backend/handler"
	"github.com/digicon-trap1-2023/backend/infrastructure"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := infrastructure.NewGormDB()
	if err != nil {
		log.Fatal(err)
	}

	api := handler.NewAPI(handler.NewPingHandler(db), handler.NewAuthHandler(), handler.NewBookMarkHandler(), handler.NewDocumentHandler())
	handler.SetUpRouter(e, api)
}
