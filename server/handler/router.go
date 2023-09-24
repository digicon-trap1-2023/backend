package handler

import "github.com/labstack/echo/v4"

type API struct {
	Auth     *AuthHandler
	Bookmark *BookMarkHandler
	Document *DocumentHandler
	Ping     *PingHandler
}

func NewAPI(ping *PingHandler, auth *AuthHandler, bookmark *BookMarkHandler, document *DocumentHandler) API {
	return API{
		Auth:     auth,
		Bookmark: bookmark,
		Document: document,
		Ping:     ping,
	}
}

func SetUpRouter(e *echo.Echo, api API) {
	e.GET("/ping", api.Ping.Ping)

	authGroup := e.Group("/auth")
	{
		authGroup.POST("/signup", nil)
	}

	e.Use(api.Auth.AuthMiddleware)

	bookmarkGroup := e.Group("/bookmark")
	{
		bookmarkGroup.GET("", nil)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
