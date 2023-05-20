package http

import (
	"net/http"
	"sync/atomic"

	"github.com/aqua-farm/farm"
	"github.com/aqua-farm/pond"
	pondUsecase "github.com/aqua-farm/pond/usecase"
	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"
)

type PondHandler struct {
	pondUsecase pondUsecase.PondUsecase
}

func NewPondHandler(e *echo.Echo, pu pondUsecase.PondUsecase) {
	handler := &PondHandler{pondUsecase: pu}
	e.GET("/pond", handler.FetchPond)
	e.POST("/pond", handler.Store)
}

// variable for counting how many fetchFarm API is called
var fetchPondCounter int64

// Handler for fetching all pond
func (ph *PondHandler) FetchPond(c echo.Context) error {

	atomic.AddInt64(&fetchPondCounter, 1)

	uniqueUserAgents := make(map[string]int)

	userAgent := c.Request().Header.Get("User-Agent")

	if uniqueUserAgents[userAgent] == 0 {
		uniqueUserAgents[userAgent]++
	}

	count := atomic.LoadInt64(&fetchPondCounter)

	listAr, err := ph.pondUsecase.Fetch()

	if err != nil {
		return c.JSON(getStatusCode(err), err.Error())
	}

	response := Response{
		Count:           count,
		UniqueUserAgent: uniqueUserAgents[userAgent],
		Data:            listAr,
	}

	return c.JSON(http.StatusOK, response)
}

// handler for Storing Farm
func (ph *PondHandler) Store(c echo.Context) (err error) {
	var pond pond.Pond
	err = c.Bind(&pond)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&pond); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := ph.pondUsecase.Store(&pond)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, res)
}

// Handler for farm Update
func (ph *PondHandler) Update(c echo.Context) (err error) {
	var pond pond.Pond
	err = c.Bind(&pond)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&pond); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := ph.pondUsecase.Update(&pond)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, res)
}

type Response struct {
	Count           int64 `json:"count"`
	UniqueUserAgent int   `json:"unique_user_agent"`
	Data            []*pond.Pond
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

func isRequestValid(m *pond.Pond) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
