package auth

import (
	"context"
	"sync"
)

const (
	O365Auth = "o365"
	LDAPAuth = "ldap"
)

var (
	authInstance sync.Map
)

type InitClientFunc func(config interface{}) IAuth

type UserInfo struct {
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Status   string   `json:"status"`
	Groups   []string `json:"groups"`
}

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type IAuth interface {
	Login(ctx context.Context, cert interface{}) (*UserInfo, error)
}

//InitAuth Have to call this func before using authClient
func InitAuth(name string, config interface{}, initFunc InitClientFunc) {
	cli := initFunc(config)
	authInstance.Store(name, cli)
}

//GetAuthClient have to call InitAuth first
func GetAuthClient(name string) IAuth {
	ins, ok := authInstance.Load(name)
	if ok {
		return ins.(IAuth)
	}
	return nil
}
