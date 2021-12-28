package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"infodesk/domain"
)

type bankUsecase struct {
	bankRepo       domain.BankRepository
	contextTimeout time.Duration
}

// NewBankUsecase will create new an articleUsecase object
// representation of domain.BankUsecase interface
func NewBankUsecase(a domain.BankRepository,
	timeout time.Duration) domain.BankUsecase {
	return &bankUsecase{
		bankRepo:       a,
		contextTimeout: timeout,
	}
}

func (a *bankUsecase) FetchBanks(c context.Context) (res []domain.BeneficiaryBank, err error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.bankRepo.FetchBanks(ctx)
	if err != nil {
		return nil, err
	}
	return
}

func convertBankCodeToTypeCode(bankCode string) int {
	typeCode := -1
	// convert bank code to type code.
	switch {
	case bankCode == "102": //工商银行
		typeCode = 4
	case bankCode == "103": //农业银行
		typeCode = 5
	case bankCode == "104": // 中国银行
		typeCode = 7
	case bankCode == "105": // 建設銀行
		typeCode = 6
	case bankCode == "301": // 交通銀行
		typeCode = 11
	case bankCode == "303": // 光大银行
		typeCode = 12
	case bankCode == "307": // 平安銀行
		typeCode = 8
	case bankCode == "308": // 招商银行
		typeCode = 18
	case bankCode == "309": // 兴业银行
		typeCode = 10
	case bankCode == "310": // 浦发银行
		typeCode = 9
	case bankCode == "2045": // 郵政銀行
		typeCode = 16
	}
	return typeCode
}

func (a *bankUsecase) FetchBankBranches(c context.Context, bankCode string, province string, city string) (res []domain.BankNode, err error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	if len(bankCode) == 0 {
		return nil, errors.New("please give me the bank code to query")
	}
	typeCode := convertBankCodeToTypeCode(bankCode)
	if typeCode == -1 {
		return nil, fmt.Errorf("unknown bankCode: %+v", bankCode)
	}

	res, err = a.bankRepo.FetchBankBranches(ctx, &typeCode, province, city)
	if err != nil {
		return nil, err
	}
	return
}
