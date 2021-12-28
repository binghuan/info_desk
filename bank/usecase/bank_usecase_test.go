package usecase_test

import (
	"context"
	"infodesk/domain"
	"infodesk/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	ucase "infodesk/bank/usecase"
)

func TestFetchBanks(t *testing.T) {
	mockBankRepo := new(mocks.BankRepository)
	mockBank := domain.BeneficiaryBank{
		ID:           "1",
		Code:         "123",
		Name:         "bank123",
		IsWithdrawal: 1,
		IsTestable:   1,
		SupportType:  0,
		CreatorID:    123,
		CreatorName:  "123",
		CreatedAt:    time.Now(),
		UpdaterID:    456,
		UpdaterName:  "456",
		UpdatedAt:    time.Now(),
		DataState:    0,
		FishBankName: "",
		FishBankCode: "",
	}

	mockListBank := make([]domain.BeneficiaryBank, 0)
	mockListBank = append(mockListBank, mockBank)

	t.Run("success", func(t *testing.T) {
		mockBankRepo.On("FetchBanks", mock.Anything).Return(mockListBank, nil).Once()

		u := ucase.NewBankUsecase(mockBankRepo, time.Second*2)
		list, err := u.FetchBanks(context.TODO())
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListBank))
		mockBankRepo.AssertExpectations(t)
	})
}

func TestFetchBranches(t *testing.T) {

	mockBankRepo := new(mocks.BankRepository)
	mockBankNode := domain.BankNode{
		ID:          3,
		Parent:      1957,
		CascadeCode: "A19|B3|",
		TypeCode:    8,
		Sheng:       "广东省",
		Shi:         "深圳市",
		Xian:        "",
		Name:        "平安银行深圳红宝支行",
		CreatorName: "---",
		CreatedAt:   time.Now(),
		UpdaterName: "---",
		UpdatedAt:   time.Now(),
	}

	mockListBank := make([]domain.BankNode, 0)
	mockListBank = append(mockListBank, mockBankNode)

	t.Run("success", func(t *testing.T) {
		mockBankRepo.On("FetchBankBranches", mock.Anything).Return(mockListBank, nil).Once()

		u := ucase.NewBankUsecase(mockBankRepo, time.Second*2)
		list, err := u.FetchBankBranches(context.TODO(), "307", "广东省", "深圳市")
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListBank))
		mockBankRepo.AssertExpectations(t)
	})
}
