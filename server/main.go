package main

import (
	"github.com/digicon-trap1-2023/backend/handler"

    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()
	
	api := handler.NewAPI(handler.NewPingHandler(), handler.NewAuthHandler(), handler.NewBookMarkHandler(), handler.NewDocumentHandler())
	handler.SetUpRouter(e, api)
}