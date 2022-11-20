package dto

type ErrorResponse struct {
	Message string `json:"message"`
}

type StatusError struct {
	ReturnCode       int32  `json:"return_code,omitempty"`
	ReturnMessage    string `json:"return_message,omitempty"`
	SubReturnCode    int32  `json:"sub_return_code,omitempty"`
	SubReturnMessage string `json:"sub_return_message,omitempty"`
}
