package common

import "errors"

var (
	PrefixLoginCode    = "spotify_login"
	CookieName         = "spotify"
	ErrEmptyStoredData = errors.New("empty data stored in kv store")
)

type ReturnCode int

const (
	Success ReturnCode = iota + 1
	Fail
)

var (
	returnCodeText = map[ReturnCode]string{
		Success: "Success",
		Fail:    "Failure",
	}
	subReturnCodeText = map[SubReturnCode]string{
		OK:                          "Success",
		InvalidRequest:              "Request is invalid, please check and try again!",
		SystemError:                 "System has been error, please try again later!",
		UsernameOrPasswordIncorrect: "Your username or password are incorrect, please check and try again!",
		Unauthorized:                "You have not login yet, please login and try again!",
		AccountBlocked:              "Your account has been blocked, please contact admin to enable your account!",
		NotPermission:               "You have not permission to perform this request, please check and try again!",
		TicketInvalid:               "Ticket is invalid, please try again!",
		NotSupport:                  "Currently, System is not support your request, please check and try again!",
		TokenInvalid:                "Token is invalid, please try again!",

		ErrExceedMaxRecordLimit: "Exceeded max record limit",
		ErrInvalidPageSize:      "Invalid page size",
		ErrInvalidPageNumber:    "Invalid page number",
	}
)

func (r ReturnCode) Message() string {
	return returnCodeText[r]
}

func (r ReturnCode) Int64() int64 {
	return int64(r)
}

type SubReturnCode int

const (
	OK SubReturnCode = iota + 1000
	InvalidRequest
	SystemError
	UsernameOrPasswordIncorrect
	Unauthorized
	NotPermission
	AccountBlocked
	TicketInvalid
	NotSupport
	TokenInvalid
	ErrExceedMaxRecordLimit
	ErrInvalidPageSize
	ErrInvalidPageNumber
)

func (r SubReturnCode) Message() string {
	return subReturnCodeText[r]
}

func (r SubReturnCode) Int64() int64 {
	return int64(r)
}
