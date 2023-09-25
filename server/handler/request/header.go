package request

import (
	"net/http"

	"github.com/digicon-trap1-2023/backend/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetUserId(ctx echo.Context) (uuid.UUID, error) {
	userIdInterface := ctx.Request().Context().Value(util.UserKey)

	userIdString, ok := userIdInterface.(string)
	if !ok {
		return uuid.UUID{}, echo.NewHTTPError(http.StatusInternalServerError, "userId is not string")
	}

	useId, err := uuid.Parse(userIdString)
	if err != nil {
		return uuid.UUID{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return useId, nil
}

func GetRole(ctx echo.Context) (string, error) {
	roleInterface := ctx.Request().Context().Value(util.RoleKey)

	role, ok := roleInterface.(string)
	if !ok {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "role is not string")
	}

	return role, nil
}