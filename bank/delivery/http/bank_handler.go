package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"infodesk/domain"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// BankHandler  represent the httphandler for article
type BankHandler struct {
	BankUsecase domain.BankUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewBankHandler(e *echo.Echo, us domain.BankUsecase) {
	handler := &BankHandler{
		BankUsecase: us,
	}
	e.GET("/banks", handler.FetchBanks)
	e.GET("/branches", handler.FetchBankBranches)
}

// FetchBanks will fetch the bank based on given params
func (a *BankHandler) FetchBanks(c echo.Context) error {
	ctx := c.Request().Context()
	listAr, err := a.BankUsecase.FetchBanks(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, listAr)
}

func (a *BankHandler) FetchBankBranches(c echo.Context) error {
	ctx := c.Request().Context()

	province := c.QueryParam("province")
	city := c.QueryParam("city")
	bankCode := c.QueryParam("bankcode")

	listAr, err := a.BankUsecase.FetchBankBranches(ctx, bankCode, province, city)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, listAr)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
