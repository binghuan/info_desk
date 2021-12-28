package domain

import (
	"context"
	"time"
)

// BeneficiaryBank 銀行資訊
type BeneficiaryBank struct {
	ID           string    `gorm:"column:id;primary_key" json:"id"`
	Code         string    `gorm:"column:code" json:"code" description:"受款銀行代碼"`
	Name         string    `gorm:"column:name" json:"name" description:"受款銀行名稱"`
	IsWithdrawal int       `gorm:"column:is_withdrawal" json:"isWithdrawal" description:""`
	IsTestable   int       `gorm:"column:is_testable" json:"IsTestable" description:""`
	SupportType  int       `gorm:"column:support_type" json:"supportType" description:""`
	CreatorID    int64     `gorm:"column:creator_id" json:"-" description:"資料建立人員"`
	CreatorName  string    `gorm:"column:creator_name" json:"-"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"-" description:"資料建立時間"`
	UpdaterID    int64     `gorm:"column:updater_id" json:"-" description:"資料更新人員"`
	UpdaterName  string    `gorm:"column:updater_name" json:"-"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"-" description:"資料更新時間"`
	DataState    int32     `gorm:"column:data_state" json:"-" description:"資料狀態"`
	FishBankName string    `gorm:"column:fish_bank_name" json:"-" description:"Fish協議對應的銀行名稱"`
	FishBankCode string    `gorm:"column:fish_bank_code" json:"-" description:"Fish協議對應的銀行代碼"`
}

type BankNode struct {
	ID          int32     `gorm:"column:id;primary_key" json:"id"`
	Parent      int32     `gorm:"column:parent" json:"parent"`
	CascadeCode string    `gorm:"column:cascade_code" json:"branchCode"`
	TypeCode    int32     `gorm:"column:type_code" json:"typeCode"`
	Sheng       string    `gorm:"column:sheng" json:"province" form:"sheng"`
	Shi         string    `gorm:"column:shi" json:"city" form:"shi"`
	Xian        string    `gorm:"column:xian" json:"-"`
	Name        string    `gorm:"column:name" json:"name" form:"name" url:"name"`
	CreatorName string    `gorm:"column:creator_name" json:"-"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"-" description:"資料建立時間"`
	UpdaterName string    `gorm:"column:updater_name" json:"-"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"-" description:"資料更新時間"`
}

// ArticleUsecase represent the article's usecases
type BankUsecase interface {
	FetchBanks(ctx context.Context) ([]BeneficiaryBank, error)
	FetchBankBranches(ctx context.Context, bankCode string, province string, city string) ([]BankNode, error)
}

// ArticleRepository represent the article's repository contract
type BankRepository interface {
	FetchBanks(ctx context.Context) (res []BeneficiaryBank, err error)
	FetchBankBranches(ctx context.Context, typeCode *int, province string, city string) ([]BankNode, error)
}
