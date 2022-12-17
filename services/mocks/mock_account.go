// Code generated by MockGen. DO NOT EDIT.
// Source: account.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	dto "spotify/dto"
	common "spotify/helper/common"

	gomock "github.com/golang/mock/gomock"
)

// MockAccountService is a mock of AccountService interface.
type MockAccountService struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServiceMockRecorder
}

// MockAccountServiceMockRecorder is the mock recorder for MockAccountService.
type MockAccountServiceMockRecorder struct {
	mock *MockAccountService
}

// NewMockAccountService creates a new mock instance.
func NewMockAccountService(ctrl *gomock.Controller) *MockAccountService {
	mock := &MockAccountService{ctrl: ctrl}
	mock.recorder = &MockAccountServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountService) EXPECT() *MockAccountServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAccountService) Create(ctx context.Context, req dto.CreateAccountRequest) common.SubReturnCode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, req)
	ret0, _ := ret[0].(common.SubReturnCode)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAccountServiceMockRecorder) Create(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccountService)(nil).Create), ctx, req)
}

// FindByUserName mocks base method.
func (m *MockAccountService) FindByUserName(ctx context.Context, username string) (dto.Account, common.SubReturnCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserName", ctx, username)
	ret0, _ := ret[0].(dto.Account)
	ret1, _ := ret[1].(common.SubReturnCode)
	return ret0, ret1
}

// FindByUserName indicates an expected call of FindByUserName.
func (mr *MockAccountServiceMockRecorder) FindByUserName(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserName", reflect.TypeOf((*MockAccountService)(nil).FindByUserName), ctx, username)
}

// Update mocks base method.
func (m *MockAccountService) Update(ctx context.Context, username string, req dto.UpdateAccountRequest) common.SubReturnCode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, username, req)
	ret0, _ := ret[0].(common.SubReturnCode)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockAccountServiceMockRecorder) Update(ctx, username, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAccountService)(nil).Update), ctx, username, req)
}