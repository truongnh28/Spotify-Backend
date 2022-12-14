// Code generated by MockGen. DO NOT EDIT.
// Source: album.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	dto "spotify/dto"
	common "spotify/helper/common"

	gomock "github.com/golang/mock/gomock"
)

// MockAlbumService is a mock of AlbumService interface.
type MockAlbumService struct {
	ctrl     *gomock.Controller
	recorder *MockAlbumServiceMockRecorder
}

// MockAlbumServiceMockRecorder is the mock recorder for MockAlbumService.
type MockAlbumServiceMockRecorder struct {
	mock *MockAlbumService
}

// NewMockAlbumService creates a new mock instance.
func NewMockAlbumService(ctrl *gomock.Controller) *MockAlbumService {
	mock := &MockAlbumService{ctrl: ctrl}
	mock.recorder = &MockAlbumServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlbumService) EXPECT() *MockAlbumServiceMockRecorder {
	return m.recorder
}

// AddAlbum mocks base method.
func (m *MockAlbumService) AddAlbum(ctx context.Context, albumIn dto.Album) common.SubReturnCode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAlbum", ctx, albumIn)
	ret0, _ := ret[0].(common.SubReturnCode)
	return ret0
}

// AddAlbum indicates an expected call of AddAlbum.
func (mr *MockAlbumServiceMockRecorder) AddAlbum(ctx, albumIn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAlbum", reflect.TypeOf((*MockAlbumService)(nil).AddAlbum), ctx, albumIn)
}

// GetAlbumByID mocks base method.
func (m *MockAlbumService) GetAlbumByID(ctx context.Context, id uint) (dto.Album, common.SubReturnCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlbumByID", ctx, id)
	ret0, _ := ret[0].(dto.Album)
	ret1, _ := ret[1].(common.SubReturnCode)
	return ret0, ret1
}

// GetAlbumByID indicates an expected call of GetAlbumByID.
func (mr *MockAlbumServiceMockRecorder) GetAlbumByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlbumByID", reflect.TypeOf((*MockAlbumService)(nil).GetAlbumByID), ctx, id)
}

// GetAlbumByName mocks base method.
func (m *MockAlbumService) GetAlbumByName(ctx context.Context, name string) ([]dto.Album, common.SubReturnCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlbumByName", ctx, name)
	ret0, _ := ret[0].([]dto.Album)
	ret1, _ := ret[1].(common.SubReturnCode)
	return ret0, ret1
}

// GetAlbumByName indicates an expected call of GetAlbumByName.
func (mr *MockAlbumServiceMockRecorder) GetAlbumByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlbumByName", reflect.TypeOf((*MockAlbumService)(nil).GetAlbumByName), ctx, name)
}

// GetAllAlbum mocks base method.
func (m *MockAlbumService) GetAllAlbum(ctx context.Context) ([]dto.Album, common.SubReturnCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAlbum", ctx)
	ret0, _ := ret[0].([]dto.Album)
	ret1, _ := ret[1].(common.SubReturnCode)
	return ret0, ret1
}

// GetAllAlbum indicates an expected call of GetAllAlbum.
func (mr *MockAlbumServiceMockRecorder) GetAllAlbum(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAlbum", reflect.TypeOf((*MockAlbumService)(nil).GetAllAlbum), ctx)
}

// UpdateAlbum mocks base method.
func (m *MockAlbumService) UpdateAlbum(ctx context.Context, albumIn dto.Album) common.SubReturnCode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAlbum", ctx, albumIn)
	ret0, _ := ret[0].(common.SubReturnCode)
	return ret0
}

// UpdateAlbum indicates an expected call of UpdateAlbum.
func (mr *MockAlbumServiceMockRecorder) UpdateAlbum(ctx, albumIn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAlbum", reflect.TypeOf((*MockAlbumService)(nil).UpdateAlbum), ctx, albumIn)
}
