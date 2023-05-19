package http

import (
	"net/http"
	"strconv"

	"github.com/aqua-farm/farm"
	pondUsecase "github.com/aqua-farm/pond/usecase"
	"github.com/labstack/echo"
)

type PondHandler struct {
	pondUsecase pondUsecase.PondUsecase
}

func NewPondHandler(e *echo.Echo, pu pondUsecase.PondUsecase) {
	handler := &PondHandler{pondUsecase: pu}
	e.GET("/pond", handler.FetchPond)
}

func (fh *PondHandler) FetchPond(c echo.Context) error {
	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)

	cursorS := c.QueryParam("cursor")
	cursor, _ := strconv.Atoi(cursorS)

	// cursor := c.QueryParam("cursor")

	listAr, nextCursor, err := fh.pondUsecase.Fetch(int64(cursor), int64(num))

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
