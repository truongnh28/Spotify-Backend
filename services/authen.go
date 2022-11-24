package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang/glog"
	"github.com/google/uuid"
	"spotify/cache"
	"spotify/config"
	"spotify/dto"
	"spotify/helper/common"
	"spotify/models"
	"spotify/repositories"
	"spotify/utils/auth"
	"time"
)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type AuthenService interface {
	Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponseData, error)
}

func (a *authenServiceImpl) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponseData, error) {
	return a.checkUserFromDB(ctx, req)
}

func (a *authenServiceImpl) checkUserFromDB(ctx context.Context, userInfo dto.LoginRequest) (*dto.LoginResponseData, error) {
	var (
		resp    *dto.LoginResponseData
		account *dto.Account
	)
	fmt.Println(userInfo.Username)
	acc, err := a.accountRepo.FindByUserName(ctx, userInfo.Username)
	if err != nil {
		fmt.Println("log err")
		fmt.Println(err)
	}
	glog.Infoln("AccountFromConfig acc: ", acc)
	if acc.UserName != "" {
		if acc.Status == models.Blocked {
			glog.Errorln("account blocked", acc.UserName)
			return nil, errors.New("account has been blocked")
		}
	} else {
		if err := a.accountRepo.Create(ctx, models.Accounts{
			UserName: userInfo.Username,
			Status:   models.Active,
			Role:     models.Admin,
		}); err != nil {
			glog.Errorln(err)
			return nil, errors.New("system error")
		}
		acc.Role = models.Admin
		acc.Status = models.Active
	}

	key := fmt.Sprintf("%s:%s", common.PrefixLoginCode, userInfo.Username)
	code, err := a.serverCache.GetCode(key)
	if err != nil && err != redis.Nil {
		glog.Errorln("Login GetCode err: ", err)
		return nil, errors.New("system error")
	}

	account = &dto.Account{
		Username: userInfo.Username,
		Code:     string(code),
		Status:   string(acc.Status),
		Role:     string(acc.Role),
	}

	if err == redis.Nil {
		account.Code = uuid.New().String()
	}

	err = a.serverCache.SetCode([]byte(account.Code), key, time.Duration(a.authConfig.ExpiredTime))
	if err != nil {
		glog.Errorln("SetCode failed:", err)
		return nil, errors.New("system error")
	}

	resp = &dto.LoginResponseData{
		Account: account,
	}
	newAcc, _ := json.Marshal(account)
	glog.Infoln("Initialize access token")
	token, err := a.jWTAuth.InitializeToken(string(newAcc))
	if err != nil {
		glog.Errorln("InitializeToken failed:", err)
		return nil, errors.New("system error")
	}
	resp.AccessToken = token
	return resp, nil
}

func NewAuthenService(jWTAuth auth.JWTAuth, serverCache cache.ServerCache, accountRepo repositories.AccountRepository, authConfig *config.Auth) AuthenService {
	return &authenServiceImpl{
		serverCache: serverCache,
		jWTAuth:     jWTAuth,
		accountRepo: accountRepo,
		authConfig:  authConfig,
	}
}

type authenServiceImpl struct {
	serverCache cache.ServerCache
	jWTAuth     auth.JWTAuth
	accountRepo repositories.AccountRepository
	authConfig  *config.Auth
}
