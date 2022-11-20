package helper

import (
	"github.com/stretchr/testify/assert"
	"gitlab.zalopay.vn/top/cicd/everest/dto"
	"gitlab.zalopay.vn/top/cicd/everest/helper/common"
	"testing"
)

func TestBuildResponseByReturnCode(t *testing.T) {
	type args struct {
		code    common.ReturnCode
		subCode common.SubReturnCode
	}
	testcases := []struct {
		name     string
		response *dto.StatusError
		args     args
		expected *dto.StatusError
	}{
		{
			name:     "OK",
			response: &dto.StatusError{},
			args: args{
				code:    common.Success,
				subCode: common.OK,
			},
			expected: &dto.StatusError{
				ReturnCode:       int32(common.Success),
				ReturnMessage:    "Success",
				SubReturnCode:    int32(common.OK),
				SubReturnMessage: "Success",
			},
		},
		{
			name:     "Fail + No Permission",
			response: &dto.StatusError{},
			args: args{
				code:    common.Fail,
				subCode: common.NotPermission,
			},
			expected: &dto.StatusError{
				ReturnCode:       2,
				ReturnMessage:    "Failure",
				SubReturnCode:    1005,
				SubReturnMessage: "You have not permission to perform this request, please check and try again!",
			},
		},
		{
			name:     "Fail + Unauthorized",
			response: &dto.StatusError{},
			args: args{
				code:    common.Fail,
				subCode: common.Unauthorized,
			},
			expected: &dto.StatusError{
				ReturnCode:       2,
				ReturnMessage:    "Failure",
				SubReturnCode:    1004,
				SubReturnMessage: "You have not login yet, please login and try again!",
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			BuildResponseByReturnCode(tc.response, tc.args.code, tc.args.subCode)
			assert.Equal(t, tc.expected, tc.response)
		})
	}
}
