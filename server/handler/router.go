package handler

import "github.com/labstack/echo/v4"

type API struct {
	Auth     *AuthHandler
	Bookmark *BookMarkHandler
	Document *DocumentHandler
	Tag      *TagHandler
	Request  *RequestHandler
	Ping     *PingHandler
}

func NewAPI(
	ping *PingHandler,
	auth *AuthHandler,
	bookmark *BookMarkHandler,
	document *DocumentHandler,
	tag *TagHandler,
	request *RequestHandler,
) API {
	return API{
		Auth:     auth,
		Bookmark: bookmark,
		Document: document,
		Tag:      tag,
		Ping:     ping,
		Request:  request,
	}
}

func SetUpRouter(e *echo.Echo, api API) {
	e.GET("/ping", api.Ping.Ping)

	authGroup := e.Group("/auth")
	{
		authGroup.POST("/signup", nil)
	}

	e.Use(api.Auth.AuthMiddleware)

	e.GET("/me", api.Auth.GetMe)

	otherGroup := e.Group("/other")
	{
		otherGroup.GET("/documents", api.Document.GetOtherDocuments)
	}

	bookmarkGroup := e.Group("/bookmark")
	{
		bookmarkGroup.GET("", nil)
	}

	documentsGroup := e.Group("/documents")
	{
		documentsGroup.GET("", api.Document.GetDocuments)
		documentsGroup.POST("", api.Document.PostDocument)
		documentsGroup.GET("/:id", api.Document.GetDocument)
		documentsGroup.PATCH("/:id", api.Document.PatchDocument)
		documentsGroup.DELETE("/:id", api.Document.DeleteDocument)
		documentsGroup.POST("/:id/bookmarks", api.Document.PostBookmark)
		documentsGroup.DELETE("/:id/bookmarks", api.Document.DeleteBookmark)
		documentsGroup.POST("/:id/reference", api.Document.PostReference)
		documentsGroup.DELETE("/:id/reference", api.Document.DeleteReference)
	}

	tagGroup := e.Group("/tags")
	{
		tagGroup.GET("", api.Tag.GetTags)
		tagGroup.POST("", api.Tag.PostTag)
	}

	requestGroup := e.Group("/requests")
	{
		requestGroup.GET("", api.Request.GetRequests)
		requestGroup.POST("", api.Request.PostRequest)
		requestGroup.GET("/withDocument", api.Request.GetRequestsWithDocument)
		requestGroup.DELETE("/:id", api.Request.DeleteRequest)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
