package http

import (
	"net/http"

	"github.com/aqua-farm/farm"
	"github.com/labstack/echo"
)

type FarmHandler struct {
}

func NewFarmHandler(e *echo.Echo) {
	handler := &FarmHandler{}
	e.GET("/farm", handler.FetchFarm)
}

func (fh *FarmHandler) FetchFarm(c echo.Context) error {
	// ctx := c.Request().Context()

	// listGP, err := fh.GPUsecase.FetchGasPriceList(ctx)
	// if err != nil {
	// 	return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	// }

	return c.JSON(http.StatusOK, "HEHE")
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
