package controller

import (
	"api-key-value-service/model"
	"api-key-value-service/service"
	"errors"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var keyService service.KeyValueService = service.KeyValueServiceImpl{}


type KeyController struct {
}

func (i KeyController) Set() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		request := new(model.KeyValue)
		if err := c.Bind(request); err != nil {
			log.Println(err)
			return handleError(err, http.StatusUnprocessableEntity, "Error Fetching Request Body")
		}
		v := validator.New()
		if err := v.Struct(request); err != nil {
			return handleError(err, http.StatusBadRequest, err.Error())
		}

		resp, err := keyService.Set(*request)
		if err != nil {
			return handleError(err, http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, resp)
	}
}
func (i KeyController) Get() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		key := c.QueryParam("key")
		if len(key) == 0 {
			err := errors.New("Missing key")
			return handleError(err, http.StatusBadRequest, err.Error())
		}
		resp, err := keyService.Get( key)
		if err != nil {
			return handleError(err, http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, resp)
	}
}
func (i KeyController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		resp, err := keyService.GetAll()
		if err != nil {
			return handleError(err, http.StatusInternalServerError, err.Error())
		}
		if resp==nil{
			return c.JSON(http.StatusNoContent, resp)
		}
		return c.JSON(http.StatusOK, resp)
	}
}
func handleError(err error, code int, message string) error {
	log.Error(err)
	appError := model.ApplicationError{Code: code, Message: message}
	log.Error(appError)
	return echo.NewHTTPError(code, appError)
}
