package auth

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWTValidateFunc func(string) (interface{}, error)
type JWTCheckSigningMethodFunc func(*jwt.Token) (interface{}, error)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type JWTAuth interface {
	InitializeToken(data string, signingMethod ...*jwt.SigningMethodHMAC) (string, error)
	UpdateDataToken(tokenStr, newData string, checkSigningMethodFunc ...JWTCheckSigningMethodFunc) (string, error)
	CheckValid(tokenStr string, checkSigningMethodFunc ...JWTCheckSigningMethodFunc) (bool, error)
	GetDataFromToken(tokenStr string, checkSigningMethodFunc ...JWTCheckSigningMethodFunc) (interface{}, error)
}

type jwtAuth struct {
	secretKey    string
	expireTime   int64
	validateFunc JWTValidateFunc
}

func NewJWTAuth(secretKey string, expiredTime int64, validateFunc JWTValidateFunc) JWTAuth {
	j := &jwtAuth{
		secretKey:    secretKey,
		expireTime:   expiredTime,
		validateFunc: validateFunc,
	}
	return j
}

//InitializeToken from data
func (j *jwtAuth) InitializeToken(data string, signingMethod ...*jwt.SigningMethodHMAC) (string, error) {
	var (
		signMethod = jwt.SigningMethodHS256
		claims     jwt.MapClaims
		token      *jwt.Token
	)
	if len(signingMethod) > 0 {
		signMethod = signingMethod[0]
	}
	token = jwt.New(signMethod)
	claims = token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Duration(j.expireTime) * time.Second).Unix()
	claims["data"] = data
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtAuth) isValid(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	return []byte(j.secretKey), nil
}

func (j *jwtAuth) CheckValid(tokenStr string, checkSigningMethodFunc ...JWTCheckSigningMethodFunc) (bool, error) {
	var (
		checkSigningMethod = j.isValid
		token              *jwt.Token
		err                error
	)
	if len(checkSigningMethodFunc) > 0 {
		checkSigningMethod = checkSigningMethodFunc[0]
	}
	token, err = jwt.Parse(tokenStr, checkSigningMethod)
	if err != nil {
		return false, err
	}
	_, ok := token.Claims.(*jwt.MapClaims)
	return ok && token.Valid, nil
}

func (j *jwtAuth) GetDataFromToken(tokenStr string, checkSigningMethodFunc ...JWTCheckSigningMethodFunc) (interface{}, error) {
	var (
		checkSigningMethod = j.isValid
		token              *jwt.Token
		err                error
	)
	if len(checkSigningMethodFunc) > 0 {
		checkSigningMethod = checkSigningMethodFunc[0]
	}
	token, err = jwt.Parse(tokenStr, checkSigningMethod)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		data, o := claims["data"].(string)
		if !o {
			return nil, fmt.Errorf("can not get data from token")
		}
		return j.validateFunc(data)
	}
	return nil, fmt.Errorf("token is invalid")
}

func (j *jwtAuth) UpdateDataToken(tokenStr, newData string, checkSigningMethodFunc ...JWTCheckSigningMethodFunc) (string, error) {
	var (
		checkSigningMethod = j.isValid
		token              *jwt.Token
		err                error
	)
	if len(checkSigningMethodFunc) > 0 {
		checkSigningMethod = checkSigningMethodFunc[0]
	}
	token, err = jwt.Parse(tokenStr, checkSigningMethod)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		claims["data"] = newData
		token.Claims = claims
		return token.SignedString([]byte(j.secretKey))
	}
	return "", fmt.Errorf("token is invalid")
}
