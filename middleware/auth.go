package middleware

import (
	"context"
	"fmt"

	"spotify/cache"
	"spotify/dto"
	"spotify/helper"

	"spotify/helper/common"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func HTTPAuthentication(ctx *gin.Context) {
	var (
		token string
		err   error
	)

	token, err = ctx.Cookie(common.CookieName)
	if err != nil {
		glog.Errorln(err)
		ctx.AbortWithStatus(401)
		return
	}
	dt, err := helper.GetJWTInstance().GetDataFromToken(token)
	if err != nil {
		glog.Errorln("GetDataFromToken failed: ", err)
		ctx.AbortWithStatus(401)
		return
	}
	actor := dt.(*dto.Account)
	glog.Infoln("checkPermissionChange", actor)
	if checkPermissionChange(ctx, actor) {
		glog.Infoln("account is change permission", true)
		ctx.AbortWithStatus(401)
		return
	}
	glog.Infoln("actor", actor)
	ctx.Set("actor", actor)

	ctx.Next()
}

func checkPermissionChange(ctx context.Context, acc *dto.Account) bool {
	dt, err := cache.GetServerCacheInstance(common.SERVERCACHE).GetCode(fmt.Sprintf("%s:%s", common.PrefixLoginCode, acc.Username))
	if err != nil || len(dt) == 0 {
		return false
	}
	if string(dt) != acc.Code {
		return true
	}
	return false
}

func GetActorFromContext(ctx context.Context) (*dto.Account, bool) {
	if ctx, ok := ctx.(*gin.Context); ok {
		iActor, ok := ctx.Get("actor")
		if !ok {
			return nil, false
		}
		actor, ok := iActor.(*dto.Account)
		return actor, ok
	}
	iActor, ok := ctx.Value("actor").(*dto.Account)
	return iActor, ok
}

func SubjectFromJWT(ctx *gin.Context) string {
	iActor, ok := ctx.Get("actor")
	if !ok {
		return ""
	}
	if actor, ok := iActor.(*dto.Account); ok {
		return actor.Role
	}
	return ""
}
