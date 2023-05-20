package http

import (
	"net/http"
	"sync/atomic"

	farmUsecase "github.com/aqua-farm/farm/usecase"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/aqua-farm/farm"
	"github.com/labstack/echo"
)

type FarmHandler struct {
	farmUsecase farmUsecase.FarmUsecase
}

func NewFarmHandler(e *echo.Echo, fu farmUsecase.FarmUsecase) {
	handler := &FarmHandler{farmUsecase: fu}
	e.GET("/farm", handler.FetchFarm)
	e.POST("/farm", handler.Store)
	e.PUT("/farm", handler.Update)
}

// variable for counting how many fetchFarm API is called
var fetchFarmCounter int64

// handler for fetching all farm
func (fh *FarmHandler) FetchFarm(c echo.Context) error {

	atomic.AddInt64(&fetchFarmCounter, 1)

	uniqueUserAgents := make(map[string]int)

	userAgent := c.Request().Header.Get("User-Agent")

	if uniqueUserAgents[userAgent] == 0 {
		uniqueUserAgents[userAgent]++
	}

	listAr, err := fh.farmUsecase.Fetch()

	if err != nil {
		return c.JSON(getStatusCode(err), err.Error())
	}

	count := atomic.LoadInt64(&fetchFarmCounter)

	response := Response{
		Count:           count,
		UniqueUserAgent: uniqueUserAgents[userAgent],
		Data:            listAr,
	}
	return c.JSON(http.StatusOK, response)
}

// handler for Storing Farm
func (fh *FarmHandler) Store(c echo.Context) (err error) {
	var farm farm.Farm
	err = c.Bind(&farm)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&farm); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := fh.farmUsecase.Store(&farm)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, res)
}

// Handler for farm Update
func (fh *FarmHandler) Update(c echo.Context) (err error) {
	var farm farm.Farm
	err = c.Bind(&farm)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&farm); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := fh.farmUsecase.Update(&farm)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, res)
}

type Response struct {
	Count           int64 `json:"count"`
	UniqueUserAgent int   `json:"unique_user_agent"`
	Data            []*farm.Farm
}

type ResponseError struct {
	Message string `json:"message"`
}

func isRequestValid(m *farm.Farm) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
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
