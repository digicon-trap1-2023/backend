package main

import (
	"log"
	"net/http"

	"github.com/digicon-trap1-2023/backend/handler"
	"github.com/digicon-trap1-2023/backend/infrastructure"
	"github.com/digicon-trap1-2023/backend/interfaces/repository"
	"github.com/digicon-trap1-2023/backend/usecases/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPut,
				http.MethodPost,
				http.MethodDelete,
			},
			AllowHeaders: []string{
				echo.HeaderOrigin,
				echo.HeaderContentType,
				echo.HeaderAccept,
				"X-Role",
				"X-UserId",
			},
			AllowOrigins: []string{
				"https://digi-con2023-trap-1.trap.show",
				"http://localhost:5173",
				"https://localhost:5173",
			},
		}))

	db, err := infrastructure.NewGormDB()
	if err != nil {
		log.Fatal(err)
	}

	client, err := infrastructure.NewS3Client()
	if err != nil {
		log.Fatal(err)
	}

	documentRepository := repository.NewDocumentRepository(db, client)
	tagRepository := repository.NewTagRepository(db)
	userRepository := repository.NewUserRepository(db)
	requestRepository := repository.NewRequestRepository(db)

	authService := service.NewAuthService(userRepository)
	documentService := service.NewDocumentService(documentRepository)
	tagService := service.NewTagService(tagRepository)
	requestService := service.NewRequestService(requestRepository)

	api := handler.NewAPI(
		handler.NewPingHandler(),
		handler.NewAuthHandler(authService),
		handler.NewBookMarkHandler(),
		handler.NewDocumentHandler(documentService),
		handler.NewTagHandler(tagService),
		handler.NewRequestHandler(requestService),
	)

	handler.SetUpRouter(e, api)
}
