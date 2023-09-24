package handler

import "github.com/labstack/echo/v4"

type API struct {
	Auth     *AuthHandler
	Bookmark *BookMarkHandler
	Document *DocumentHandler
	Tag      *TagHandler
	Ping     *PingHandler
}

func NewAPI(
	ping *PingHandler,
	auth *AuthHandler,
	bookmark *BookMarkHandler,
	document *DocumentHandler,
	tag *TagHandler,
) API {
	return API{
		Auth:     auth,
		Bookmark: bookmark,
		Document: document,
		Tag:      tag,
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

	documentsGroup := e.Group("/documents")
	{
		documentsGroup.GET("", api.Document.GetDocuments)
		documentsGroup.POST("", nil)
		documentsGroup.GET("/:id", nil)
		documentsGroup.PATCH("/:id", nil)
	}

	tagGroup := e.Group("/tags")
	{
		tagGroup.GET("", api.Tag.GetTags)
		tagGroup.POST("", api.Tag.PostTag)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
