package v1

import (
	"github.com/gin-gonic/gin"
	"spotify/dto"
	"spotify/helper"
	"spotify/helper/common"
	"spotify/services"
)

type AccountHandlerImpl struct {
	accountService services.AccountService
}

func NewAccountHandler(accountService services.AccountService) *AccountHandlerImpl {
	return &AccountHandlerImpl{
		accountService: accountService,
	}
}

func (s *AccountHandlerImpl) CreateAccount(context *gin.Context) {
	var (
		out = &dto.AccountsResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	username := context.Request.FormValue("username")
	email := context.Request.FormValue("email")
	password := context.Request.FormValue("password")
	passwordHash, err := helper.HashPassword(password)
	if err != nil {
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	in := dto.CreateAccountRequest{
		Username: username,
		Email:    email,
		Password: passwordHash,
		Role:     "User",
		Status:   "Active",
	}
	code := s.accountService.Create(context, in)
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Accounts = append(out.Accounts, dto.Account{
		Username: username,
		Email:    email,
		Role:     "User",
		Status:   "Active",
	})
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}
