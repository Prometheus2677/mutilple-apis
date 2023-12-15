package entity

// BankBalanceHistory struct is represents a bank_balance_histories table in database
type BankBalanceHistory struct {
	Id            uint64      `gorm:"primary_key:auto_increment" json:"id"`
	BankBalanceId uint64      `gorm:"not null" json:"-"`
	BankBalance   BankBalance `gorm:"foreignkey:BankBalanceId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"bank_balance"`
	BalanceBefore int64       `gorm:"not null" json:"balance_before"`
	BalanceAfter  int64       `gorm:"not null" json:"balance_after"`
	Activtiy      string      `gorm:"not null" json:"activity"`
	Type          HistoryType `sql:"type:ENUM('debit','kredit')" json:"type"`
	Ip            string      `gorm:"not null" json:"ip_address"`
	Location      string      `gorm:"not null" json:"location"`
	UserAgent     string      `gorm:"not null" json:"user_agent"`
	Author        string      `gorm:"not null" json:"author"`
}
