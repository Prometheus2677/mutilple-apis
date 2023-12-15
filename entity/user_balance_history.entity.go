package entity

type HistoryType string

const (
	Debit  HistoryType = "debit"
	Kredit HistoryType = "kredit"
)

// UserBalanceHistory struct is represents a user_balance_histories table in database
type UserBalanceHistory struct {
	Id            uint64      `gorm:"primary_key:auto_increment" json:"id"`
	UserBalanceId uint64      `gorm:"not null" json:"-"`
	UserBalance   UserBalance `gorm:"foreignkey:UserBalanceId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user_balance"`
	BalanceBefore int64       `gorm:"not null" json:"balance_before"`
	BalanceAfter  int64       `gorm:"not null" json:"balance_after"`
	Activtiy      string      `gorm:"not null" json:"activity"`
	Type          HistoryType `sql:"type:ENUM('debit','kredit');not null" json:"type"`
	Ip            string      `gorm:"not null" json:"ip_address"`
	Location      string      `gorm:"not null" json:"location"`
	UserAgent     string      `gorm:"not null" json:"user_agent"`
	Author        string      `gorm:"not null" json:"author"`
}
