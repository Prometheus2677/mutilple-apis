package entity

// User struct is represents a user_balances table in database
type UserBalance struct {
	Id             uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UserId         uint64 `gorm:"not null" json:"-"`
	User           User   `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	Balance        int64  `gorm:"not null" json:"balance"`
	BalanceAchieve int64  `gorm:"not null" json:"balance_achieve"`
}

func EmptyUserBalance(userId uint64) *UserBalance {
	return &UserBalance{
		UserId: userId,
		Balance: 0,
		BalanceAchieve: 0,
	}
}
