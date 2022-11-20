package helper

import (
	"reflect"

	"spotify/helper/common"
)

var (
	fieldReturnCodes = []string{"ReturnCode", "ReturnMessage", "SubReturnCode", "SubReturnMessage"}
)

func BuildResponseByReturnCode(resp interface{}, code common.ReturnCode, subCode common.SubReturnCode) {
	stype := reflect.ValueOf(resp).Elem()
	for index, fieldName := range fieldReturnCodes {
		field := stype.FieldByName(fieldName)
		if field.IsValid() && field.CanSet() {
			switch index {
			case 0:
				field.SetInt(code.Int64()) //set return_code
			case 1:
				field.SetString(code.Message()) //set return_message
			case 2:
				field.SetInt(subCode.Int64()) //set sub_return_code
			case 3:
				field.SetString(subCode.Message()) //set sub_return_message
			}
		}
	}
}
