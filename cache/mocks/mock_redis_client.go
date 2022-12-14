// Code generated by MockGen. DO NOT EDIT.
// Source: redis_client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	v8 "github.com/go-redis/redis/v8"
	gomock "github.com/golang/mock/gomock"
)

// MockRedisClient is a mock of RedisClient interface.
type MockRedisClient struct {
	ctrl     *gomock.Controller
	recorder *MockRedisClientMockRecorder
}

// MockRedisClientMockRecorder is the mock recorder for MockRedisClient.
type MockRedisClientMockRecorder struct {
	mock *MockRedisClient
}

// NewMockRedisClient creates a new mock instance.
func NewMockRedisClient(ctrl *gomock.Controller) *MockRedisClient {
	mock := &MockRedisClient{ctrl: ctrl}
	mock.recorder = &MockRedisClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisClient) EXPECT() *MockRedisClientMockRecorder {
	return m.recorder
}

// Del mocks base method.
func (m *MockRedisClient) Del(ctx context.Context, key string) *v8.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", ctx, key)
	ret0, _ := ret[0].(*v8.IntCmd)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockRedisClientMockRecorder) Del(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockRedisClient)(nil).Del), ctx, key)
}

// Get mocks base method.
func (m *MockRedisClient) Get(ctx context.Context, key string) *v8.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(*v8.StringCmd)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockRedisClientMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRedisClient)(nil).Get), ctx, key)
}

// HDel mocks base method.
func (m *MockRedisClient) HDel(ctx context.Context, key, field string) *v8.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HDel", ctx, key, field)
	ret0, _ := ret[0].(*v8.IntCmd)
	return ret0
}

// HDel indicates an expected call of HDel.
func (mr *MockRedisClientMockRecorder) HDel(ctx, key, field interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HDel", reflect.TypeOf((*MockRedisClient)(nil).HDel), ctx, key, field)
}

// HGet mocks base method.
func (m *MockRedisClient) HGet(ctx context.Context, key, field string) *v8.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HGet", ctx, key, field)
	ret0, _ := ret[0].(*v8.StringCmd)
	return ret0
}

// HGet indicates an expected call of HGet.
func (mr *MockRedisClientMockRecorder) HGet(ctx, key, field interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HGet", reflect.TypeOf((*MockRedisClient)(nil).HGet), ctx, key, field)
}

// HSet mocks base method.
func (m *MockRedisClient) HSet(ctx context.Context, key, field string, values interface{}) *v8.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HSet", ctx, key, field, values)
	ret0, _ := ret[0].(*v8.IntCmd)
	return ret0
}

// HSet indicates an expected call of HSet.
func (mr *MockRedisClientMockRecorder) HSet(ctx, key, field, values interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HSet", reflect.TypeOf((*MockRedisClient)(nil).HSet), ctx, key, field, values)
}

// LPop mocks base method.
func (m *MockRedisClient) LPop(ctx context.Context, key string) *v8.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LPop", ctx, key)
	ret0, _ := ret[0].(*v8.StringCmd)
	return ret0
}

// LPop indicates an expected call of LPop.
func (mr *MockRedisClientMockRecorder) LPop(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPop", reflect.TypeOf((*MockRedisClient)(nil).LPop), ctx, key)
}

// LPush mocks base method.
func (m *MockRedisClient) LPush(ctx context.Context, key string, values ...interface{}) *v8.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LPush", varargs...)
	ret0, _ := ret[0].(*v8.IntCmd)
	return ret0
}

// LPush indicates an expected call of LPush.
func (mr *MockRedisClientMockRecorder) LPush(ctx, key interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPush", reflect.TypeOf((*MockRedisClient)(nil).LPush), varargs...)
}

// RPush mocks base method.
func (m *MockRedisClient) RPush(ctx context.Context, key string, values ...interface{}) *v8.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RPush", varargs...)
	ret0, _ := ret[0].(*v8.IntCmd)
	return ret0
}

// RPush indicates an expected call of RPush.
func (mr *MockRedisClientMockRecorder) RPush(ctx, key interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPush", reflect.TypeOf((*MockRedisClient)(nil).RPush), varargs...)
}

// Set mocks base method.
func (m *MockRedisClient) Set(ctx context.Context, key string, value interface{}, expireTime time.Duration) *v8.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, value, expireTime)
	ret0, _ := ret[0].(*v8.StatusCmd)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockRedisClientMockRecorder) Set(ctx, key, value, expireTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRedisClient)(nil).Set), ctx, key, value, expireTime)
}
