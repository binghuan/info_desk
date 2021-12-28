package http_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	bankHttp "infodesk/bank/delivery/http"
	"infodesk/domain"
	"infodesk/domain/mocks"
)

func TestFetchBanks(t *testing.T) {
	var mockBank domain.BeneficiaryBank
	err := faker.FakeData(&mockBank)
	assert.NoError(t, err)
	mockUCase := new(mocks.BankUsecase)
	mockListBank := make([]domain.BeneficiaryBank, 0)
	mockListBank = append(mockListBank, mockBank)
	mockUCase.On("FetchBanks", mock.Anything).Return(mockListBank, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/banks", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := bankHttp.BankHandler{
		BankUsecase: mockUCase,
	}
	err = handler.FetchBanks(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestFetchBankBranches(t *testing.T) {
	var mockBankNode domain.BankNode
	err := faker.FakeData(&mockBankNode)
	assert.NoError(t, err)
	mockUCase := new(mocks.BankUsecase)
	mockListBankNode := make([]domain.BankNode, 0)
	mockListBankNode = append(mockListBankNode, mockBankNode)

	bankCode := "104"
	province := "广东省"
	city := "深圳市"
	mockUCase.On("FetchBankBranches", mock.Anything, bankCode, province, city).Return(mockListBankNode, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("/branches?bankcode=%s&province=%s&city=%s", bankCode, province, city), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := bankHttp.BankHandler{
		BankUsecase: mockUCase,
	}
	err = handler.FetchBankBranches(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestFetchBranchesError(t *testing.T) {
	mockUCase := new(mocks.BankUsecase)

	bankCode := "304" // This is a situation where the bank code is not converted to a type code.
	province := "广东省"
	city := "深圳市"
	mockUCase.On("FetchBankBranches", mock.Anything, bankCode, province, city).Return(nil, domain.ErrInternalServerError)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("/branches?bankcode=%s&province=%s&city=%s", bankCode, province, city), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := bankHttp.BankHandler{
		BankUsecase: mockUCase,
	}
	err = handler.FetchBankBranches(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}
