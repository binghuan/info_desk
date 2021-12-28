package mysql_test

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	bankMysqlRepo "infodesk/bank/repository/mysql"
	"infodesk/domain"
)

func TestFetchBanks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockBanks := []domain.BeneficiaryBank{
		{
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
		},
	}

	rows := sqlmock.NewRows([]string{
		"id", "code", "name", "is_withdrawal", "is_testable",
		"support_type", "creator_id", "creator_name", "created_at",
		"updater_id", "updater_name", "updated_at",
		"data_state", "fish_bank_name", "fish_bank_code"}).
		AddRow(
			mockBanks[0].ID,
			mockBanks[0].Code,
			mockBanks[0].Name,
			mockBanks[0].IsWithdrawal,
			mockBanks[0].IsTestable,
			mockBanks[0].SupportType,
			mockBanks[0].CreatorID,
			mockBanks[0].CreatorName,
			mockBanks[0].CreatedAt,
			mockBanks[0].UpdaterID,
			mockBanks[0].UpdaterName,
			mockBanks[0].UpdatedAt,
			mockBanks[0].DataState,
			mockBanks[0].FishBankName,
			mockBanks[0].FishBankCode,
		)

	query := `SELECT id, code, name, is_withdrawal, is_testable,
	support_type, creator_id, creator_name, created_at,
	updater_id, updater_name, updated_at,
	data_state, fish_bank_name, fish_bank_code FROM beneficiary_bank WHERE is_withdrawal=1`

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
	a := bankMysqlRepo.NewMysqlBankRepository(db, db)
	list, err := a.FetchBanks(context.TODO())
	assert.NoError(t, err)
	assert.Len(t, list, 1)
}

func TestFetchBranches(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockBanks := []domain.BankNode{
		{
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
		},
	}

	rows := sqlmock.NewRows([]string{
		"id", "parent", "cascade_code", "type_code", "sheng", "shi", "xian", "name",
		"creator_name", "created_at", "updater_name", "updated_at"}).
		AddRow(
			mockBanks[0].ID,
			mockBanks[0].Parent,
			mockBanks[0].CascadeCode,
			mockBanks[0].TypeCode,
			mockBanks[0].Sheng,
			mockBanks[0].Shi,
			mockBanks[0].Xian,
			mockBanks[0].Name,
			mockBanks[0].CreatorName,
			mockBanks[0].CreatedAt,
			mockBanks[0].UpdaterName,
			mockBanks[0].UpdatedAt,
		)

	query := `SELECT id, parent, cascade_code, type_code, sheng, shi, xian, name,
	creator_name, created_at, updater_name, updated_at FROM bank_node`

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
	a := bankMysqlRepo.NewMysqlBankRepository(db, db)

	typeCode := int(8)
	list, err := a.FetchBankBranches(context.TODO(), &typeCode, "广东省", "深圳市")
	assert.NoError(t, err)
	assert.Len(t, list, 1)
}
