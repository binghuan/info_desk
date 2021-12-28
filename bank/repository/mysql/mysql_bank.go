package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"

	"infodesk/domain"
)

type mysqlBankRepository struct {
	Conn4AgentDB  *sql.DB
	Conn4DeviceDB *sql.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewMysqlBankRepository(Conn4AgentDB *sql.DB, Conn4DeviceDB *sql.DB) domain.BankRepository {
	return &mysqlBankRepository{
		Conn4AgentDB:  Conn4AgentDB,
		Conn4DeviceDB: Conn4DeviceDB,
	}
}

func (m *mysqlBankRepository) FetchBanks(ctx context.Context) (res []domain.BeneficiaryBank, err error) {
	query := `SELECT id, code, name, is_withdrawal, is_testable,
	support_type, creator_id, creator_name, created_at,
	updater_id, updater_name, updated_at,
	data_state, fish_bank_name, fish_bank_code FROM beneficiary_bank WHERE is_withdrawal=1`

	res, err = m.doFetchBanks(ctx, query)
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlBankRepository) doFetchBanks(ctx context.Context, query string, args ...interface{}) (result []domain.BeneficiaryBank, err error) {
	rows, err := m.Conn4AgentDB.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.BeneficiaryBank, 0)
	for rows.Next() {
		t := domain.BeneficiaryBank{}
		err = rows.Scan(
			&t.ID,
			&t.Code,
			&t.Name,
			&t.IsWithdrawal,
			&t.IsTestable,
			&t.SupportType,
			&t.CreatorID,
			&t.CreatorName,
			&t.CreatedAt,
			&t.UpdaterID,
			&t.UpdaterName,
			&t.UpdatedAt,
			&t.DataState,
			&t.FishBankName,
			&t.FishBankCode,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlBankRepository) FetchBankBranches(ctx context.Context, typeCode *int, province string, city string) (res []domain.BankNode, err error) {

	query := `SELECT id, parent, cascade_code, type_code, sheng, shi, xian, name,
	creator_name, created_at, updater_name, updated_at FROM bank_node`

	where := "WHERE"

	if len(province) > 0 {
		where = fmt.Sprintf(`%s sheng = '%s'`, where, province)
	} else {
		where = fmt.Sprintf(`%s sheng = sheng`, where)
	}

	if len(city) > 0 {
		where = fmt.Sprintf(`%s AND shi = '%s'`, where, city)
	} else {
		where = fmt.Sprintf(`%s AND shi = shi`, where)
	}

	if typeCode == nil {
		where = fmt.Sprintf(`%s AND type_code = %d`, where, typeCode)
	} else {
		where = fmt.Sprintf(`%s AND type_code = type_code`, where)
	}

	query = fmt.Sprintf(`%s %s`, query, where)
	log.Println("Query:", query)

	res, err = m.doFetchBankBranches(ctx, query)
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlBankRepository) doFetchBankBranches(ctx context.Context, query string, args ...interface{}) (result []domain.BankNode, err error) {
	rows, err := m.Conn4DeviceDB.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.BankNode, 0)
	for rows.Next() {
		t := domain.BankNode{}
		err = rows.Scan(
			&t.ID,
			&t.Parent,
			&t.CascadeCode,
			&t.TypeCode,
			&t.Sheng,
			&t.Shi,
			&t.Xian,
			&t.Name,
			&t.CreatorName,
			&t.CreatedAt,
			&t.UpdaterName,
			&t.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}
