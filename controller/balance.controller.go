package controller

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mr687/simple-go-rest-api/dto"
	"github.com/mr687/simple-go-rest-api/entity"
	"github.com/mr687/simple-go-rest-api/repository"
	"github.com/mr687/simple-go-rest-api/response"
	"github.com/mr687/simple-go-rest-api/service"
)

func (s *Server) GetBalance(c *gin.Context) {
	repo := repository.NewBalanceRepository(s.DB)

	userId, _ := service.GetTokenId(c)
	if userId == 0 {
		response.Unauthorized(c)
		return
	}

	balance, err := repo.GetCurrentBalance(userId)
	var amount int64
	if err != nil {
		amount = 0
	}

	if balance == nil {
		balance = entity.EmptyUserBalance(userId)
	}

	amount = balance.BalanceAchieve
	res := gin.H{"balance": amount}
	response.Ok(c, "Get current balance successfully", res)
}

func (s *Server) AddBalance(c *gin.Context) {
	var request dto.AddBalanceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.ValidationError(c, err)
		return
	}

	repo := repository.NewBalanceRepository(s.DB)
	userRepo := repository.NewUserRepository(s.DB)
	bankBalanceRepo := repository.NewBankBalanceRepository(s.DB)

	userId, _ := service.GetTokenId(c)
	if userId == 0 {
		response.Unauthorized(c)
		return
	}
	user, _ := userRepo.FindByID(userId)

	var newBalance *entity.UserBalance
	currentBalance, err := repo.GetCurrentBalance(userId)
	if err != nil {
		// create new balance
		newBalance = entity.EmptyUserBalance(userId)
	} else {
		newBalance = currentBalance
	}

	newBalance.Balance = request.Amount
	newBalance.BalanceAchieve += request.Amount

	// save UserBalance
	err = repo.SaveBalance(c, newBalance)
	if err != nil {
		response.ServerError(c, err)
		return
	}

	// save UserBalanceHistory
	err = repo.SaveBalanceHistory(c, newBalance, "debit", fmt.Sprintf("Add balance %v", request.Amount))
	if err != nil {
		response.ServerError(c, err)
		return
	}

	// save user balance to bank balance & bank history
	currentBankBalance := bankBalanceRepo.GetBankBalance()
	if currentBankBalance.Enabled {
		bankBalance, _ := bankBalanceRepo.SaveBalanceToBank(currentBankBalance, newBalance)
		bankBalanceRepo.SaveBalanceHistory(c, bankBalance, fmt.Sprintf("Add balance %v from %s", request.Amount, user.Username))
	}

	response.Ok(c, "Add balance successfully", nil)
}

func (s *Server) SendBalance(c *gin.Context) {
	var request dto.SendBalanceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.ValidationError(c, err)
		return
	}

	repo := repository.NewBalanceRepository(s.DB)
	userRepo := repository.NewUserRepository(s.DB)

	userId, _ := service.GetTokenId(c)
	if userId == 0 {
		response.Unauthorized(c)
		return
	}
	sender, _ := userRepo.FindByID(userId)

	// 1. Check if user has enough balance
	// - Get current balance
	senderBalance, err := repo.GetCurrentBalance(userId)
	if err != nil {
		// - If not found, it's means user has no balance
		response.BadRequest(c, "Not enough balance")
		return
	}
	// - then check if current balance is enough
	if senderBalance.BalanceAchieve < request.Amount {
		response.BadRequest(c, "Not enough balance")
		return
	}

	// 2. Check if recipient user exists
	// - Get recipient user
	recipient, err := userRepo.FindByUsernameOrEmail(request.Recipient)
	if err != nil {
		response.ValidationError(c, errors.New("recipient user not found"))
	}

	// 3. Check if recipient user is not the same as sender
	// - Check if sender user id is the same as recipient user
	if recipient.Id == userId {
		response.BadRequest(c, "Recipient user is the same as sender")
		return
	}

	// 4. Send Balance
	// Fo Sender:
	// - Subtract balance
	senderBalance.Balance = -request.Amount
	senderBalance.BalanceAchieve -= request.Amount
	// - Save balance
	repo.SaveBalance(c, senderBalance)
	// - Save balance history
	repo.SaveBalanceHistory(c, senderBalance, "kredit", fmt.Sprintf("Send %v to %s", request.Amount, recipient.Username))

	// For Recipient:
	// - Get current balance
	recipientBalance, _ := repo.GetCurrentBalance(recipient.Id)
	// - Add balance
	recipientBalance.Balance = request.Amount
	recipientBalance.BalanceAchieve += request.Amount
	// - Save balance
	repo.SaveBalance(c, recipientBalance)
	// - Save balance history
	repo.SaveBalanceHistory(c, recipientBalance, "debit", fmt.Sprintf("Receive %v from %s", request.Amount, sender.Username))

	// 5. Return response
	response.Ok(c, "Send balance successfully", nil)
}
