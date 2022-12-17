// Code generated by MockGen. DO NOT EDIT.
// Source: interaction.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	models "spotify/models"

	gomock "github.com/golang/mock/gomock"
)

// MockInteractionRepository is a mock of InteractionRepository interface.
type MockInteractionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockInteractionRepositoryMockRecorder
}

// MockInteractionRepositoryMockRecorder is the mock recorder for MockInteractionRepository.
type MockInteractionRepositoryMockRecorder struct {
	mock *MockInteractionRepository
}

// NewMockInteractionRepository creates a new mock instance.
func NewMockInteractionRepository(ctrl *gomock.Controller) *MockInteractionRepository {
	mock := &MockInteractionRepository{ctrl: ctrl}
	mock.recorder = &MockInteractionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInteractionRepository) EXPECT() *MockInteractionRepositoryMockRecorder {
	return m.recorder
}

// AddInteraction mocks base method.
func (m *MockInteractionRepository) AddInteraction(ctx context.Context, userId, songId uint) (models.Interaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddInteraction", ctx, userId, songId)
	ret0, _ := ret[0].(models.Interaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddInteraction indicates an expected call of AddInteraction.
func (mr *MockInteractionRepositoryMockRecorder) AddInteraction(ctx, userId, songId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddInteraction", reflect.TypeOf((*MockInteractionRepository)(nil).AddInteraction), ctx, userId, songId)
}

// RemoveInteraction mocks base method.
func (m *MockInteractionRepository) RemoveInteraction(ctx context.Context, userId, songId uint) (models.Interaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveInteraction", ctx, userId, songId)
	ret0, _ := ret[0].(models.Interaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveInteraction indicates an expected call of RemoveInteraction.
func (mr *MockInteractionRepositoryMockRecorder) RemoveInteraction(ctx, userId, songId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveInteraction", reflect.TypeOf((*MockInteractionRepository)(nil).RemoveInteraction), ctx, userId, songId)
}
