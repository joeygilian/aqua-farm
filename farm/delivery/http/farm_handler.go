package http

import (
	"net/http"
	"strconv"

	farmUsecase "github.com/aqua-farm/farm/usecase"

	"github.com/aqua-farm/farm"
	"github.com/labstack/echo"
)

type FarmHandler struct {
	farmUsecase farmUsecase.FarmUsecase
}

func NewFarmHandler(e *echo.Echo, fu farmUsecase.FarmUsecase) {
	handler := &FarmHandler{farmUsecase: fu}
	e.GET("/farm", handler.FetchFarm)
}

func (fh *FarmHandler) FetchFarm(c echo.Context) error {
	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)

	cursorS := c.QueryParam("cursor")
	cursor, _ := strconv.Atoi(cursorS)

	// cursor := c.QueryParam("cursor")

	listAr, nextCursor, err := fh.farmUsecase.Fetch(int64(cursor), int64(num))

	if err != nil {
		return c.JSON(getStatusCode(err), err.Error())
	}
	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, listAr)

	// return c.JSON(http.StatusOK, "HEHE")
}

type ResponseError struct {
	Message string `json:"message"`
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case farm.ErrInternalServerError:
		return http.StatusInternalServerError
	case farm.ErrNotFound:
		return http.StatusNotFound
	case farm.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
