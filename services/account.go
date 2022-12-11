package services

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"github.com/google/uuid"
	"spotify/cache"
	"spotify/config"
	"spotify/dto"
	"spotify/helper/common"
	"spotify/models"
	"spotify/repositories"
	"time"
)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type UserService interface {
	Update(ctx context.Context, username string, req dto.UpdateAccountRequest) common.SubReturnCode
	FindByUserName(ctx context.Context, username string) (dto.Account, common.SubReturnCode)
}

func (a *AccountServiceImpl) FindByUserName(ctx context.Context, username string) (dto.Account, common.SubReturnCode) {
	acc, err := a.accountRepository.FindByUserName(ctx, username)
	if err != nil {
		glog.Errorln("FindByUserName failed: ", err)
		return dto.Account{}, common.SystemError
	}
	return dto.Account{
		Username: acc.UserName,
		Role:     string(acc.Role),
		Status:   string(acc.Status),
	}, common.OK
}

func (a *AccountServiceImpl) Update(ctx context.Context, username string, req dto.UpdateAccountRequest) common.SubReturnCode {
	glog.Infoln("Update request: ", req)
	rowsAffected, err := a.accountRepository.UpdateByUsername(ctx, username, models.Account{
		Status: models.AccountStatus(req.Status),
		Role:   models.AccountRole(req.Role),
	})
	if err != nil {
		glog.Errorln("Update failed: ", err)
		return common.SystemError
	}

	if rowsAffected == 0 {
		glog.Errorln("User not exist", username)
		return common.SystemError
	}

	err = a.serverCache.SetCode([]byte(uuid.New().String()), fmt.Sprintf("%s:%s", common.PrefixLoginCode, username), time.Duration(a.authConfig.ExpiredTime))
	if err != nil {
		glog.Errorln(err)
		return common.SystemError
	}

	return common.OK
}

func NewAccountService(usersRepository repositories.AccountRepository, serverCache cache.ServerCache, authConfig *config.Auth) UserService {
	return &AccountServiceImpl{
		accountRepository: usersRepository,
		serverCache:       serverCache,
		authConfig:        authConfig,
	}
}

type AccountServiceImpl struct {
	accountRepository repositories.AccountRepository
	serverCache       cache.ServerCache
	authConfig        *config.Auth
}
