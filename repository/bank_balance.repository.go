package repository

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mr687/simple-go-rest-api/entity"
	"github.com/mr687/simple-go-rest-api/service"
	"gorm.io/gorm"
)

// Store all user balances
type bankBalanceRepository struct {
	db *gorm.DB
}

func NewBankBalanceRepository(db *gorm.DB) *bankBalanceRepository {
	return &bankBalanceRepository{db}
}

func (repo *bankBalanceRepository) GetBankBalance() *entity.BankBalance {
	var bankBalance *entity.BankBalance
	err := repo.db.Where("enabled = ?", true).First(&bankBalance).Error
	if err != nil {
		// return new bank balank with empty balance
		return entity.EmptyBankBalance()
	}
	return bankBalance
}

func (repo *bankBalanceRepository) SaveBalanceToBank(bankBalance *entity.BankBalance, userBalance *entity.UserBalance) (*entity.BankBalance, error) {
	now := time.Now().Unix()

	bankBalance.Balance = userBalance.Balance
	bankBalance.BalanceAchieve += userBalance.Balance
	bankBalance.Code = fmt.Sprintf("%v-%v", userBalance.Id, now)
	bankBalance.Enabled = true

	return bankBalance, repo.db.Save(bankBalance).Error
}

func (repo *bankBalanceRepository) SaveBalanceHistory(c *gin.Context, bankBalance *entity.BankBalance, activity string) error {
	geoInfo := c.MustGet("geoip").(*service.GeoIp)

	balanceHistory := entity.BankBalanceHistory{
		BankBalanceId: bankBalance.Id,
		BalanceBefore: bankBalance.BalanceAchieve - bankBalance.Balance,
		BalanceAfter:  bankBalance.BalanceAchieve,
		Ip:            geoInfo.Ip,
		Location:      fmt.Sprintf("%s, %s", geoInfo.City, geoInfo.CountryName),
		UserAgent:     c.GetHeader("User-Agent"),
		Type:          "debit",
		Activtiy:      activity,
		Author:        "System",
	}
	return repo.db.Save(&balanceHistory).Error
}
