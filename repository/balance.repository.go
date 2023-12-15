package repository

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mr687/simple-go-rest-api/entity"
	"github.com/mr687/simple-go-rest-api/service"
	"gorm.io/gorm"
)

type balanceRepository struct {
	db *gorm.DB
}

func NewBalanceRepository(db *gorm.DB) *balanceRepository {
	return &balanceRepository{db}
}

func (repo *balanceRepository) GetCurrentBalance(userId uint64) (*entity.UserBalance, error) {
	var balance *entity.UserBalance
	err := repo.db.Table("user_balances").Where("user_id = ?", userId).Last(&balance).Error
	if err != nil {
		return entity.EmptyUserBalance(userId), err
	}

	return balance, nil
}

func (repo *balanceRepository) SaveBalance(c *gin.Context, newBalance *entity.UserBalance) error {
	if err := repo.db.Save(&newBalance).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (repo *balanceRepository) SaveBalanceHistory(c *gin.Context, newBalance *entity.UserBalance, historyType entity.HistoryType, activity string) error {
	geoInfo := c.MustGet("geoip").(*service.GeoIp)

	balanceHistory := entity.UserBalanceHistory{
		UserBalanceId: newBalance.Id,
		BalanceBefore: newBalance.BalanceAchieve - newBalance.Balance,
		BalanceAfter:  newBalance.BalanceAchieve,
		Ip:            geoInfo.Ip,
		Location:      fmt.Sprintf("%s, %s", geoInfo.City, geoInfo.CountryName),
		UserAgent:     c.GetHeader("User-Agent"),
		Type:          historyType,
		Activtiy:      activity,
		Author:        "System",
	}
	return repo.db.Save(&balanceHistory).Error
}
