package entity

// User struct is represents a bank_balances table in database
type BankBalance struct {
	Id             uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Balance        int64  `gorm:"not null" json:"balance"`
	BalanceAchieve int64  `gorm:"not null" json:"balance_achieve"`
	Code           string `gorm:"not null" json:"code"`
	Enabled        bool   `gorm:"not null" json:"enabled"`
}

func EmptyBankBalance() *BankBalance {
	return &BankBalance{
		Balance:        0,
		BalanceAchieve: 0,
		Enabled:        true,
	}
}
