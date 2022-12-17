// Code generated by MockGen. DO NOT EDIT.
// Source: playlist.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	dto "spotify/dto"
	models "spotify/models"

	gomock "github.com/golang/mock/gomock"
)

// MockPlayListRepository is a mock of PlayListRepository interface.
type MockPlayListRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPlayListRepositoryMockRecorder
}

// MockPlayListRepositoryMockRecorder is the mock recorder for MockPlayListRepository.
type MockPlayListRepositoryMockRecorder struct {
	mock *MockPlayListRepository
}

// NewMockPlayListRepository creates a new mock instance.
func NewMockPlayListRepository(ctrl *gomock.Controller) *MockPlayListRepository {
	mock := &MockPlayListRepository{ctrl: ctrl}
	mock.recorder = &MockPlayListRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlayListRepository) EXPECT() *MockPlayListRepositoryMockRecorder {
	return m.recorder
}

// AddPlayList mocks base method.
func (m *MockPlayListRepository) AddPlayList(ctx context.Context, playListIn dto.PlayList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPlayList", ctx, playListIn)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPlayList indicates an expected call of AddPlayList.
func (mr *MockPlayListRepositoryMockRecorder) AddPlayList(ctx, playListIn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPlayList", reflect.TypeOf((*MockPlayListRepository)(nil).AddPlayList), ctx, playListIn)
}

// GetAllPlayList mocks base method.
func (m *MockPlayListRepository) GetAllPlayList(ctx context.Context) ([]models.PlayList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPlayList", ctx)
	ret0, _ := ret[0].([]models.PlayList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPlayList indicates an expected call of GetAllPlayList.
func (mr *MockPlayListRepositoryMockRecorder) GetAllPlayList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPlayList", reflect.TypeOf((*MockPlayListRepository)(nil).GetAllPlayList), ctx)
}

// GetPlayListByID mocks base method.
func (m *MockPlayListRepository) GetPlayListByID(ctx context.Context, id uint) (models.PlayList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlayListByID", ctx, id)
	ret0, _ := ret[0].(models.PlayList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlayListByID indicates an expected call of GetPlayListByID.
func (mr *MockPlayListRepositoryMockRecorder) GetPlayListByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlayListByID", reflect.TypeOf((*MockPlayListRepository)(nil).GetPlayListByID), ctx, id)
}

// GetPlayListByName mocks base method.
func (m *MockPlayListRepository) GetPlayListByName(ctx context.Context, name string) ([]models.PlayList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlayListByName", ctx, name)
	ret0, _ := ret[0].([]models.PlayList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlayListByName indicates an expected call of GetPlayListByName.
func (mr *MockPlayListRepositoryMockRecorder) GetPlayListByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlayListByName", reflect.TypeOf((*MockPlayListRepository)(nil).GetPlayListByName), ctx, name)
}

// GetPlayListByUserID mocks base method.
func (m *MockPlayListRepository) GetPlayListByUserID(ctx context.Context, userId uint) ([]models.PlayList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlayListByUserID", ctx, userId)
	ret0, _ := ret[0].([]models.PlayList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlayListByUserID indicates an expected call of GetPlayListByUserID.
func (mr *MockPlayListRepositoryMockRecorder) GetPlayListByUserID(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlayListByUserID", reflect.TypeOf((*MockPlayListRepository)(nil).GetPlayListByUserID), ctx, userId)
}

// UpdatePlayList mocks base method.
func (m *MockPlayListRepository) UpdatePlayList(ctx context.Context, playListIn dto.PlayList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePlayList", ctx, playListIn)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePlayList indicates an expected call of UpdatePlayList.
func (mr *MockPlayListRepositoryMockRecorder) UpdatePlayList(ctx, playListIn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePlayList", reflect.TypeOf((*MockPlayListRepository)(nil).UpdatePlayList), ctx, playListIn)
}
