package helper

import (
	"encoding/json"
	"sync"

	"spotify/config"
	"spotify/dto"

	"spotify/utils/auth"
)

var (
	jwtOnce     sync.Once
	jwtInstance auth.JWTAuth
)

func GetJWTInstance() auth.JWTAuth {
	jwtOnce.Do(func() {
		var (
			cfg = config.AuthConfig()
		)
		jwtInstance = auth.NewJWTAuth(cfg.SecretKey, cfg.ExpiredTime, getInfoFromToken)
	})
	return jwtInstance
}

func getInfoFromToken(dt string) (interface{}, error) {
	var (
		acc = &dto.Account{}
		err error
	)
	err = json.Unmarshal([]byte(dt), acc)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
